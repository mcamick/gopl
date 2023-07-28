package math

import (
	"fmt"
	"testing"
)

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{10, 20, 30},
	{0, 0, 0},
	{-1, 1, 0},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Error("Test Failed: {} {} inputted, {} expected, recieved: {}", test.arg1, test.arg2, test.expected, output)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 5)
	}
}

func ExampleAdd() {
	fmt.Println(Add(4, 6))
	// Output: 10
}
