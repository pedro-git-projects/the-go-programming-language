/*
 *	Exercise 1.1: Modify the echo prog ram to als o print os.Args[0],
 * 	the name of the command that invoked it.
 */

package main

import (
	"fmt"
	"os"
)

func main() {

	var s, sep string

	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
