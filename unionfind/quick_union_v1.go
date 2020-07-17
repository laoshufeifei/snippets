package unionfind

// QuickUnionV1 最简单的一种实现
type QuickUnionV1 struct {
	parents []int
}

func newQuickUnionV1(size int) *QuickUnionV1 {
	u := &QuickUnionV1{
		parents: make([]int, size),
	}

	for i := 0; i < size; i++ {
		u.parents[i] = i
	}
	return u
}

// Find 返回所有集合的父节点
func (u *QuickUnionV1) Find(idx int) int {
	for idx != u.parents[idx] {
		idx = u.parents[idx]
	}
	return idx
}

// Union 合并 idx1 与 idx2 所在的集合
func (u *QuickUnionV1) Union(idx1, idx2 int) {
	p1, p2 := u.Find(idx1), u.Find(idx2)
	if p1 == p2 {
		return
	}

	u.parents[p1] = p2
}

// IsSameUnion 判断 idx1 和 idx2 是否在同一集合里
func (u *QuickUnionV1) IsSameUnion(idx1, idx2 int) bool {
	return u.Find(idx1) == u.Find(idx2)
}
