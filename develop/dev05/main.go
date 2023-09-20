package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func grep(searchTerm string, options Options, reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)

	var outputLines []string
	var lineNum int
	var prevLines []string

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if options.IgnoreCase {
			line = strings.ToLower(line)
			searchTerm = strings.ToLower(searchTerm)
		}

		match := strings.Contains(line, searchTerm)

		if (options.Inverse && !match) || (!options.Inverse && match) {
			if options.LineNum {
				outputLines = append(outputLines, fmt.Sprintf("%d:%s", lineNum, line))
			} else {
				outputLines = append(outputLines, line)
			}
		}

		if match {
			if options.Before > 0 {
				outputLines = append(outputLines, prevLines...)
				prevLines = nil
			}

			if options.After > 0 {
				outputLines = append(outputLines, line)
				for i := 0; i < options.After; i++ {
					if scanner.Scan() {
						nextLine := scanner.Text()
						lineNum++
						outputLines = append(outputLines, nextLine)
					} else {
						break
					}
				}
			}
		} else if options.Before > 0 {
			prevLines = append(prevLines, line)
			if len(prevLines) > options.Before {
				prevLines = prevLines[1:]
			}
		}

		if options.Count && !options.Inverse {
			options.NumMatches++
		}
	}

	for _, line := range outputLines {
		fmt.Fprintln(writer, line)
	}
}

type Options struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Inverse    bool
	Fixed      bool
	LineNum    bool
	NumMatches int
}

func main() {
	var options Options
	flag.IntVar(&options.After, "A", 0, "print +N lines after match")
	flag.IntVar(&options.Before, "B", 0, "print +N lines before match")
	flag.IntVar(&options.Context, "C", 0, "print ±N lines around match")
	flag.BoolVar(&options.Count, "c", false, "print count of matching lines")
	flag.BoolVar(&options.IgnoreCase, "i", false, "ignore case")
	flag.BoolVar(&options.Inverse, "v", false, "invert match")
	flag.BoolVar(&options.Fixed, "F", false, "fixed match")
	flag.BoolVar(&options.LineNum, "n", false, "print line numbers")

	flag.Parse()

	searchTerm := flag.Arg(0)
	if searchTerm == "" {
		fmt.Println("Usage: grep [options] <pattern> [filename]")
		os.Exit(1)
	}

	var reader io.Reader
	filename := flag.Arg(1)
	if filename == "" {
		reader = os.Stdin
	} else {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	grep(searchTerm, options, reader, os.Stdout)

	if options.Count {
		fmt.Printf("Total matches: %d\n", options.NumMatches)
	}
}
