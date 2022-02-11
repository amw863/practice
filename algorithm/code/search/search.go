package search

/**
  二分查找是对有序的集合, 类似分治的思
    要点：
        1. 二分查找依赖线性表，也即数组
        2. 数据要有序，所以适应一次排序多次查找，也即增删不频繁的场景
        3. 数据量小不适合，但是如果比较耗时则推荐二分查找
        4. 数据量大不适合，因为数组线性的要求内存连续
    注意：
        1. mid取值
        2. low 和 high 变更
        3. 循环退出

	二分变形：
		- 第一个等于给定的
		- 最后一个等于给定的
		- 第一个大于等于给定的
		- 最后一个小于等于给定的
*/

import (
	"fmt"
)

func bsearchV3(s []int, v int) bool {
	l := 0
	h := len(s) - 1

	for l <= h {
		m := l + (h-l)>>1
		if s[m] == v {
			return true
		} else if s[m] > v {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}

// 递归方式和非递归方式一致
func bsearch(s []int, v int) bool {
	return bs(v, s)
	//  return  search(s, v, 0, len(s) -1)
}

func search(s []int, v, l, r int) bool {
	if l > r {
		return false
	}

	m := l + (r-l)>>1
	fmt.Printf("l:%d, r:%d, m:%d\n", l, r, m)
	if s[m] == v {
		return true
	} else if s[m] > v {
		return search(s, v, l, m-1)
	} else {
		return search(s, v, m+1, r)
	}
}

func bsearchV2(s []int, v int) bool {
	low := 0
	high := len(s) - 1

	// 退出循环的条
	// low <= high ??? 什么时候 low > high
	for low <= high {
		m := low + (high-low)>>1
		fmt.Printf("【测试 %d】h:%d, l:%d m:%d \n", v, high, low, m)
		if s[m] == v {
			return true
		} else if s[m] > v {
			high = m - 1
		} else {
			// 此处会导致low > high 会退出循环
			low = m + 1
		}
	}

	return false
}

// 扩展：转成整数
func sqrt(s float64) float64 {
	t := int64(s * 1000000)
	var l int64 = 0
	h := t
	var m int64
	for l <= h {
		m = (l + h) >> 1
		fmt.Printf("t:%v, m2:%v, l:%v, h:%v, m:%v\n", t, m*m, l, h, m)
		if m*m == t {
			return s
		} else if m*m > t {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return float64(m) / 1000
}

func SearchV3(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	// target = 8
	// 4, 5, 6, 7, 8, 1, 2, 3
	// 7, 8,  1, 2, 3, 4, 5, 6
	for l <= r {
		m := l + (r-l)>>1
		fmt.Printf("l:%d, m:%d, r:%d\n", l, m, r)
		if nums[l] > nums[m] { // m 左边有序
			if target {
				// todo
			}
		} else { //

		}
		//if nums[m] == target {
		//	return m
		//} else if nums[l] > nums[r] {
		//	if target > nums[m] && target >= nums[l] {
		//		r = m - 1
		//	} else if target < nums[m] && target <= nums[r] {
		//		l = m + 1
		//	}
		//} else {
		//	if nums[m] > target {
		//		r = m - 1
		//	} else {
		//		l = m + 1
		//	}
		//}
	}

	return -1
}
