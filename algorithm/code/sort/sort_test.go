package sort

import (
	"demo1/practice/algorithm/code/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	given  = []int{}
	wanted = []int{}
)

type data struct {
	given    []int
	expected []int
}

func TestSort(t *testing.T) {
	expected := utils.NewSlice(4)
	give := utils.Shuffle(expected)

	testCase := []data{
		{given: give, expected: expected},
	}

	for _, d := range testCase {
		fmt.Println(d)
		assert.Equal(t, d.expected, InsertSort(d.given))
		assert.Equal(t, d.expected, BubbleSort(d.given))
		assert.Equal(t, d.expected, SelectSort(d.given))
		assert.Equal(t, d.expected, ShellSort(d.given))
		assert.Equal(t, d.expected, MergeSort(d.given))
	}
}

func TestInsertSort(t *testing.T) { assert.Equal(t, wanted, InsertSort(given)) }

func TestBubbleSortSort(t *testing.T) { assert.Equal(t, wanted, BubbleSort(given)) }

func TestSelectSort(t *testing.T) { assert.Equal(t, wanted, SelectSort(given)) }

func TestShellSort(t *testing.T) { assert.Equal(t, wanted, ShellSort(given)) }
func TestMergeSort(t *testing.T) { assert.Equal(t, wanted, MergeSort(given)) }

var s = utils.Shuffle(utils.NewSlice(10000))

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort(s)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(s)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort(s)
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShellSort(s)
	}
}

// go 参数是传值
// 但是 map slice channel 指针 引用类型 巨坑!!!!!!
// slice 作为参数或者返回函数都会影响外部
// 注意这个坑
// 另一个坑是越界
func TestSlice(t *testing.T) {
	s := []int{1, 2, 3}
	v := sliceHandler(s)
	v[1] = 9
	fmt.Println(s)
	fmt.Println(v)
}

func sliceHandler(s []int) []int {
	// 此处要判断长度，否则会溢出
	s[1] = 1000

	// 直接返回也会影响外部的
	//return s
	b := make([]int, len(s))
	copy(b, s)
	return b
}

func TestM(t *testing.T) {
	m := make(map[string]string)
	m["a"] = "aa"
	c := mapHandler(m)
	c["a"] = "ccc"
	fmt.Println(m)
}

func mapHandler(m map[string]string) map[string]string {
	t := make(map[string]string)
	m["a"] = "bbb"
	t["t"] = m["a"]
	return t
}
