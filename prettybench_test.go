package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

// Test_main is a basic test of main. It writes the benchmark output file to
// STDIN and checks that the output is as expected.
func Test_main(t *testing.T) {
	// backup and resutore OsSTDIN to mock os.Stdin
	oldOsSTDIN := OsSTDIN
	defer func() { OsSTDIN = oldOsSTDIN }()

	nameFileOutput := "bench_01_output.txt"
	nameFileExpect := "bench_01_expect.txt"
	pathFileOutput := filepath.Join(".", "testdata", nameFileOutput)
	filePathExpect := filepath.Join(".", "testdata", nameFileExpect)

	var err error

	// Set input data to STDIN
	OsSTDIN, err = os.Open(pathFileOutput)
	if err != nil {
		t.Fatalf("error reading file %s: %s", pathFileOutput, err)
	}
	defer OsSTDIN.Close()

	// Capture STDOUT and STDERR
	out := capturer.CaptureOutput(func() {
		main()
	})

	// Read file of expect output
	data, err := os.ReadFile(filePathExpect)
	if err != nil {
		t.Fatalf("error reading file %s: %s", filePathExpect, err)
	}

	expect := strings.TrimSpace(string(data))
	actual := strings.TrimSpace(out)

	assert.Equal(t, expect, actual)
}
