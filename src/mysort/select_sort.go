package mysort

type selectSorter struct {
	Sorter
}

func newSelectSorter() *selectSorter {
	s := &selectSorter{}
	s.sortImple = s.selectSort
	return s
}

func (s *selectSorter) selectSort(ints []int) {
	s.array = ints

	count := len(ints)
	for i := 0; i < count; i++ {
		minIdx := s.selectMinIndex(i, count)
		s.Swap(i, minIdx)
	}
}

// [start, end)
func (s *selectSorter) selectMinIndex(start, end int) int {
	minIndex, minValue := start, s.array[start]
	for i := start + 1; i < end; i++ {
		if s.CmpValue(s.array[i], minValue) < 0 {
			minIndex, minValue = i, s.array[i]
		}
	}

	return minIndex
}
