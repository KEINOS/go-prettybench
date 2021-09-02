package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

const (
	nameDirTest = "testdata"
)

func pathFileTest(t *testing.T, nameFileTest string) string {
	t.Helper()

	path := filepath.Join(".", nameDirTest, nameFileTest)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("file not found: %v\nerror: %v", path, err)
	}

	return path
}

// Test_main is a basic test of main. It writes the benchmark output file to
// STDIN and checks that the output is as expected.
func Test_main(t *testing.T) {
	// backup and resutore OsSTDIN to mock os.Stdin
	oldOsSTDIN := OsSTDIN
	defer func() { OsSTDIN = oldOsSTDIN }()

	pathFileOutput := pathFileTest(t, "sample_benchmark01.txt")
	filePathExpect := pathFileTest(t, "Test_main-expect.txt")

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

func Test_main_sort(t *testing.T) {
	// Backup and restore OsSTDIN to mock os.Stdin
	oldOsSTDIN := OsSTDIN
	defer func() { OsSTDIN = oldOsSTDIN }()

	// Backup and restore os.Args to mock os.Args
	oldOsArgs := os.Args
	defer func() { os.Args = oldOsArgs }()

	pathFileOutput := pathFileTest(t, "sample_benchmark01.txt")
	filePathExpect := pathFileTest(t, "Test_main_sort-expect.txt")

	var err error

	// Set input data to STDIN
	OsSTDIN, err = os.Open(pathFileOutput)
	if err != nil {
		t.Fatalf("error reading file %s: %s", pathFileOutput, err)
	}
	defer OsSTDIN.Close()

	// Capture STDOUT and STDERR
	out := capturer.CaptureOutput(func() {
		// Mock user command line flat
		os.Args = []string{t.Name(), "-sort", "time"}
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
