package code

import "testing"

func TestChain(t *testing.T) {
	ch := &WordsFilterChain{}
	a := A{}
	b := B{}
	ch.AddFilter(a)
	ch.AddFilter(b)
	ch.Check()
}
