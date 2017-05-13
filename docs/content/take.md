---
weight: 50
---

# Take 

```go
err = b.Take(5)

err = b.Take(5)
// err.Error() => "Insufficient tokens."
```

You can take as many tokens from a bucket as a bucket currently contains. If not enough tokens exist
in the bucket, or if some other error such as a network error occurs, it will return an error. Operations against
a bucket occur transactionally so it may be assumed that upon an error the actual value of the bucket,
even in a distributed system, did not change.
