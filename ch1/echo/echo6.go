// Exercis e 1.2: Mo dif y the echo prog ram to print the index and value of each of its arguments, on e per line.
// same exercise, but this time I'm using join

package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	i := 0

	for _, arg := range os.Args[0:] {
		s = arg
		fmt.Println(i, s)		
		i++
	}
}
