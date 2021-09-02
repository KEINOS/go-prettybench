/*
Package prettify provides a set of functions to format benchmark results prettily.
*/
package prettify

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/tools/benchmark/parse"
)

// ----------------------------------------------------------------------------
//  Structs
// ----------------------------------------------------------------------------

type BenchOutputGroup struct {
	Lines    []*parse.Benchmark
	Measured int // Columns which are in use
}

type Table struct {
	MaxLengths []int
	Cells      [][]string
}

// ----------------------------------------------------------------------------
//  Methods for BenchOutputGroup
// ----------------------------------------------------------------------------

// AddLine adds the "line" to the Lines' field and updates Measured field to
// indicate which measurements ware used.
func (g *BenchOutputGroup) AddLine(line *parse.Benchmark) {
	g.Lines = append(g.Lines, line)
	g.Measured |= line.Measured // bitwise OR to up the bit flags
}

// ColumnNames returns a list of column header names according to the current
// bench results.
func (g *BenchOutputGroup) ColumnNames() []string {
	// Add basic header
	columnNames := []string{"benchmark", "iter", "time/iter"}

	// Add "throughput" column if found bitwisely
	if (g.Measured & parse.MBPerS) > 0 {
		columnNames = append(columnNames, "throughput")
	}

	// Add "bytes alloc" column if found bitwisely
	if (g.Measured & parse.AllocedBytesPerOp) > 0 {
		columnNames = append(columnNames, "bytes alloc")
	}

	// Add "allocs" column if found bitwisely
	if (g.Measured & parse.AllocsPerOp) > 0 {
		columnNames = append(columnNames, "allocs")
	}

	return columnNames
}

// FormattRow returns a row with formatted column values from the given
// benchmark line.
func (g *BenchOutputGroup) FormattRow(line *parse.Benchmark) (row []string) {
	// Create a row
	row = []string{
		line.Name,
		FormatIterations(line.N),
		g.FormatTimeUnit(line.NsPerOp), // uniform nanoseconds to ns/μs/ms/s
	}

	if (g.Measured & parse.MBPerS) > 0 {
		row = append(row, FormatMegaBytesPerSecond(line))
	}

	if (g.Measured & parse.AllocedBytesPerOp) > 0 {
		row = append(row, FormatBytesAllocPerOp(line))
	}

	if (g.Measured & parse.AllocsPerOp) > 0 {
		row = append(row, FormatAllocsPerOp(line))
	}

	return row
}

// String is a stringer for BenchOutputGroup which returns a formatted strin of
// the benchmark results.
func (g *BenchOutputGroup) String() string {
	if len(g.Lines) == 0 {
		return ""
	}

	// Get column names
	columnNames := g.ColumnNames()
	// Get blank table with headers(columnNames and underlines)
	table := g.Tabulate(columnNames)

	// Loop table's cells and write them to the buffer

	var buf bytes.Buffer

	for _, row := range table.Cells {
		for i, cell := range row {
			var format string

			switch i {
			case 0:
				format = "%%-%ds   "
			case len(row) - 1:
				format = "%%%ds"
			default:
				format = "%%%ds   "
			}

			fmt.Fprintf(&buf, fmt.Sprintf(format, table.MaxLengths[i]), cell)
		}

		fmt.Fprint(&buf, "\n")
	}

	return buf.String()
}

// Tabulate creates a blank table with an underlined header of the given column
// names.
func (g *BenchOutputGroup) Tabulate(columnNames []string) *Table {
	table := &Table{Cells: [][]string{columnNames}}

	// Create underline with the length of each column
	underlines := []string{}

	for _, name := range columnNames {
		underlines = append(underlines, strings.Repeat("-", len(name)))
	}
	// Add underline to column name
	table.Cells = append(table.Cells, underlines)

	// Read each line of the benchmark results and add it to the table
	// after formatting it.
	for _, line := range g.Lines {
		table.Cells = append(table.Cells, g.FormattRow(line))
	}

	// Find max length of column names
	for i := range columnNames {
		maxLength := 0
		for _, row := range table.Cells {
			if len(row[i]) > maxLength {
				maxLength = len(row[i])
			}
		}

		table.MaxLengths = append(table.MaxLengths, maxLength)
	}

	return table
}

// FormatTimeUnit uniforms the given nanoseconds (float64) to time unit of
// ns/μs/ms/s relativelly to the smallest number.
func (g *BenchOutputGroup) FormatTimeUnit(ns float64) string {
	// Find the smallest time
	smallest := g.Lines[0].NsPerOp
	for _, line := range g.Lines[1:] {
		if line.NsPerOp < smallest {
			smallest = line.NsPerOp
		}
	}

	switch {
	case smallest < float64(10000*time.Nanosecond):
		return fmt.Sprintf("%.2f ns/op", ns)
	case smallest < float64(time.Millisecond):
		return fmt.Sprintf("%.2f μs/op", ns/1000)
	case smallest < float64(10*time.Second):
		return fmt.Sprintf("%.2f ms/op", ns/1e6)
	default:
		return fmt.Sprintf("%.2f s/op", ns/1e9)
	}
}

// ----------------------------------------------------------------------------
//  Functions for generating output
// ----------------------------------------------------------------------------

// FormatAllocsPerOp formats the AllocsPerOp value to an "allocs/op" string.
func FormatAllocsPerOp(l *parse.Benchmark) string {
	if (l.Measured & parse.AllocsPerOp) == 0 {
		return ""
	}

	return fmt.Sprintf("%d allocs/op", l.AllocsPerOp)
}

// FormatBytesAllocPerOp formats the AllocedBytesPerOp value to "B/op" string.
func FormatBytesAllocPerOp(l *parse.Benchmark) string {
	if (l.Measured & parse.AllocedBytesPerOp) == 0 {
		return ""
	}

	return fmt.Sprintf("%d B/op", l.AllocedBytesPerOp)
}

// FormatIterations converts the number of iterations to a string.
func FormatIterations(iter int) string {
	return strconv.FormatInt(int64(iter), 10)
}

// FormatMegaBytesPerSecond formats the MBPerS bench value to "MB/s" string.
func FormatMegaBytesPerSecond(l *parse.Benchmark) string {
	if (l.Measured & parse.MBPerS) == 0 {
		return ""
	}

	return fmt.Sprintf("%.2f MB/s", l.MBPerS)
}
