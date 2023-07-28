// Echo1 prints its command-line arguments.
// Exercise 1.1 asks that we print os.args[0] to print the command that invoked it
package exercise1_1

import (
	"fmt"
	"os"
)

func Echo1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Print(i, " ")
		fmt.Println(os.Args[i])
	}
}
