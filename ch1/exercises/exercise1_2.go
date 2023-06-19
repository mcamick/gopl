// Program prints its command-line arguments.
// Exercise 1.2 asks that we modify the program to print the index and value of each of the command line arguements it has
package main

import (
	"fmt"
	"os"
)

func exercise1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Print(i, " ")
		fmt.Println(os.Args[i])
	}
	fmt.Println(s)
}
