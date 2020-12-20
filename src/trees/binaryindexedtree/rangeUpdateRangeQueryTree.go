package binaryindexedtree

// 树状数组
// https://en.wikipedia.org/wiki/Fenwick_tree
/*
8   ----------------------------------------------------1000--
4   ----------------------0100--
2   ------0010--                -------0110--
1   0001           0011           0101          0111
      1       2      3      4       5      6      7       8
    0001    0010   0011   0100    0101   0110   0111    1000
*/

// 这里是区间修改，区间查询的一种实现，需要用到差分数组
type rangeUpdateRangeQueryTree struct {
	size   int
	array1 []int
	array2 []int
}

func newRangeUpdateRangeQueryTree(values []int) *rangeUpdateRangeQueryTree {
	length := len(values)
	if length == 0 {
		return nil
	}

	t := rangeUpdateRangeQueryTree{
		size:   length,
		array1: make([]int, length+1),
		array2: make([]int, length+1),
	}

	values2 := make([]int, length)
	for i := length - 1; i > 0; i-- {
		values[i] = values[i] - values[i-1]
		values2[i] = values[i] * i
	}

	for i := 1; i <= length; i++ {
		t.array1[i] = values[i-1]
		for j := i - lowbit(i); j < i-1; j++ {
			t.array1[i] += values[j]
		}
	}

	for i := 1; i <= length; i++ {
		t.array2[i] = values2[i-1]
		for j := i - lowbit(i); j < i-1; j++ {
			t.array2[i] += values2[j]
		}
	}

	return &t
}

// pointUpdate 更新原始的值
// idx 从 1 开始
func (t *rangeUpdateRangeQueryTree) pointUpdate(idx, delta int) {
	x := idx
	for idx <= t.size {
		t.array1[idx] += delta
		t.array2[idx] += delta * (x - 1)
		idx += lowbit(idx)
	}
}

func (t *rangeUpdateRangeQueryTree) rangeUpdate(fromIdx, toIdx, delta int) {
	t.pointUpdate(fromIdx, delta)
	t.pointUpdate(toIdx+1, -delta)
}

// 前面 k 个元素的和
func (t *rangeUpdateRangeQueryTree) prefixSum(idx int) (result int) {
	if idx > t.size {
		idx = t.size
	}

	x := idx
	for idx > 0 {
		result += t.array1[idx]*x - t.array2[idx]
		idx -= lowbit(idx)
	}
	return
}

// pointQuery 求前 idx 个元素的和，即原来元素的值
// idx 从 1 开始
func (t *rangeUpdateRangeQueryTree) pointQuery(idx int) int {
	return t.prefixSum(idx) - t.prefixSum(idx-1)
}

// rangeSum 求 [fromIdx, toIdx] 元素的和
func (t *rangeUpdateRangeQueryTree) rangeSum(fromIdx, toIdx int) int {
	return t.prefixSum(toIdx) - t.prefixSum(fromIdx-1)
}
