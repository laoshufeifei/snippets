package mysort

// 大顶堆
type heapSorter struct {
	Sorter
	heapSize int // 堆的 size (也就是未排序的个数)
}

func newHeapSorter() *heapSorter {
	s := &heapSorter{}
	s.sortImple = s.heapSort
	return s
}

func (h *heapSorter) heapSort(ints []int) {
	// h := newHeap(ints)
	h.array = ints
	size := len(ints)
	h.heapSize = size

	// 自底而上的下溢
	index := size>>1 - 1
	for i := index; i >= 0; i-- {
		h.shifDown(i)
	}

	for h.heapSize > 0 {
		h.Swap(0, h.heapSize-1)
		h.heapSize--
		h.shifDown(0)
	}
}

// 大顶堆
// 如果使用几个中间变量把 value 值保存下来，然后比较 value，会快一点点
func (h *heapSorter) shifDown(index int) {
	leftIndex := index*2 + 1
	rightIndex, biggerIndex := 0, 0
	for leftIndex < h.heapSize {
		biggerIndex = leftIndex

		rightIndex = leftIndex + 1
		if rightIndex < h.heapSize {
			if h.CmpIndex(rightIndex, leftIndex) > 0 {
				biggerIndex = rightIndex
			}
		}

		// swap index, biggerIndex
		if h.CmpIndex(biggerIndex, index) > 0 {
			h.Swap(biggerIndex, index)
		}

		index = biggerIndex
		leftIndex = index*2 + 1
	}
}
