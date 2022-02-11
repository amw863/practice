package search

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	assert.Equal(t, math.Sqrt(30), sqrt(30))
}

func TestSearchV3(t *testing.T) {
	assert.Equal(t, -1, SearchV3([]int{1}, 2))
	assert.Equal(t, 5, SearchV3([]int{4, 5, 6, 0, 1, 2, 3}, 2))
	assert.Equal(t, 4, SearchV3([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	assert.Equal(t, 0, SearchV3([]int{3, 5, 1}, 3))
	assert.Equal(t, 4, SearchV3([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8))
}

func TestBSearch(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if bsearch(s, 0) == false {
		t.Fatal("【0】 测试失败")
	}

	if bsearch(s, 9) == false {
		t.Fatal(" 【9】 测试失败")
	}

	if bsearch(s, 3) == false {
		t.Fatal("【3】 测试失败")
	}

	if bsearch(s, 10) {
		t.Fatal("【10】 测试失败")
	}
}
