/*
Exercise 1.3 wants to measure the run time variance between strings.join and the
for loop that iterates over the arguments presented in echo 1

This link was much better at explaining testing in go than the book:
https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package
*/

package exercise1_1

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1()
	}
}
