// Echo1 prints its command-line arguments.
// Exercise 1.1 asks that we print os.args[0] to print the command that invoked it
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Print(i, " ")
		fmt.Println(os.Args[i])
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)
}
