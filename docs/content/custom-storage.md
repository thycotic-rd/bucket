---
weight: 45
---

# Custom Storage


```go 
import (
    "github.com/b3ntly/bucket/storage"
    "github.com/go-redis/redis"
)    

store := &storage.RedisStorage{
    Client: redis.NewClient(&redis.Options{
        Addr: ":6379",
        DB: 5,
        PoolSize: 30,
    }),
}	

b2, err := bucket.New(&bucket.Options{
    Capacity: 10,
    Name: "My redis bucket with custom config",
    Storage: store,
})
```

Instead of the provided In-memory and Redis storage back-ends you can pass any object that fulfills
the Storage interface.

Here's an example which modifies the Redis storage with a custom configuration.
It uses [Go-Redis](https://github.com/go-redis/redis) internally.