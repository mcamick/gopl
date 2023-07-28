// Program prints its command-line arguments.
// Exercise 1.2 asks that we modify the program to print the index and value of each of the command line arguments it has
package exercise1_2

import (
	"fmt"
	"os"
)

func Echo2() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Print(i, " ")
		fmt.Println(os.Args[i])
	}
}
