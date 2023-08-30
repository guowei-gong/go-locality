打开终端，进入项目根目录，执行如下命令。

`$go test -bench=. -count=3`

输出:
```shell
goos: darwin
goarch: amd64
pkg: weizicoding.com/go-locality
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkLoopStep1-12             413377              2870 ns/op
BenchmarkLoopStep1-12             417328              2909 ns/op
BenchmarkLoopStep1-12             425328              2910 ns/op
BenchmarkLoopStep2-12             222739              5198 ns/op
BenchmarkLoopStep2-12             229764              5124 ns/op
BenchmarkLoopStep2-12             227998              5199 ns/op
BenchmarkLoopStep3-12             289923              3824 ns/op
BenchmarkLoopStep3-12             320758              3833 ns/op
BenchmarkLoopStep3-12             309706              3859 ns/op
BenchmarkLoopStep4-12             235686              5172 ns/op
BenchmarkLoopStep4-12             227913              5201 ns/op
BenchmarkLoopStep4-12             213157              5201 ns/op
BenchmarkLoopStep5-12             218772              5646 ns/op
BenchmarkLoopStep5-12             212197              5708 ns/op
BenchmarkLoopStep5-12             207944              5932 ns/op
BenchmarkLoopStep6-12             174680              6750 ns/op
BenchmarkLoopStep6-12             181030              6595 ns/op
BenchmarkLoopStep6-12             183346              6600 ns/op
BenchmarkLoopStep12-12            123316             10138 ns/op
BenchmarkLoopStep12-12            123435              9518 ns/op
BenchmarkLoopStep12-12            125121              9746 ns/op
BenchmarkLoopStep16-12             91929             12622 ns/op
BenchmarkLoopStep16-12             93736             12830 ns/op
BenchmarkLoopStep16-12             90597             15198 ns/op
PASS
ok      weizicoding.com/go-locality     34.396s
```

跨度 1 时的性能，是跨度 16 时的 4 ~ 5 倍。
