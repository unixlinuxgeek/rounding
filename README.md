![push](https://github.com/unixlinuxgeek/rounding/actions/workflows/main.yml/badge.svg?event=push)


![Go rounding utility](https://raw.githubusercontent.com/unixlinuxgeek/logos/main/256x256/go.png)

### Go rounding utility

Run tests:
```shell
go test -test.v
```

Example:
```shell
$ go test -test.v
=== RUN   TestRound
    rounding_test.go:10: Generate random float64 number(1-9): 8.704279
    rounding_test.go:12: after round: 8.710
    rounding_test.go:13: rand: 8.704
    rounding_test.go:18: Test Round Passed!!!
--- PASS: TestRound (0.00s)
=== RUN   TestCeil
    rounding_test.go:24: Generate random float64 number(1-9): 4.019324
    rounding_test.go:26: after ceil: 4.020000
    rounding_test.go:27: rand: 4.019324
    rounding_test.go:32: Test Ceil Passed!!!
--- PASS: TestCeil (0.00s)
=== RUN   TestFloor
    rounding_test.go:38: Generate random float64 number(1-9): 1.450042
    rounding_test.go:40: after floor: 1.460000
    rounding_test.go:41: rand: 1.450042
    rounding_test.go:46: Test Floor Passed!!!
--- PASS: TestFloor (0.00s)
PASS
ok      github.com/unixlinuxgeek/rounding       0.002s
```
