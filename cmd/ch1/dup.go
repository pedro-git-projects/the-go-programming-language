package ch1

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Note ctrl+d enter EOF in the terminal

// Dup prints each line that appears more than once in the standard input,
// preceded by its count
func Dup() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
		// line := input.Text()
		// counts[line] = counts[line]  + 1
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// countLines takes a pointer to a file and a map of strings to ints
// and when a line is repeated adds it to the map and increments
// its value
func countLines(f *os.File, counts map[string]int) error {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		err := input.Err()
		if err != nil {
			return err
		}
	}
	return nil
}

// Dup2 reads the lines of as many files as are passed as aruguments when
// invoking the program
// and prints the repeated ones with the number of repetitions
func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		err := countLines(os.Stdin, counts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "DupInFile: %v\n", err)
		}
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "DupInFile: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Dup3 uses ReadFile to read the whole file as a slice of bytes
// instead of streaming it into smaller pieces
func Dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Exercise 1.4
// Modify dup2 to print the names of all files in which each duplicated line occurs
func countLinesIn(f *os.File, countedLines *countedLines) (filename string, err error) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		err := input.Err()
		if err != nil {
			return "", err
		}
		countedLines.counts[input.Text()]++
		countedLines.filename[input.Text()] = f.Name()
	}
	return f.Name(), nil
}

type countedLines struct {
	filename map[string]string
	counts   map[string]int
}

func newCountedLines() *countedLines {
	return &countedLines{
		filename: make(map[string]string),
		counts:   make(map[string]int),
	}
}

func Dup4() {
	countedLines := newCountedLines()
	files := os.Args[1:]

	if len(files) == 0 {
		err := countLines(os.Stdin, countedLines.counts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "DupInFile: %v\n", err)
		}
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "DupInFile: %v\n", err)
				continue
			}
			countLinesIn(f, countedLines)
			f.Close()
		}
	}

	for k, n := range countedLines.counts {
		if n > 1 {
			fmt.Printf("%s: %d\t%s\n", countedLines.filename[k], n, k)
		}
	}
}
