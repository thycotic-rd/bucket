---
weight: 60
---

# Put

```go
// put 5 tokens back into the bucket
err = b.Put(5)
```

Put tokens into the bucket. 

Because this library represents numeric values as 32-bit integers the maximum value is
2,147,483,647.

