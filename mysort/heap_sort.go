package mysort

// 大顶堆
type heap struct {
	ints []int
	size int
}

func newHeap(ints []int) *heap {
	size := len(ints)
	h := &heap{
		// 拷贝个结构体代价不大
		ints: ints,
		size: size,
	}

	index := size>>1 - 1
	for i := index; i >= 0; i-- {
		h.shifDown(i)
	}

	return h
}

// 大顶堆
func (h *heap) shifDown(index int) {
	leftIndex := index*2 + 1
	for leftIndex < h.size {
		value := h.ints[index]

		leftValue := h.ints[leftIndex]
		biggerIndex, biggerValue := leftIndex, leftValue

		rightIndex := leftIndex + 1
		if rightIndex < h.size {
			rightValue := h.ints[rightIndex]
			if rightValue > leftValue {
				biggerIndex, biggerValue = rightIndex, rightValue
			}
		}

		// swap index, biggerIndex
		if biggerValue > value {
			h.ints[index], h.ints[biggerIndex] = h.ints[biggerIndex], h.ints[index]
		}

		index = biggerIndex
		leftIndex = index*2 + 1
	}
}

func (h *heap) pop() int {
	if h.size == 0 {
		return 0
	}

	header := h.ints[0]
	h.size--

	h.ints[0] = h.ints[h.size]
	h.shifDown(0)
	return header
}

func heapSort(ints []int) {
	h := newHeap(ints)

	count := len(ints)
	for index := 1; index <= count; index++ {
		ints[count-index] = h.pop()
		// fmt.Println(ints)
	}
}
