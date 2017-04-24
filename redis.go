package distributed_token_bucket

import (
	"github.com/go-redis/redis"
	"time"
	"strconv"
	"errors"
)

const (
	// Redis and our redis client library allow lua scripting. We use the following script to conditionally decrement
	// a key-value pair by a given amount. This allows us not only to do both actions in a single round-trip to the
	// database but provide a lock-free implementation for logic that would otherwise be unsafe in a concurrent environment.
	//
	// Consider if we make two calls: .Get() and .Decr() with a conditional that .Decr() should only be called if .Get()
	// returns a certain amount... what if the value had been modified by a separate client during the
	// conditional statement? Then we might decrement the value incorrectly.
	//
	// Notably lua scripting is less performent then normal Redis commands thus this library should not share a database
	// with a super high-performance requirement
	//
	// Notice the use of tonumber() in Lua because everything Redis takes in and spits out is a string
	luaGetAndDecr = `
		local key = KEYS[1]
		local amount = tonumber(ARGV[1])
		local count = tonumber(redis.call("get", key))

		if count >= amount then
			return redis.call("DECRBY", key, amount)
		else
			error("Insufficient tokens/")
		end
	`
)

type RedisStorage struct {
	client *redis.Client
}

func (rs *RedisStorage) Ping() error {
	return rs.client.Ping().Err()
}

// bucket.Create will create a new bucket with the given parameters if one does not exist, if no bucket can be created it will return an error
func (rs *RedisStorage) Create(name string, capacity int) error {
	// query redis and see if the key already exists
	//
	// we receive tokensCount as a string in case 'name' already existed as a key and was not intended to be used here
	// i.e. "test" might be already taken and have some string value
	strTokensCount, err := rs.client.Get(name).Result()

	// if the name key does not exist in redis create it with the value of capacity
	if err == redis.Nil || len(strTokensCount) == 0 {
		// the last param 0 indicates the key will never expire
		return rs.client.Set(name, capacity, 0).Err()
	}

	// if strTokensCount can not be converted to an integer we will assume this key is already taken for something else
	// and return an error
	tokensCount, err := strconv.Atoi(strTokensCount)

	if err != nil {
		return err
	}

	// throw an error if a bucket already exists but is fully depleted, in order to prevent user confusion
	//
	// for example this might happen if a user does not think a bucket exists and it really does
	if tokensCount == 0 {
		return errors.New("Bucket exists in redis but contains a value of 0. Try putting tokens back into this bucket.")
	}

	// great we know the bucket exists and has an acceptable value
	return nil
}

// Executes a lua script which decrements the token value by tokensDesired if tokensDesired >= the token value.
func (rs *RedisStorage) Take(bucketName string, tokens int) error {
	return rs.client.Eval(luaGetAndDecr, []string{bucketName}, tokens).Err()
}

// Increment the token value by a given amount
func (rs *RedisStorage) Put(bucketName string, tokens int) error {
	return rs.client.IncrBy(bucketName, int64(tokens)).Err()
}

// attempt on a 500ms interval to asynchronously call bucket.Take until timeout is exceeded
// returns a channel which will fire nil or error on completion
func (rs *RedisStorage) Watch(bucketName string, tokens int, duration time.Duration) chan error {
	done := make(chan error)
	timeout := time.After(duration)

	go func(tokensDesired int, timeout <-chan time.Time, done chan error) {
		// time.Ticker returns a channel which fires every time the duration provided is passed
		ticker := time.NewTicker(time.Millisecond * 500)
		defer ticker.Stop()

		for {
			select {

			// attempt to take the desiredTokens on every ticker event
			case <-ticker.C:
				if err := rs.Take(bucketName, tokensDesired); err == nil {
					done <- nil
					break
				}

			// return an error if our timeout has passed
			case <-timeout:

				done <- errors.New("Watch timeout.")
				break
			}
		}
	}(tokens, timeout, done)

	return done
}