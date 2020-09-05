package mysort

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// BenchmarkTestSorters ...
// TODO: convert this test to benchmark test
func BenchmarkTestSorters() {
	rand.Seed(time.Now().UnixNano())

	N := 10 * 10000
	origins := make([]int, N)
	for i := 0; i < N; i++ {
		origins[i] = rand.Intn(N * 10)
	}

	handlers := []SortHandler{
		// 千万级别测试
		newQuickSorter(),
		newMergeSorter(),
		newHeapSorter(),
		newShellSorterV1(),
		newShellSorterV2(),

		// 以下测试最好不要超过 10 万
		newInsertSorterV1(),
		newInsertSorterV2(),
		newInsertSorterV3(),
		// newSelectSorter(),

		// 以下测试最好不要超过 5 万
		// newBubbleSorterV1(),
		// newBubbleSorterV2(),
	}

	for _, handler := range handlers {
		ints := make([]int, N)
		copy(ints, origins)

		handler.StartSort(ints)
		if !sort.IntsAreSorted(ints) {
			fmt.Println("error.................................")
		}
	}
}
