// Modify dup2 to print the names of all files in which each duplicated line occurs.
// Used command  "go run exercise1_4.go example1.txt example2.txt example3.txt" to invoke and run against three files

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]string)
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

			// Reset the file pointer to the beginning of the file
			_, err = f.Seek(0, 0)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %s\n", err)
				f.Close()
				continue // go to next iteration of the loop
			}
			determineNames(f, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%-30s\t%s\n", n, line, filenames[line])
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

func determineNames(f *os.File, filenames map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		filename := f.Name()

		// Check if the line is already present in the filenames map
		if filenamesValue, ok := filenames[line]; ok {
			// If the filename is not already in the value associated with the line, append it
			if !strings.Contains(filenamesValue, filename) {
				filenames[line] = filenamesValue + ", " + filename
			}
		} else {
			// If the line is not present in the map, add it with the filename
			filenames[line] = filename
		}
	}
}
