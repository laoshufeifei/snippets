package mysort

type insertSorter struct {
	Sorter
}

func newInsertSorterV1() *insertSorter {
	s := &insertSorter{}
	s.sortImple = s.insertSortV1
	return s
}

func newInsertSorterV2() *insertSorter {
	s := &insertSorter{}
	s.sortImple = s.insertSortV2
	return s
}

func newInsertSorterV3() *insertSorter {
	s := &insertSorter{}
	s.sortImple = s.insertSortV3
	return s
}

func (s *insertSorter) insertSortV1(ints []int) {
	s.array = ints
	count := len(ints)

	for i := 1; i < count; i++ {
		for j := i; j > 0; j-- {
			if s.CmpIndex(j-1, j) <= 0 {
				break
			}
			s.Swap(j-1, j)
		}
	}
}

func (s *insertSorter) insertSortV2(ints []int) {
	s.array = ints
	count := len(ints)

	for i := 1; i < count; i++ {
		bak := s.array[i]
		j := i
		for ; j > 0; j-- {
			if s.CmpValue(s.array[j-1], bak) <= 0 {
				break
			}
			ints[j] = ints[j-1]
		}
		ints[j] = bak
	}
}

// 使用二分查找优化
func (s *insertSorter) insertSortV3(ints []int) {
	s.array = ints
	count := len(ints)

	for i := 1; i < count; i++ {
		// 这句话优化有限可加可不加
		// if s.CmpIndex(i, i-1) >= 0 {
		// 	continue
		// }

		pos := s.findPosition(i)
		s.insert(i, pos)
	}
}

// 将 srcIndex 位置的元素插入到 dstIndex
func (s *insertSorter) insert(srcIndex, dstIndex int) {
	if srcIndex == dstIndex {
		return
	}

	bak := s.array[srcIndex]
	// [srcIndex, dstIndex)
	for i := srcIndex; i > dstIndex; i-- {
		s.array[i] = s.array[i-1]
	}

	s.array[dstIndex] = bak
}

// [0, index) 已经排好序了
// 查找 ints[index] 元素应该放在什么位置
func (s *insertSorter) findPosition(index int) (pos int) {
	start, end := 0, index
	for start < end {
		mid := (start + end) >> 1
		if s.CmpIndex(mid, index) > 0 {
			end = mid
		} else {
			start = mid + 1
		}
	}
	pos = start
	return
}
