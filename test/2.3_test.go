//go test -bench=.
package test

import (
	"testing"
	"awesomeProject/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

func BenchmarkPopCountFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}

func BenchmarkPopCount25(b *testing.B)  {
	for i :=0; i < b.N; i++ {
		popcount.PopCountr25(uint64(i))
	}
}