package mysort

// 临时空间，提前分配好更快
var tmp []int

func mergeSort(ints []int) {
	tmp = make([]int, len(ints)>>1)
	mergeSortImpl(ints, 0, len(ints))
}

func mergeSortImpl(ints []int, start, end int) {
	if end-start < 2 {
		return
	}

	mid := (end + start) >> 1
	mergeSortImpl(ints, start, mid)
	mergeSortImpl(ints, mid, end)
	merge(ints, start, mid, end)
}

// [start, mid) [mid, end)
func merge(ints []int, start, mid, end int) {
	copy(tmp, ints[start:mid+1])

	// fmt.Printf("merge: [%d, %d) [%d, %d)\n", start, mid, mid, end)
	i, j := 0, mid
	index := start
	for i < (mid-start) && j < end {
		vi, vj := tmp[i], ints[j]
		if vi < vj {
			ints[index] = tmp[i]
			i++
		} else {
			ints[index] = ints[j]
			j++
		}

		index++
	}

	for i < (mid - start) {
		ints[index] = tmp[i]
		index++
		i++
	}

	// fmt.Println(ints)
}
