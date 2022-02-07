package sort

// BubbleSort 4, 5, 6, 1, 2, 3
// 只会操作相邻的元素，每次操作至少让一个元素移动到正确的位置
// 原地排序：
// 稳定排序：
// 时间复杂度：
func BubbleSort(s []int) []int {
	//for i := 0; i < len(s)-1; i++ {
	//	f := false
	//	for j := 0; j < len(s)-1; j++ {
	//		// 控制升序或者降序
	//		if s[j] > s[j+1] {
	//			s[j], s[j+1] = s[j+1], s[j]
	//			f = true
	//		}
	//	}
	//
	//	if f == false {
	//		break
	//	}
	//}

	for i := range s {
		for j := i; j < len(s); j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}

	return s
}

// InsertSort 在未排序的数据中取插入到已排序的数据中且保障已排序的顺序
// 2 5 6 3 4 1
// i = 3, t = 3, j = 2 s[j]=6 > 3, 256641
// j = 1, s[j]=5 > 3, 255641
// j = 0, s[j]=2 < 3,break, j = 0, s[j+1]=s[1]=3.235641
// 原地？稳定？？时间复杂度？？？
func InsertSort(s []int) []int {
	//for i := range s {
	//	t := s[i]
	//	j := i - 1
	//	for ; j >= 0; j-- {
	//		// 控制升序获取降序
	//		if s[j] < t {
	//			s[j+1] = s[j] // 对前面已排序的进行移动
	//		} else {
	//			break
	//		}
	//	}
	//
	//	s[j+1] = t
	//}

	for i := range s {
		t := s[i]
		j := i - 1
		for j >= 0 && s[j] > t {
			s[j+1] = s[j]
			j--
		}

		s[j+1] = t
	}
	return s
}

// SelectSort 从未排序中获取最小原始放到已排序末尾
// 原地??稳定??时间复杂度??
// 4 2 3 5 6 1
func SelectSort(s []int) []int {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	return s
}

// ShellSort 插入排序加分段
// 步长
func ShellSort(s []int) []int {
	for step := len(s) / 2; step >= 1; step = step / 2 {
		for i := step; i < len(s); i++ {
			t := s[i]
			j := i - step
			for ; j >= 0; j = j - step {
				if s[j] > t {
					s[j+step] = s[j]
				} else {
					break
				}
			}

			s[j+step] = t
		}
	}
	return s
}

// MergeSort 分治递归
func MergeSort(s []int) []int {
	if len(s) == 1 {
		return s
	}

	m := len(s) / 2

	return append()
}
