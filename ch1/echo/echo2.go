package main

import (
	"fmt"
	"os"
)

func main() {
	
	s, sep := "", ""
	// _ is what is called a blank identifier
	// it is used whenever syntax requires a variable name but logic does not
	for _, arg := range os.Args[1:] {
		s += sep + arg 
		/* trough each iteration of the loop, s gets a whole new value
		 * the old value is garbage collected
		 * this could be costly, sou the usage of Join from the strings package could be
		 * more efficient
		 */
		sep = " "
	}

	fmt.Println(s)
}
