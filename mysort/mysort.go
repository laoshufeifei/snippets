package mysort

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"time"
)

// SortFunc ...
type SortFunc func(ints []int)

// TestBench ...
func TestBench() {
	N := 10 * 10000
	origins := make([]int, N)
	for i := 0; i < N; i++ {
		origins[i] = rand.Intn(N * 10)
	}

	funcs := []SortFunc{
		// 千万级别测试
		// quickSort,
		// mergeSort,
		// heapSort,
		shellSort,
		shellSortV1,

		// 以下测试最好不要超过 10 万
		// insertSort,
		// insertSortV3,
		// insertSortV2,
		// insertSortV1,
		// selectSort,

		// 以下测试最好不要超过 5 万
		// bubbleSort,
		// bubbleSortV1,
	}

	for _, fun := range funcs {
		ints := make([]int, N)
		copy(ints, origins)
		start := time.Now()
		fun(ints)
		fmt.Printf("%s \tcost = %v\n", getFunctionName(fun), time.Since(start))

		if !sort.IntsAreSorted(ints) {
			fmt.Println("error.................................")
		}
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
