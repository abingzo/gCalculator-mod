# gCalculator-mod

##### 功能

- [x] 基于操作系统`StdI/O`,可以将该程序插入到任何地方
- [x] 加减乘除
- [x] 复杂及简单算式的解析
- [x] `20`位除法精度
- [x] 整数次幂,余
- [x] 通过模拟实现的二进制左移和右移
- [ ] 三角函数
- [ ] 对数函数

---

> 包含一个自实现的大数库，性能比较低，一些测试文件阐述了基本操作的性能,详细的测试可以执行`go test -bench=. ./test/alg_test.go ./test/task_test.go`查看，因为性能比较低，大于`2<<756`的数不能`1`秒之内算出

`Benchmark`

```shell
goos: darwin
goarch: amd64
BenchmarkBigNum/Add-8              74541             15642 ns/op
BenchmarkBigNum/Sub-8             200166              5753 ns/op
BenchmarkBigNum/Ride-8                 4         293896553 ns/op
BenchmarkBigNum/Std_Lib_Big-8   1000000000               0.000022 ns/op
BenchmarkBigNum/Except-8              24          46392853 ns/op
BenchmarkFormat/Decode-8          660482              1836 ns/op
BenchmarkFormat/Encode-8          725082              1619 ns/op
PASS
ok      command-line-arguments  8.504s
```

`Example`

```shell
echo "2^756" | ./eval
```

`OutPut`

```shell
379032737378102767370356320425415662904513187772631008578870126471203845870697482014374611530431269030880793627229265919475483409207718357286202948008100864063587640630090308972232735749901964068667724412528434753635948938919936
```

---

##### Build :dart:

```shell
make build-linux
```

##### Test :dagger:

```shell
make test
```

