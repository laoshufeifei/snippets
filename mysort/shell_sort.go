package mysort

import (
	"math"
)

func shellSortV1(ints []int) {
	steps := shellStep(len(ints))
	for _, step := range steps {
		shellImple(ints, step)
	}
}

func shellSort(ints []int) {
	steps := sedgewickStep(len(ints))
	for i := len(steps) - 1; i >= 0; i-- {
		// shellImple2(ints, steps[i])
		shellImple(ints, steps[i])
	}
}

// 无二分查找优化的版本
func shellImple(ints []int, step int) {
	count := len(ints)
	for col := 0; col < step; col++ {
		// col+step*0, col+step*1, col+step*2
		for begin := col + step; begin < count; begin += step {
			cur := begin
			for cur > col {
				pre := cur - step
				if ints[cur] >= ints[pre] {
					break
				}
				ints[cur], ints[pre] = ints[pre], ints[cur]
				cur -= step
			}
		}
	}
}

// 有二分查找优化的版本
// 使用二分查找后反而变慢了，因为逆序对个数已经减少了，使用二分查找后比较的次数变多了
func shellImple2(ints []int, step int) {
	count := len(ints)
	for col := 0; col < step; col++ {
		// col+step*0, col+step*1, col+step*2
		for begin := col + step; begin < count; begin += step {
			if ints[begin] >= ints[begin-step] {
				continue
			}

			// find pos for curl
			pos := findSheelPos(ints, col, step, begin)
			// inset ints[begin] to pos
			insertSheelPosition(ints, begin, step, pos)
		}
	}
}

// [col, col+step, col+step*2, ..., curl)
func findSheelPos(ints []int, col, step, cur int) int {
	start, end := 0, cur/step
	for start < end {
		mid := (start + end) >> 1
		index := col + step*mid
		if ints[index] > ints[cur] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return col + step*start
}

func insertSheelPosition(ints []int, cur, step, pos int) {
	bak := ints[cur]
	for i := cur; i > pos; i -= step {
		ints[i] = ints[i-step]
	}

	ints[pos] = bak
}

// 生成的序列是从大到小的
func shellStep(count int) (steps []int) {
	step := count >> 1
	for step > 0 {
		steps = append(steps, step)
		step = step >> 1
	}
	return
}

// 生成的序列是从小到大的，使用的时候要倒序使用
func sedgewickStep(count int) (steps []int) {
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
