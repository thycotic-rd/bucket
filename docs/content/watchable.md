---
weight: 70
---

# Watchable

```go 
type Watchable struct {
	Success chan error

	Cancel  chan error

	Failed  chan error

	// The final observable which the user is likely to read from. Though it can only be fired once it is buffered
	// so that is may be ignored.
	Finished chan error
}

watchable := b.Watch(10, time.Second * 5)
watchable.Cancel <- errors.New("I wasn't happy with this watcher :/")

// capture the error as the watcher exits
err := <- watchable.Done()
// err.Error() => "I wasn't happy with this watcher :/"
```

A basic structure from which to cancel or observe an asynchronous action.

