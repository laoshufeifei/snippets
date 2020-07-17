package unionfind

// QuickFindUnion ...
type QuickFindUnion struct {
	parents []int
}

func newQuickFindUnion(size int) *QuickFindUnion {
	u := &QuickFindUnion{
		parents: make([]int, size),
	}

	for i := 0; i < size; i++ {
		u.parents[i] = i
	}
	return u
}

// Find 返回所有集合的父节点
func (u *QuickFindUnion) Find(idx int) int {
	return u.parents[idx]
}

// Union 合并 idx1 所在的集合 到 idx2 的父节点下
func (u *QuickFindUnion) Union(idx1, idx2 int) {
	p1, p2 := u.Find(idx1), u.Find(idx2)
	if p1 == p2 {
		return
	}

	for i := 0; i < len(u.parents); i++ {
		if u.parents[i] == p1 {
			u.parents[i] = p2
		}
	}
}

// IsSameUnion 判断 idx1 和 idx2 是否在同一集合里
func (u *QuickFindUnion) IsSameUnion(idx1, idx2 int) bool {
	return u.Find(idx1) == u.Find(idx2)
}
