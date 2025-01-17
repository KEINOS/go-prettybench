<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# prettify

```go
import "github.com/KEINOS/go-prettybench/prettify"
```

Package prettify provides a set of functions to format benchmark results prettily\.

## Index

- [func FormatAllocsPerOp(l *bench.Benchmark) string](<#func-formatallocsperop>)
- [func FormatBytesAllocPerOp(l *bench.Benchmark) string](<#func-formatbytesallocperop>)
- [func FormatIterations(iter int) string](<#func-formatiterations>)
- [func FormatMegaBytesPerSecond(l *bench.Benchmark) string](<#func-formatmegabytespersecond>)
- [type BenchOutputGroup](<#type-benchoutputgroup>)
  - [func (g *BenchOutputGroup) AddLine(line *bench.Benchmark)](<#func-benchoutputgroup-addline>)
  - [func (g *BenchOutputGroup) ColumnNames() []string](<#func-benchoutputgroup-columnnames>)
  - [func (g *BenchOutputGroup) FormatTimeUnit(ns float64) string](<#func-benchoutputgroup-formattimeunit>)
  - [func (g *BenchOutputGroup) FormattRow(line *bench.Benchmark) (row []string)](<#func-benchoutputgroup-formattrow>)
  - [func (g *BenchOutputGroup) Sort()](<#func-benchoutputgroup-sort>)
  - [func (g *BenchOutputGroup) String() string](<#func-benchoutputgroup-string>)
  - [func (g *BenchOutputGroup) Tabulate(columnNames []string) *Table](<#func-benchoutputgroup-tabulate>)
- [type Table](<#type-table>)


## func [FormatAllocsPerOp](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L244>)

```go
func FormatAllocsPerOp(l *bench.Benchmark) string
```

FormatAllocsPerOp formats the AllocsPerOp value to an "allocs/op" string\.

## func [FormatBytesAllocPerOp](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L253>)

```go
func FormatBytesAllocPerOp(l *bench.Benchmark) string
```

FormatBytesAllocPerOp formats the AllocedBytesPerOp value to "B/op" string\.

## func [FormatIterations](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L262>)

```go
func FormatIterations(iter int) string
```

FormatIterations converts the number of iterations to a string\.

## func [FormatMegaBytesPerSecond](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L267>)

```go
func FormatMegaBytesPerSecond(l *bench.Benchmark) string
```

FormatMegaBytesPerSecond formats the MBPerS bench value to "MB/s" string\.

## type [BenchOutputGroup](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L23-L27>)

BenchOutputGroup holds a set of benchmark results and the options used to format them\.

```go
type BenchOutputGroup struct {
    ColNameSort string             // Column ID to sort by column if not empty
    Lines       []*bench.Benchmark // Benchmark lines parsed
    Measured    int                // Columns which are in use
}
```

### func \(\*BenchOutputGroup\) [AddLine](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L41>)

```go
func (g *BenchOutputGroup) AddLine(line *bench.Benchmark)
```

AddLine adds the "line" to the Lines' field and updates Measured field to indicate which measurements ware used\.

### func \(\*BenchOutputGroup\) [ColumnNames](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L48>)

```go
func (g *BenchOutputGroup) ColumnNames() []string
```

ColumnNames returns a list of column header names according to the current bench results\.

### func \(\*BenchOutputGroup\) [FormatTimeUnit](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L218>)

```go
func (g *BenchOutputGroup) FormatTimeUnit(ns float64) string
```

FormatTimeUnit uniforms the given nanoseconds \(float64\) to time unit of ns/μs/ms/s relativelly to the smallest number\.

### func \(\*BenchOutputGroup\) [FormattRow](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L72>)

```go
func (g *BenchOutputGroup) FormattRow(line *bench.Benchmark) (row []string)
```

FormattRow returns a row with formatted column values from the given benchmark line\.

### func \(\*BenchOutputGroup\) [Sort](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L96>)

```go
func (g *BenchOutputGroup) Sort()
```

Sort sorts by columns specified by command option\.

### func \(\*BenchOutputGroup\) [String](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L139>)

```go
func (g *BenchOutputGroup) String() string
```

String is a stringer for BenchOutputGroup which returns a formatted strin of the benchmark results\.

### func \(\*BenchOutputGroup\) [Tabulate](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L183>)

```go
func (g *BenchOutputGroup) Tabulate(columnNames []string) *Table
```

Tabulate returns a Table object with header of the given column names containing formatted comumns\.

## type [Table](<https://github.com/KEINOS/go-prettybench/blob/master/prettify/BenchOutputGroup.go#L30-L33>)

Table holds a set of cells and the maximum width of each column\.

```go
type Table struct {
    MaxLengths []int
    Cells      [][]string
}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
