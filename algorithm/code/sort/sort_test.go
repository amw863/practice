package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var s = []int{4, 5, 6, 1, 2, 3}

func TestBubbleSortSort(t *testing.T) { assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, InsertSort(s)) }

func TestInsertSort(t *testing.T) { assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, BubbleSort(s)) }

func TestSelectSort(t *testing.T) { assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, SelectSort(s)) }

func TestShellSort(t *testing.T) { assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, ShellSort(s)) }
