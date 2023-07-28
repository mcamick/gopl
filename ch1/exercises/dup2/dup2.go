// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %s\n", err)
				continue // go to next iteration of the loop
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 { // print only lines that appear more than once
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) { // f is a pointer to an os.File
	input := bufio.NewScanner(f) // input is a pointer to a Scanner
	for input.Scan() {           // input.Scan() returns true if there is a line and false if there is none
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
