package sort

// 链表的插入冒泡选择
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
		for j := i; j < len(s)-1; j++ {
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
		// 挪牌
		for j >= 0 && s[j] > t {
			s[j+1] = s[j]
			j--
		}

		// 插入，这里为什么是j+1因为这个时候j已经减到位置了
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

// MergeSort 递归算法 vs 非递归算法
// merge(s) = merge_sort(merge(s[0:q]), merge(s[q:]))
// 6 2 4 5 3 1
// 624 531
// 6 24
// 2 4
// 稳定：相同元素的先后顺序
// 时间复杂度
// 空间复杂度（是否原地，是否借用额外空间）:
func MergeSort(s []int) []int {
	l := len(s)
	if l <= 1 {
		return s
	}
	m := l / 2

	// 分割到底之后应用 merge?? 为什么这样就快呢？
	return merge(MergeSort(s[:m]), MergeSort(s[m:]))
}

func merge(s1, s2 []int) []int {
	t := make([]int, len(s1)+len(s2))
	for i, j := 0, 0; ; {
		// 主要理解这一段, 创建个临时的切片，
		// 前半部分0~i
		// 后半部分0~j
		// i 合理 当后半部分没有值获取 前半部分的值小于后半部分时,
		// 妙啊
		if i < len(s1) && (j == len(s2) || s1[i] < s2[j]) {
			t[i+j] = s1[i]
			i++
			// 当前半部分没有值，后半部分还有值时，把后半部分值添加到临时切片
		} else if j < len(s2) {
			t[i+j] = s2[j]
			j++
		} else {
			break
		}
	}
	return t
}


// QuickSort
// 数组中任意找一个分割点, 大于分割点的放左侧，小于分割点的放右侧
// 到最后一个停止
func QuickSort(s []int) []int {
	return quickSort(s, 0, len(s)-1)
}

func quickSort(s []int, left, right int) []int {
    if (left > right) {
        return s
    }

    m := partition(s, left, right)
    quickSort(s, left, m-1)
    quickSort(s, m+1, right)
    return s
}

func partition(s []int, left, right int) int {
    // 选首位置和最后一个位置
    pivot := left
    // 记录比pivot小的位置
    index := pivot + 1

    for i:=index;i<=right;i++ {
        if s[i] < s[pivot] {
            if(i != index) {
                // 当某个值比pivot小，则把改值和index替换，此时index++
                // 此处妙在index值也替换了, 记录位置也增加了
                s[i], s[index] = s[index], s[i]
            }
            index++
        }
    }

    // 此处为什么再替换一次, index - 1记录的是比pivot小的值
    // 所以中间值为此时进行了替换所以是index-1
    s[pivot], s[index -1] = s[index -1], s[pivot]
    
    return index -1
}



// 看懂不
// 堆排序
// 应用场景第K大数
// 优先级队列
func HeapSort(s []int) []int {
    l := len(s)
    buildMaxHeap(s, l)
    for i := l-1;i>=0;i-- {
        s[i], s[0] = s[0], s[i]
        heapify(s, 0, l)
    }

    return s
}

func buildMaxHeap(s []int, l int) {
    for i := l/2;i>=0;i++ {
        heapify(s, i,l)
    }
}

func heapify(s []int, i, l int) {
    left := 2*i+1
    right := 2*i+2

    largest := i

    if left < l && s[left] > s[largest] {
        largest = left
    }

    if right < l && s[right] > s[largest] {
        largest = right
    }

    if largest != i {
        s[largest], s[i] = s[i], s[largest]
        heapify(s, largest, l)
    }
}


// 没看懂
// Q1: 求无序数组第K大元素要求O(n)的时间复杂度
// Q2: 10个300M的日志文件，文件中日志按照时间戳有序,希望合成一个有序。如果内存有限怎么处理?
// 10 个句柄 内存大小无所谓了


// BucketSort 如果给100w用户的年龄进行排序
// 思路：将排序的数据放进有序的桶里，再对桶里进行排序【分桶+插入或者快排+归并】
// 要求比较严格：
//	1. 分桶比较简单
//  2. 内桶比较均匀
//  3. 外排序
func BucketSort(s []int) []int {
	return s
}

// CountingSort
// 思想：分桶排序的一种特殊情况，当分通的个数有限，全部分桶，省略桶内排序的耗时
// 比如：高考排名 可以分0~900个桶
// 前半部分还能看懂啥意思。后面的就看不懂了
// 找出待排序中最大和最小值
// 统计每个数出现的次数
// 对所有的计次累加
// 反向填充目标数据
// 比较耗内存，不适合字母排序
func CountingSort(s []int) []int {
    var maxVal, minVal int
    // 找出最大最小值
    for i:=0;i<len(s);i++ {
        if(s[i] > maxVal) {
            maxVal = s[i]
        }

        if(s[i] < minVal) {
            minVal = s[i]
        }
    }


    s1 := make([]int, maxVal - minVal+1)
    // 统计每个值出现的次数
    for i:=0;i<len(s);i++ {
        s1[s[i]]++
    }
//println(s1)
    s2 := make([]int, 0, len(s))
    for k, v := range s1 {
        for i:=0;i<v;i++ {
            s2 = append(s2, k)
        }
    }

    return s2
}

// RadixSort
// 数据比较特殊：如手机号排序，如果 a 的前面比 b 的前面几的位大则 后面的就不用比较了
// 感觉是前缀排序
// 本质是另一种分桶方式
func RadixSort() {

}

// 总结：核心的快排+归并+插入（希尔）。桶、计数、基数比较特殊
