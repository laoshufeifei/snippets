package segmenttree

// 线段树简单用来求和
// https://www.bilibili.com/video/BV1cb411t7AM/
type segmentTree struct {
	size  int
	tree  []int
	array []int
}

func newSegmentTree(values []int) *segmentTree {
	length := len(values)
	t := segmentTree{
		size:  length,
		tree:  make([]int, length*2+1),
		array: make([]int, length), // 其实没必要保存
	}

	copy(t.array, values)
	t.recursionBuildTree(0, length-1, 0)
	return &t
}

func (t *segmentTree) recursionBuildTree(start, end, nodeIdx int) {
	if start == end {
		t.tree[nodeIdx] = t.array[start]
		return
	}

	mid := (start + end) >> 1
	leftNodeIdx, rightNodeIdx := 2*nodeIdx+1, 2*nodeIdx+2

	t.recursionBuildTree(start, mid, leftNodeIdx)
	t.recursionBuildTree(mid+1, end, rightNodeIdx)

	// 求和
	t.tree[nodeIdx] = t.tree[leftNodeIdx] + t.tree[rightNodeIdx]
}

// 将原始的 array[updateIdx] 改为 value
func (t *segmentTree) update(idx, value int) {
	t.updateTree(0, t.size-1, 0, idx, value)
}

func (t *segmentTree) updateTree(start, end, nodeIdx, updateIdx, value int) {
	if start == end {
		t.array[updateIdx] = value
		t.tree[nodeIdx] = value
		return
	}

	mid := (start + end) >> 1
	leftNodeIdx, rightNodeIdx := 2*nodeIdx+1, 2*nodeIdx+2

	if updateIdx >= start && updateIdx <= mid {
		t.updateTree(start, mid, leftNodeIdx, updateIdx, value)
	} else {
		t.updateTree(mid+1, end, rightNodeIdx, updateIdx, value)
	}

	t.tree[nodeIdx] = t.tree[leftNodeIdx] + t.tree[rightNodeIdx]
}

func (t *segmentTree) query(L, R int) int {
	return t.queryTree(0, t.size-1, 0, L, R)
}

func (t *segmentTree) queryTree(start, end, nodeIdx, L, R int) int {
	if R < start || L > end {
		return 0
	} else if (L <= start && end <= R) || start == end {
		return t.tree[nodeIdx]
	}

	mid := (start + end) >> 1
	leftNodeIdx, rightNodeIdx := 2*nodeIdx+1, 2*nodeIdx+2

	leftSum := t.queryTree(start, mid, leftNodeIdx, L, R)
	rightSum := t.queryTree(mid+1, end, rightNodeIdx, L, R)
	return leftSum + rightSum
}
