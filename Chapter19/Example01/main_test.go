package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func add(x, y int) int {
	return x + y
}

func main() {
	result := add(1, 2)
	fmt.Printf("addition result: %v\n", result)
}

func FuzzAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int, j int) {
		got := add(i, j)
		assert.Equal(t, i+j, got, "the values i and j should be summed together properly")
	})
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
}
