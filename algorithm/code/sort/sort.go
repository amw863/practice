package sort

// 4, 5, 6, 1, 2, 3
// 只会操作相邻的元素，每次操作至少让一个元素移动到正确的位置
// 原地排序：
// 稳定排序：
// 时间复杂度：
func BubbleSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		f := false
		for j := 0; j < len(s)-1; j++ {
			// 控制升序或者降序
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				f = true
			}
		}

		if f == false {
			break
		}
	}

	return s
}

// 在未排序的数据中取插入到已排序的数据中且保障已排序的顺序
// 2 5 6 3 4 1
// i = 3, t = 3, j = 2 s[j]=6 > 3, 256641
// j = 1, s[j]=5 > 3, 255641
// j = 0, s[j]=2 < 3,break, j = 0, s[j+1]=s[1]=3.235641
// 原地？稳定？？时间复杂度？？？
func InsertSort(s []int) []int {
	for i := range s {
		t := s[i]
		j := i - 1
		for ; j >= 0; j-- {
			// 控制升序获取降序
			if s[j] < t {
				s[j+1] = s[j] // 对前面已排序的进行移动
			} else {
				break
			}
		}

		s[j+1] = t
	}
	return s
}

// 从未排序中获取最小原始放到已排序末尾
// 原地??稳定??时间复杂度??
func SelectSort(s []int) []int {
	for i := range s {
		t := s[i]
		for
	}
	return s
}
