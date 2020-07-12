package mysort

type bubbleSorter struct {
	Sorter
}

func newBubbleSorterV1() *bubbleSorter {
	s := &bubbleSorter{}
	s.sortImple = s.sortV1
	return s
}

func newBubbleSorterV2() *bubbleSorter {
	s := &bubbleSorter{}
	s.sortImple = s.sortV2
	return s
}

func (s *bubbleSorter) sortV1(ints []int) {
	s.array = ints
	count := len(ints)
	for i := 0; i < count; i++ {
		// 使用 count-i 来实现要排序的范围
		for j := 1; j < count-i; j++ {
			if s.CmpIndex(j-1, j) > 0 {
				s.Swap(j-1, j)
			}
		}
	}
}

func (s *bubbleSorter) sortV2(ints []int) {
	s.array = ints
	count := len(ints)

	for end := count - 1; end >= 0; end-- {
		sortedIndex := end
		for j := 1; j <= end; j++ {
			if s.CmpIndex(j-1, j) >= 0 {
				s.Swap(j-1, j)
				sortedIndex = j
			}
		}

		end = sortedIndex
	}
}
