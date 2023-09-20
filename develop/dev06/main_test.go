package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Prepare input data
	input := "1\tJohn Doe\tjohn@example.com\t555-1234\n2\tJane Smith\tjane@example.com\t555-5678"
	expectedOutput := "1\n2"

	// Create a scanner for input data
	scanner, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	defer scanner.Close()
	defer writer.Close()

	// Replace standard input with the configured scanner
	oldStdin := os.Stdin
	os.Stdin = scanner
	defer func() {
		os.Stdin = oldStdin
	}()

	// Capture the program's output
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
	}()

	// Write the input data to the scanner
	writer.WriteString(input)
	writer.Close()

	// Run the main function
	main()

	// Read the program's output
	w.Close()
	out, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	actualOutput := string(out)
	//expectedOutput = strings.TrimSuffix(expectedOutput, "\n")

	if actualOutput != expectedOutput {
		panic(fmt.Sprintf("Output mismatch!\nExpected: %s\nActual: %s", expectedOutput, actualOutput))
	}

	os.Exit(m.Run())
}
