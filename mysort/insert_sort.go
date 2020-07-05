package mysort

func insertSortV1(ints []int) {
	count := len(ints)
	for i := 1; i < count; i++ {
		for j := i; j > 0; j-- {
			if ints[j-1] <= ints[j] {
				break
			}
			ints[j-1], ints[j] = ints[j], ints[j-1]
		}
	}
}

func insertSortV2(ints []int) {
	count := len(ints)
	for i := 1; i < count; i++ {
		bak := ints[i]
		j := i
		for ; j > 0; j-- {
			if ints[j-1] <= bak {
				break
			}
			ints[j] = ints[j-1]
		}
		ints[j] = bak
	}
}

func insertSortV3(ints []int) {
	count := len(ints)
	for i := 1; i < count; i++ {
		pos := findPosition(ints, i)
		insert(ints, i, pos)
	}
}

// 在完全随机的数组排序中优化提升有限
func insertSort(ints []int) {
	count := len(ints)
	for i := 1; i < count; i++ {
		if ints[i] >= ints[i-1] {
			continue
		}

		pos := findPosition(ints, i)
		insert(ints, i, pos)
	}
}

// 将 srcIndex 位置的元素插入到 dstIndex
func insert(ints []int, srcIndex, dstIndex int) {
	bak := ints[srcIndex]
	// [srcIndex, dstIndex)
	for i := srcIndex; i > dstIndex; i-- {
		ints[i] = ints[i-1]
	}

	ints[dstIndex] = bak
}

// [0, index) 已经排好序了
// 查找 ints[index] 元素应该放在什么位置
func findPosition(ints []int, index int) (pos int) {
	start, end := 0, index
	for start < end {
		mid := (start + end) >> 1
		if ints[mid] > ints[index] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	pos = start
	return
}
