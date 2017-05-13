---
weight: 30
---

# New

```go
// bucket will use in-memory storage as default
b, err := bucket.New(&bucket.Options{
    Name: "my_bucket",
    Capacity: 10,
})
```

Creating a new bucket sets a key-value pair with whatever storage back-end you are using. By default this
will happen in-memory. With the Redis back-end this will set the key "my_bucket" with the value of 10.

There are some protections in place for name-collisions. For example when using the redis backend if the designated key
already contains a value that is not an integer, or is an integer with a value of 0, bucket.New will return an error.

Bucket names are not unique. If two buckets share the same name and storage back-end they will share the same value. 


