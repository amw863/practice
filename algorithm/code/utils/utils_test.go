package utils

import (
	"fmt"
	"testing"
)

var (
	s = []int{1, 2, 3, 4}
)

func TestShuffle(t *testing.T) {
	s := []int{1, 2, 3, 4}
	fmt.Println(Shuffle(s))
	fmt.Println(ShuffleV2(s))
	fmt.Println(ShuffleV3(s))
}

func BenchmarkShuffleV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShuffleV3(s)
	}
}
func BenchmarkShuffleV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShuffleV2(s)
	}
}
func BenchmarkShuffle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Shuffle(s)
	}
}
