package mysort

import (
	"math"
)

type shellSorter struct {
	Sorter
}

func newShellSorterV1() *shellSorter {
	s := &shellSorter{}
	s.sortImple = s.shellSortV1
	return s
}

func newShellSorterV2() *shellSorter {
	s := &shellSorter{}
	s.sortImple = s.shellSortV2
	return s
}

func (s *shellSorter) shellSortV1(ints []int) {
	s.array = ints

	steps := s.shellStep(len(ints))
	for _, step := range steps {
		s.shellImple(step)
	}
}

func (s *shellSorter) shellSortV2(ints []int) {
	s.array = ints

	steps := s.sedgewickStep(len(ints))
	for i := len(steps) - 1; i >= 0; i-- {
		s.shellImple(steps[i])
	}
}

func (s *shellSorter) shellImple(step int) {
	count := len(s.array)
	for col := 0; col < step; col++ {
		// col+step*0, col+step*1, col+step*2
		for begin := col + step; begin < count; begin += step {
			cur := begin
			for cur > col {
				pre := cur - step
				if s.CmpIndex(cur, pre) >= 0 {
					break
				}

				s.Swap(cur, pre)
				cur -= step
			}
		}
	}
}

// 生成的序列是从大到小的
func (s *shellSorter) shellStep(count int) (steps []int) {
	step := count >> 1
	for step > 0 {
		steps = append(steps, step)
		step = step >> 1
	}
	return
}

// 生成的序列是从小到大的，使用的时候要倒序使用
// 1, 5, 19, 41, 109...
func (s *shellSorter) sedgewickStep(count int) (steps []int) {
	k, step := 0, 0
	for {
		// https://en.wikipedia.org/wiki/Shellsort
		pow2k := int(math.Pow(2, float64(k)))
		if (k & 0x1) == 0 {
			pow2kMinus1 := int(math.Pow(2, float64(k>>1)))
			step = 9*(pow2k-pow2kMinus1) + 1
		} else {
			pow2kPlus1 := int(math.Pow(2, float64((k+1)>>1)))
			step = 8*pow2k - 6*pow2kPlus1 + 1
		}

		if step >= count {
			break
		}
		// append step to beginning of steps
		// steps = append([]int{step}, steps...)
		steps = append(steps, step)
		k++
	}
	return
}

// 有二分查找的版本，运行速度反而变慢了，所以还是删了吧
// 因为逆序对个数已经减少了，使用二分查找后比较的次数变多了
// func newShellSorterV3() *shellSorter {
// 	s := &shellSorter{}
// 	s.sortImple = s.shellSortV3
// 	return s
// }

// func (s *shellSorter) shellSortV3(ints []int) {
// 	s.array = ints

// 	steps := s.sedgewickStep(len(ints))
// 	for i := len(steps) - 1; i >= 0; i-- {
// 		s.shellImple2(steps[i])
// 	}
// }

// // 有二分查找的版本
// func (s *shellSorter) shellImple2(step int) {
// 	count := len(s.array)
// 	for col := 0; col < step; col++ {
// 		// col+step*0, col+step*1, col+step*2
// 		for begin := col + step; begin < count; begin += step {
// 			if s.CmpIndex(begin, begin-step) >= 0 {
// 				continue
// 			}

// 			// find pos for curl
// 			pos := s.findPosition(col, step, begin)
// 			// inset ints[begin] to pos
// 			s.insertPosition(begin, step, pos)
// 		}
// 	}
// }

// // [col, col+step, col+step*2, ..., curl)
// func (s *shellSorter) findPosition(col, step, cur int) int {
// 	start, end := 0, cur/step
// 	for start < end {
// 		mid := (start + end) >> 1
// 		index := col + step*mid
// 		if s.CmpIndex(index, cur) > 0 {
// 			end = mid
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return col + step*start
// }

// func (s *shellSorter) insertPosition(cur, step, pos int) {
// 	bak := s.array[cur]
// 	for i := cur; i > pos; i -= step {
// 		s.array[i] = s.array[i-step]
// 	}

// 	s.array[pos] = bak
// }
