---
weight: 40
---

# NewWithRedis
 
```go 
b, err := bucket.NewWithRedis(&bucket.Options{
    Capacity: 10,
    Name: "My redis bucket with default config",
})
``` 

Create a new bucket with a redis backend. By default it will attempt
to connect to a Redis instance with the default Redis configuration,
running at 127.0.0.1:6379.
