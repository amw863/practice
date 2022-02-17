package sort

func ISort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {

	}

	return s
}

func MSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	m := len(s) >> 1
	return _merge(MSort(s[:m]), MSort(s[m:]))
}

func _merge(s1, s2 []int) []int {
	s := make([]int, len(s1)+len(s2))
	i, j := 0, 0
	for {
		if i < len(s1) && (j == len(s2) || s1[i] < s2[j]) {
			s[i+j] = s1[i]
			i++
		} else if j < len(s2) {
			s[i+j] = s2[j]
			j++
		} else {
			break
		}
	}

	return s
}

// 找出一个中间值，把大的小的分开
func PQSort(s []int) []int {
	return _qsort(s, 0, len(s)-1)
}

func _qsort(s []int, l, r int) []int {
	if l < r {
		m := _partition(s, l, r)
		_qsort(s, l, m-1)
		_qsort(s, m+1, r)
	}

	return s
}

func _partition(s []int, l, r int) int {
	p := l
	index := p + 1

	for i := index; i <= r; i++ {
		if s[p] > s[i] {
			if i != index {
				s[i], s[index] = s[index], s[i]
			}

			index++
		}
	}

	s[index-1], s[p] = s[p], s[index-1]

	return index - 1
}
