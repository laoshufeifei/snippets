package mysort

type mergeSorter struct {
	Sorter
	// 临时空间，提前分配好更快
	tmp []int
}

func newMergeSorter() *mergeSorter {
	s := &mergeSorter{}
	s.sortImple = s.mergeSort
	return s
}

func (s *mergeSorter) mergeSort(ints []int) {
	s.array = ints
	s.tmp = make([]int, len(ints)>>1)

	s.mergeSortImpl(0, len(ints))
}

func (s *mergeSorter) mergeSortImpl(start, end int) {
	if end-start < 2 {
		return
	}

	mid := (end + start) >> 1
	s.mergeSortImpl(start, mid)
	s.mergeSortImpl(mid, end)
	s.merge(start, mid, end)
}

// [start, mid) [mid, end)
func (s *mergeSorter) merge(start, mid, end int) {
	copy(s.tmp, s.array[start:mid+1])

	// fmt.Printf("merge: [%d, %d) [%d, %d)\n", start, mid, mid, end)
	i, j := 0, mid
	index := start
	for i < (mid-start) && j < end {
		if s.CmpValue(s.tmp[i], s.array[j]) < 0 {
			s.array[index] = s.tmp[i]
			i++
		} else {
			s.array[index] = s.array[j]
			j++
		}

		index++
	}

	for i < (mid - start) {
		s.array[index] = s.tmp[i]
		index++
		i++
	}
}
