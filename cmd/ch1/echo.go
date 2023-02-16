package ch1

import (
	"fmt"
	"os"
	"strings"
)

func Echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep += " "
	}
	fmt.Println(s)
}

// Note: avoid using += Join has better performance
func Echo2() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += arg
		sep += " "
	}
	fmt.Println(s)
}

func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// Exercise 1.1
// Modify the echo program to also print os.Args[0] the name of the command
// that invoked it
func EchoZero() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

// Exercise 1.2
// Modify the echo program to print the index and value of each
// of its arguments, one per line
func EchoWIndex() {
	for idx, val := range os.Args {
		fmt.Printf("[%d] %s\n", idx, val)
	}
}

// TODO ex 1.3
