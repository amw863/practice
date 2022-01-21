package code

import "testing"

func TestStrategy(t *testing.T) {
	NewBook(59.90, NovelStrategy{}).Print()
}
