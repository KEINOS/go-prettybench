package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/KEINOS/go-prettybench/prettify"
	"golang.org/x/tools/benchmark/parse"
)

var (
	errNotBenchLine = errors.New("not a bench line")

	// Regular expression to find bench data lines.
	benchLineMatcher = regexp.MustCompile(`^Benchmark.*\t.*\d+`)
	// Regular expression to find the test result.
	okLineMatcher = regexp.MustCompile(`^ok\s`)

	// Flag options.
	noPassthrough = flag.Bool("no-passthrough", false, "Don't print non-benchmark lines")

	// OsSTDIN is a copy of os.Stdin to ease mocking in tests.
	OsSTDIN = os.Stdin
	// OsExit is a copy of os.Exit to ease mocking in tests.
	OsExit = os.Exit
)

func main() {
	flag.Parse()

	currentBenchmark := &prettify.BenchOutputGroup{}
	scanner := bufio.NewScanner(OsSTDIN)

	// Loop each line of bench results from the input and print the parsed output
	for scanner.Scan() {
		lineText := scanner.Text()
		lineParsed, err := ParseLine(lineText)

		switch err {
		case errNotBenchLine:
			// Spit out the parsed lines if the line reaches to test result ans is "ok"
			if okLineMatcher.MatchString(lineText) {
				fmt.Print(currentBenchmark)                     // Spits out
				currentBenchmark = &prettify.BenchOutputGroup{} // Reset
			}

			// Print if noPassthrough flag is false (print info)
			if !*noPassthrough {
				fmt.Println(lineText)
			}
		case nil:
			currentBenchmark.AddLine(lineParsed)
		default:
			fmt.Fprintln(os.Stderr, "prettybench unrecognized line:")
			fmt.Println(lineText)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "prettybench unrecognized line:", err)
		OsExit(1)
	}
}

// ParseLine parses a line of benchmark output to testing.Benchmark.
// It returns an errNotBenchLine error if the line doesn't seem to contain benchmark data.
func ParseLine(line string) (*parse.Benchmark, error) {
	if !benchLineMatcher.MatchString(line) {
		return nil, errNotBenchLine
	}

	fields := strings.Split(line, "\t")
	if len(fields) < 3 {
		return nil, errNotBenchLine
	}

	return parse.ParseLine(line)
}
