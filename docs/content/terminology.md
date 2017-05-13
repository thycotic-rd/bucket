---
weight: 15
---

# Terminology

A bucket is simple a key-value pair matching a string to an integer type.
The key represents the name and identifier for the object and the value the number
of tokens held within a bucket.

This is currently represented over two storage options, in-memory or redis. In-memory
stores the pairs within a Golang map while Redis stores it using its own key-value system.