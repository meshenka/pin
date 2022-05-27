# package pin

A simple project that generate a secure card pin.

## Performance

'''
make bench
go test -bench=. -count=5
goos: linux
goarch: amd64
pkg: github.com/meshenka/pin
cpu: Intel(R) Core(TM) i7-2670QM CPU @ 2.20GHz
BenchmarkGenerator-8   	 702619	     4317 ns/op
BenchmarkGenerator-8   	 235971	     4618 ns/op
BenchmarkGenerator-8   	 218743	     4640 ns/op
BenchmarkGenerator-8   	 279468	     4424 ns/op
BenchmarkGenerator-8   	 488704	     4178 ns/op
PASS
ok  	github.com/meshenka/pin	9.528s
'''
