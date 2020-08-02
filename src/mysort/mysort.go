package mysort

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// SortHandler ...
type SortHandler interface {
	StartSort(ints []int)
}

// SortFunc ...
type SortFunc func(ints []int)

// Sorter 排序基类
// 只想对比下几种排序算法，还没想要在实际生产中使用，所以只支持 int 类型
type Sorter struct {
	array       []int
	swapCounter int
	cmpCounter  int
	sortImple   SortFunc
}

// Swap swap with index
func (s *Sorter) Swap(i, j int) {
	s.swapCounter++
	s.array[i], s.array[j] = s.array[j], s.array[i]
}

// CmpIndex compare with index
func (s *Sorter) CmpIndex(i, j int) int {
	s.cmpCounter++
	switch {
	case s.array[i] < s.array[j]:
		return -1
	case s.array[i] > s.array[j]:
		return 1
	default:
		return 0
	}
}

// CmpValue compare with value
func (s *Sorter) CmpValue(v1, v2 int) int {
	s.cmpCounter++
	switch {
	case v1 < v2:
		return -1
	case v1 > v2:
		return 1
	default:
		return 0
	}
}

// StartSort ...
func (s *Sorter) StartSort(ints []int) {
	startTime := time.Now()
	s.sortImple(ints)

	fmt.Printf("%s cost %v, swap: %s, compare: %s\n",
		getFunctionName(s.sortImple), time.Since(startTime), formatCounter(s.swapCounter), formatCounter(s.cmpCounter))
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func formatCounter(counter int) string {
	switch {
	case counter < 10000:
		return fmt.Sprintf("%d", counter)
	default:
		return fmt.Sprintf("%f w", float64(counter)/10000.)
	}
}
