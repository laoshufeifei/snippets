package mysort

func selectSort(ints []int) {
	count := len(ints)
	for i := 0; i < count; i++ {
		minIdx := selectMinIndex(ints, i, count)
		ints[i], ints[minIdx] = ints[minIdx], ints[i]
	}
}

func selectMinIndex(ints []int, start, end int) int {
	if start >= end {
		return -1
	}

	minIndex, minValue := start, ints[start]
	for i := start + 1; i < end; i++ {
		if ints[i] < minValue {
			minIndex, minValue = i, ints[i]
		}
	}

	return minIndex
}
