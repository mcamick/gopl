// test to benchmark and compare for exercise 1_3

package join

import (
	"testing"
)

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join()
	}
}
