# go-cache-benchmark

This is a simple benchmark of these LRU cache packages:
1. [Freecache](https://github.com/coocood/freecache)
2. [Hashicorp LRU](https://github.com/hashicorp/golang-lru)

Results are: 

| Benchmark name | iter | ns/op | B/op | alloc/op |
|----|----|----|----|----|
| BenchmarkFreeCacheGet-12 |          100000 |             23396 ns/op |            9584 B/op |        254 allocs/op |
| BenchmarkHashicorpGet-12 |        20000000 |               114 ns/op |              16 B/op |          1 allocs/op |
| | | | | |
| BenchmarkFreeCacheSet-12 |          200000 |              6895 ns/op |            2628 B/op |         51 allocs/op |
| BenchmarkHashicorpSet-12 |         1000000 |              1662 ns/op |             459 B/op |          6 allocs/op |


 