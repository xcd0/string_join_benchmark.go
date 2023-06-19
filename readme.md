# Golangにおける文字列連結処理のベンチマーク

https://gist.github.com/matope/89ceb77d39e3c0407d67

を `m2 += v + ","` に修正したベンチマーク。  
`make`で実行できるようにしている。

## go version go1.20.5 linux/amd64 5.15.90.1-microsoft-standard-WSL2 x86_64 GNU/Linux
```
goos: linux
goarch: amd64
pkg: string_join_benchmark
cpu: Intel(R) Core(TM) i7-10850H CPU @ 2.70GHz
BenchmarkStringsJoin____-4   	10386705	       116.0 ns/op	     160 B/op	       2 allocs/op
BenchmarkAppendOperator_-4   	 3582710	       328.3 ns/op	     408 B/op	       8 allocs/op
BenchmarkHardCoding_____-4   	13436542	        88.94 ns/op	      80 B/op	       1 allocs/op
BenchmarkByteArray______-4   	 7242062	       166.3 ns/op	     320 B/op	       5 allocs/op
BenchmarkCapByteArray___-4   	20393071	        54.47 ns/op	      80 B/op	       1 allocs/op
BenchmarkBytesBuffer____-4   	 5692730	       211.2 ns/op	     272 B/op	       3 allocs/op
BenchmarkCapBytesBuffer_-4   	 9032692	       124.3 ns/op	      80 B/op	       1 allocs/op
BenchmarkCapBytesBuffer2-4   	13200518	        89.68 ns/op	      80 B/op	       1 allocs/op
PASS
ok  	string_join_benchmark	10.640s
```
