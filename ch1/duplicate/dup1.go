package main

import (
	"bufio"
	"fmt"
	"os"
)

// prints each line that appears more than once preceded by its count
func main() {

	// make creates a new empty map where the keys are strings and the values integers
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	// ignoring potential errors from input.Err()
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
	
}
