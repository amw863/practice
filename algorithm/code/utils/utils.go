package utils

import (
	"math/rand"
	"time"
)

func ShuffleV2(s []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
}

func ShuffleV3(s []int) []int {
	l := len(s)
	for l > 0 {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(l)
		s[l-1], s[randNum] = s[randNum], s[l-1]
		l -= 1
	}
	return s
}

func Shuffle(s []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := len(s)
	for l > 0 {
		randNum := r.Intn(l)
		s[l-1], s[randNum] = s[randNum], s[l-1]
		l -= 1
	}
	return s
}

func RandSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Intn(n)
	}

	return s
}

func NewSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}
	return s
}
