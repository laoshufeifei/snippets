package mysort

func bubbleSortV1(ints []int) {
	count := len(ints)
	for i := 0; i < count; i++ {
		// 使用 count-i 来实现要排序的范围
		for j := 1; j < count-i; j++ {
			if ints[j-1] > ints[j] {
				ints[j-1], ints[j] = ints[j], ints[j-1]
			}
		}
	}
}

func bubbleSort(ints []int) {
	count := len(ints)

	for end := count - 1; end >= 0; end-- {
		sortedIndex := end
		for j := 1; j <= end; j++ {
			if ints[j-1] > ints[j] {
				ints[j-1], ints[j] = ints[j], ints[j-1]
				sortedIndex = j
			}
		}

		end = sortedIndex
	}
}
