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

// 这里是单点修改，区间查询的一种实现
type pointUpdateRangeQueryTree struct {
	size  int
	array []int
}

func newPointUpdateRangeQueryTree(values []int) *pointUpdateRangeQueryTree {
	length := len(values)
	t := pointUpdateRangeQueryTree{
		size:  length,
		array: make([]int, length+1),
	}

	for i := 1; i <= length; i++ {
		t.array[i] = values[i-1]

		// lowbit(i) 是 t.array[i] 可以覆盖到的长度
		// 即在 values 中的第 [i-lowbit(i)+1, i]
		// 换成索引表示就是 [i-lowbit(i), i-1]
		// 但是 values[i-1] 已经存储到 array[i] 中了
		for j := i - lowbit(i); j < i-1; j++ {
			t.array[i] += values[j]
		}
	}

	return &t
}

// idx 从 1 开始
func (t *pointUpdateRangeQueryTree) pointUpdate(idx, delta int) {
	for idx <= t.size {
		t.array[idx] += delta
		idx += lowbit(idx)
	}
}

// prefixSum 求前 idx 个元素的和
// idx 从 1 开始
func (t *pointUpdateRangeQueryTree) prefixSum(idx int) (result int) {
	if idx > t.size {
		idx = t.size
	}

	for idx > 0 {
		result += t.array[idx]
		idx -= lowbit(idx)
	}
	return
}

// rangeSum 求 [fromIdx, toIdx] 元素的和
// 从 1 开始
func (t *pointUpdateRangeQueryTree) rangeSum(fromIdx, toIdx int) int {
	return t.prefixSum(toIdx) - t.prefixSum(fromIdx-1)
}

// pointQuery 单点查询
func (t *pointUpdateRangeQueryTree) pointQuery(idx int) (result int) {
	return t.prefixSum(idx) - t.prefixSum(idx-1)
}
