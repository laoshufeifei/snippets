package mysort

func quickSort(ints []int) {
	quickSortImple(ints, 0, len(ints))
}

// [start, end)
func quickSortImple(ints []int, start, end int) {
	if start+1 >= end {
		return
	}
	pivot := partition(ints, start, end)
	quickSortImple(ints, start, pivot)
	quickSortImple(ints, pivot+1, end)
}

func partition(ints []int, start, end int) int {
	// fmt.Printf("[%d, %d)\n", start, end)
	i, j := start, end-1
	bak := ints[i]
	for i < j {
		for i < j {
			if ints[j] > bak {
				j--
			} else {
				ints[i] = ints[j]
				i++ // 不要忘记
				break
			}
		}

		for i < j {
			if ints[i] < bak {
				i++
			} else {
				ints[j] = ints[i]
				j-- // 不要忘记
				break
			}
		}
	}
	ints[i] = bak
	return i
}

// https://www.runoob.com/w3cnote/quick-sort-2.html
// 另一种思路实现的，比上面稍微慢一点点，因为交换的次数会比较多
func partition2(ints []int, start, end int) int {
	pivot := start
	index := start + 1

	for i := index; i < end; i++ {
		if ints[i] < ints[pivot] {
			if i != index {
				ints[i], ints[index] = ints[index], ints[i]
			}
			index++
		}
	}
	ints[pivot], ints[index-1] = ints[index-1], ints[pivot]
	return index - 1
}
