/*
Exercise 1.3 wants to measure the run time variance between strings.join and the
for loop that iterates over the arguments presented in echo 1
*/

package exercise1_2

import (
	"testing"
)

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2()
	}
}
