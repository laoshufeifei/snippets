package unionfind

// QuickUnionV2 使用基于 rank 的优化
type QuickUnionV2 struct {
	parents []int
	rank    []int
}

func newQuickUnionV2(size int) *QuickUnionV2 {
	u := &QuickUnionV2{
		parents: make([]int, size),
		rank:    make([]int, size),
	}

	for i := 0; i < size; i++ {
		u.parents[i] = i
		u.rank[i] = 1
	}
	return u
}

// Find 返回所有集合的父节点
func (u *QuickUnionV2) Find(idx int) int {
	for idx != u.parents[idx] {
		idx = u.parents[idx]
	}
	return idx
}

// Union 合并 idx1 与 idx2 所在的集合
func (u *QuickUnionV2) Union(idx1, idx2 int) {
	p1, p2 := u.Find(idx1), u.Find(idx2)
	if p1 == p2 {
		return
	}

	if u.rank[p1] < u.rank[p2] {
		u.parents[p1] = p2
	} else if u.rank[p1] > u.rank[p2] {
		u.parents[p2] = p1
	} else {
		u.parents[p1] = p2
		u.rank[p2]++
	}
}

// IsSameUnion 判断 idx1 和 idx2 是否在同一集合里
func (u *QuickUnionV2) IsSameUnion(idx1, idx2 int) bool {
	return u.Find(idx1) == u.Find(idx2)
}
