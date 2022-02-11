// Echo1 prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

// will print all command line arguments passed when invoking the program
func main() {

	var s, sep string // delcares two variables of type string
	// the variables in the line above are implicitly initialized to their zero values

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		// equivalent to s = s + sep + os.Args[i]
		sep = " "
	}
	
	/*	
	 *	fmt.Println(reflect.TypeOf(os.Args))
	 *	as os.Args has type of string array, sep + os.Args[i] will concatenate sep 
	 *	with the iesm element of os.Args[i]
	 */

	fmt.Println(s)

	/* 
	 *	The string s starts its lifecycle empty, that is, with value "" 
	 * 	after that, with each trip trough the loop, some text is added to it
	 * 	This is a quadritic process and can be inneficient if the number of arguments gets too big
	 */
}
