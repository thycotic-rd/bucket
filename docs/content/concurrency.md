---
weight: 85
---

# Concurrency 

This library was built to provide bucket primitives for a distributed system. As such
both the in-memory and redis storage back-ends provide concurrent safe access to tokens.

For the Redis back-end single operations are safe automatically thanks to the underlying single-threaded
design of Redis. Methods needing more then one operation are done with Lua scripting which is executed as
a single network call.

For the in-memory back-end everything is protected by a read-write mutex.