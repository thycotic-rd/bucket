---
weight: 80
---

# DynamicFill
 
```go
import "time"

// ticker creates a channel which fires on a given interval
ticker := time.NewTicker(time.Second)

watchable := bucket.DynamicFill(300, ticker.C)

time.Sleep(time.Second * 3)
watchable.Cancel <- nil
``` 

Puts tokens into a bucket on an interval.

It returns a watchable which may be canceled.