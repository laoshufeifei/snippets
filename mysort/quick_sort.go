package mysort

type quickSorter struct {
	Sorter
}

func newQuickSorter() *quickSorter {
	s := &quickSorter{}
	s.sortImple = s.quickSort
	return s
}

func (s *quickSorter) quickSort(ints []int) {
	s.array = ints
	s.quickSortImple(0, len(ints))
}

// [start, end)
func (s *quickSorter) quickSortImple(start, end int) {
	if end-start < 2 {
		return
	}
	pivot := s.partition(start, end)
	s.quickSortImple(start, pivot)
	s.quickSortImple(pivot+1, end)
}

func (s *quickSorter) partition(start, end int) int {
	// fmt.Printf("[%d, %d)\n", start, end)
	// 也可以随机选择一个元素与 start 交换，防止出现逆序这种情况
	i, j := start, end-1
	bak := s.array[i]
	for i < j {
		for i < j {
			if s.CmpValue(s.array[j], bak) > 0 {
				j--
			} else {
				s.array[i] = s.array[j]
				i++ // 不要忘记
				break
			}
		}

		for i < j {
			// if ints[i] < bak {
			if s.CmpValue(s.array[i], bak) < 0 {
				i++
			} else {
				s.array[j] = s.array[i]
				j-- // 不要忘记
				break
			}
		}
	}
	s.array[i] = bak
	return i
}

// https://www.runoob.com/w3cnote/quick-sort-2.html
// 另一种思路实现的，比上面稍微慢一点点，因为交换的次数会比较多
func (s *quickSorter) partition2(start, end int) int {
	pivot := start
	index := start + 1

	for i := index; i < end; i++ {
		if s.CmpIndex(i, pivot) < 0 {
			if i != index {
				s.Swap(i, index)
			}
			index++
		}
	}
	s.Swap(pivot, index-1)
	return index - 1
}
