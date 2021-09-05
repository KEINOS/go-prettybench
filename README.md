[![Go Report Card](https://goreportcard.com/badge/github.com/KEINOS/go-prettybench)](https://goreportcard.com/report/github.com/KEINOS/go-prettybench)
[![CodeQL](https://github.com/KEINOS/go-prettybench/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/KEINOS/go-prettybench/actions/workflows/codeql-analysis.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-prettybench.svg)](https://pkg.go.dev/github.com/KEINOS/go-prettybench)

# Prettybench <sub><sup><sup>FORK</sup></sup></sub>

`prettybench` is a formating tool to **"pretty-print" the benchmark results of Go**. It helps a bit to make it nicer for humans to read ;-)

> This repo, [KEINOS/go-prettybench](https://github.com/KEINOS/go-prettybench), is a fork of [cespare/prettybench](https://github.com/cespare/prettybench) with `-sort` option implementation.

## Usage

```shellsession
$ prettybench -h
Usage of prettybench:
  -no-passthrough
        Don't print non-benchmark lines
  -sort string
        Sort by column (string: name, iter, time, bytes, allocs)
```

```shellsession
$ # Basic usage
$ go test -bench=. | go-prettybench
```
```shellsession
$ # Sort by iteration (faster iteration)
$ go test -bench=. | go-prettybench -sort iter
```

To sort more than one column, pipe the command as below.

```shellsession
$ # Find the fastest benchmark
$ go test -bench=. | go-prettybench -sort=iter | go-prettybench -sort=time | go-prettybench -sort=bytes
```

## Sample Output

```shellsession
$ go test -failfast -benchmem -shuffle=on -benchtime=10s -count=5 -bench . | go-prettybench -sort iter
-test.shuffle 1630544600714497930
goos: linux
goarch: amd64
pkg: github.com/KEINOS/go-blake3-example
cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
PASS
benchmark                   iter        time/iter   bytes alloc        allocs
---------                   ----        ---------   -----------        ------
BenchmarkBlake3_32Byte-2   21348     579.62 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_32Byte-2   20786     580.41 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_64Byte-2   20275    1309.79 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_32Byte-2   19292     698.38 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_64Byte-2   18704     651.92 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_32Byte-2   18151     604.64 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_32Byte-2   16812     688.91 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_64Byte-2   15891     818.48 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_64Byte-2   15384    1756.33 μs/op        0 B/op   0 allocs/op
BenchmarkBlake3_64Byte-2   15085     814.88 μs/op        0 B/op   0 allocs/op
BenchmarkFNV1_4Byte-2       8626    1930.83 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1_4Byte-2       8566    1956.97 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1A_4Byte-2      8452    1499.16 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1A_4Byte-2      8396    1401.46 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1A_4Byte-2      8137    1661.25 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1A_4Byte-2      8035    1442.88 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1_4Byte-2       7635    1725.90 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1A_4Byte-2      7329    1369.91 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1_4Byte-2       6999    1625.71 μs/op        8 B/op   1 allocs/op
BenchmarkFNV1_4Byte-2       6796    1662.90 μs/op        8 B/op   1 allocs/op
BenchmarkSHA3_32Byte-2      3169    5188.47 μs/op      960 B/op   4 allocs/op
BenchmarkSHA3_32Byte-2      3120    4498.84 μs/op      960 B/op   4 allocs/op
BenchmarkSHA3_32Byte-2      2427    6027.57 μs/op      960 B/op   4 allocs/op
BenchmarkSHA3_32Byte-2      2258    5363.21 μs/op      960 B/op   4 allocs/op
BenchmarkSHA3_32Byte-2      2218    5628.21 μs/op      960 B/op   4 allocs/op
BenchmarkSHA3_64Byte-2      1698    6975.16 μs/op     1024 B/op   4 allocs/op
BenchmarkSHA3_64Byte-2      1662    7139.66 μs/op     1024 B/op   4 allocs/op
BenchmarkSHA3_64Byte-2      1653    7034.54 μs/op     1024 B/op   4 allocs/op
BenchmarkSHA3_64Byte-2      1633    8689.48 μs/op     1024 B/op   4 allocs/op
BenchmarkSHA3_64Byte-2      1574    7693.19 μs/op     1024 B/op   4 allocs/op
ok      github.com/KEINOS/go-blake3-example     557.763s
```

## Insatall

- `go install github.com/KEINOS/go-prettybench@latest` (Go v1.16 or above)
- `go get -u github.com/KEINOS/go-prettybench@latest` (Go v1.16 or below)

> **Note:** As of Go v.1.17 installing executables with '`go get -u`' in module mode is deprecated.

## Statuses

- Supported Go versions and monthly test results:
  - [![go1.14](https://github.com/KEINOS/go-prettybench/actions/workflows/runGo1_14.yml/badge.svg)](https://github.com/KEINOS/go-prettybench/actions/workflows/runGo1_14.yml)
  [![go1.15](https://github.com/KEINOS/go-prettybench/actions/workflows/runGo1_15.yml/badge.svg)](https://github.com/KEINOS/go-prettybench/actions/workflows/runGo1_15.yml)
  [![go1.16](https://github.com/KEINOS/go-prettybench/actions/workflows/runGo1_16.yml/badge.svg)](https://github.com/KEINOS/go-prettybench/actions/workflows/runGo1_16.yml)
  [![go1.17](https://github.com/KEINOS/go-prettybench/actions/workflows/runGo1_17.yml/badge.svg)](https://github.com/KEINOS/go-prettybench/actions/workflows/runGo1_17.yml)
- Security Alert Status
  - [Security advisories](https://github.com/KEINOS/go-prettybench/security/advisories)
  - [Code scanning alerts (CodeQL)](https://github.com/KEINOS/go-prettybench/security/code-scanning)
  - [Dependencies vulnerability scan (Dependabot)](https://github.com/KEINOS/go-prettybench/security/dependabot)

## Notes

- About the branch:
  - [main](https://github.com/KEINOS/go-prettybench/tree/original): Main branch of this repo to PR.
  - [original](https://github.com/KEINOS/go-prettybench/tree/original): Equal to upstream's [`master`](https://github.com/cespare/prettybench/tree/master) branch.
- Right now the units for the time are chosen based on the smallest value in the column.
- Prettybench has to buffer all the rows of output before it can print them (for
  column formatting), so you won't see intermediate progress. If you want to see
  that too, you could tee your output so that you see the unmodified version as
  well. If you do this, you'll want to use the prettybench's `-no-passthrough`
  flag so it doesn't print all the other lines (because then they'd be printed
  twice):

        $ go test -bench=. | tee >(prettybench -no-passthrough)

- Any PR for better are welcome!

## Wishlist

- [x] Add a `-sort` flag to sort by the given column
- [ ] Show average and dispersity (in ±%) per benchmark name (if more than one exists)
