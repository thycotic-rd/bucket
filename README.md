[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/b3ntly/bucket/master/LICENSE.txt) 
[![Build Status](https://travis-ci.org/b3ntly/bucket.svg?branch=master)](https://travis-ci.org/b3ntly/bucket)
[![Coverage Status](https://coveralls.io/repos/github/b3ntly/bucket/badge.svg?branch=master)](https://coveralls.io/github/b3ntly/bucket?branch=master?q=1) 
[![GoDoc](https://godoc.org/github.com/b3ntly/bucket?status.svg)](https://godoc.org/github.com/b3ntly/bucket)
[![Documentation](https://readthedocs.org/projects/docs/badge/?version=latest)](https://b3ntly.github.io/bucket)

## Features

* In-Memory or Redis based storage
* Concurrency-safe 
* Put, Take, Count, DynamicFill, and many more primitives
* Designed for distributed systems
* Well tested
* Well documented


## Install

```shell
go get github.com/b3ntly/bucket
```

## Documentation

Can be found [here](https://b3ntly.github.io/bucket)

## Notes

* Test coverage badge is stuck in some cache and is out of date, click the badge to see the actual current coverage


## Changelog for version 0.2

* Abstracted storage to its own interface, see storage.go and redis.go for examples
* Added in-memory storage option
* Replaced storageOptions parameter with IStorage, allowing a single storage option to be shared
  between multiple buckets. This should make it much more efficient to have a large number of buckets, i.e.
  per logged in user.
  
## Changelog for version 0.3
  
* Renamed the repository from distributed-token-bucket to bucket
* Moved storage interface and providers to the /storage subpackage
* Added unit testing to the /storage subpackage
* Added watchable.go and changed signatures of all async functions to return a watchable
* Fixed examples
* Added more documentation and comments  

## Changelog for version 0.4

* Shortened "constructor" names
* Default options
* Better "constructor" signatures
* bucket.DynamicFill()
* bucket.TakeAll()

# Changelog for version 0.5

* Added full documentation via Hugo

## Benchmarks

```golang
go test -bench .
```

Benchmarks were done on a 2.2GHz quad-core Intel Core i7 with 16GB of RAM.

Version 0.4

Memory

| Benchmark                | Operations | ns/op  |
|--------------------------|------------|--------|
| BenchmarkBucket_Create-8 | 10000      | 715 ns/op |
| BenchmarkBucket_Take-8   | 30000      | 132 ns/op  |
| BenchmarkBucket_Put-8    | 50000      | 142 ns/op  |

Redis

| Benchmark                | Operations | ns/op  |
|--------------------------|------------|--------|
| BenchmarkBucket_Create-8 | 10000      | 98582 ns/op |
| BenchmarkBucket_Take-8   | 30000      | 47716 ns/op  |
| BenchmarkBucket_Put-8    | 50000      | 31350 ns/op  |
