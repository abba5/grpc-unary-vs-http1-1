# grpc-unary-vs-http1-1
benchmarking GRPC unary vs HTTP 1.1 in golang

golang: 1.14

GRPC Testing
```
# server
go build -o server -a ./cmd/grpc
./server

# test
go test -v -bench=. -run=^# github.com/abba5/grpc-unary-vs-http1-1/loadtest/grpc -count 5

# output
goos: linux
goarch: amd64
pkg: github.com/abba5/grpc-unary-vs-http1-1/loadtest/grpc
Benchmark_grpc_some
Benchmark_grpc_some-8              16276             73537 ns/op
Benchmark_grpc_some-8              16389             73444 ns/op
Benchmark_grpc_some-8              16368             73063 ns/op
Benchmark_grpc_some-8              16375             73416 ns/op
Benchmark_grpc_some-8              16320             73143 ns/op
PASS
ok      github.com/abba5/grpc-unary-vs-http1-1/loadtest/grpc    9.733s
```


HTTP Testing
```
# server
go build -o server -a ./cmd/http
./server

# test
go test -v -bench=. -run=^# github.com/abba5/grpc-unary-vs-http1-1/loadtest/http -count 5

# output
goos: linux
goarch: amd64
pkg: github.com/abba5/grpc-unary-vs-http1-1/loadtest/http
Benchmark_grpc_some
Benchmark_grpc_some-8              18914             62793 ns/op
Benchmark_grpc_some-8              19122             62545 ns/op
Benchmark_grpc_some-8              19137             62639 ns/op
Benchmark_grpc_some-8              19198             62928 ns/op
Benchmark_grpc_some-8              19207             62622 ns/op
PASS
ok      github.com/abba5/grpc-unary-vs-http1-1/loadtest/http    9.180s
```


