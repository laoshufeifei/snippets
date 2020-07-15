package unionfind

// QuickUnionV3 在使用 rank 的基础上再加上路径减半
type QuickUnionV3 struct {
	parents []int
	rank    []int
}

func newQuickUnionV3(size int) *QuickUnionV3 {
	u := &QuickUnionV3{
		parents: make([]int, size),
		rank:    make([]int, size),
	}

	for i := 0; i < size; i++ {
		u.parents[i] = i
		u.rank[i] = 1
	}
	return u
}

// Find 返回所有集合的父节点(路径压缩)
func (u *QuickUnionV3) Find(idx int) int {
	return u.FindWithPathHalf(idx)
}

// FindWithPathCompression 返回所有集合的父节点(路径压缩)
func (u *QuickUnionV3) FindWithPathCompression(idx int) int {
	if idx != u.parents[idx] {
		// p := u.parents[idx]
		// root := u.FindWithPathCompression(p)
		// u.parents[idx] = root
		u.parents[idx] = u.FindWithPathCompression(u.parents[idx])
	}
	return u.parents[idx]
}

// FindWithPathSplit 返回所有集合的父节点(路径分裂)
func (u *QuickUnionV3) FindWithPathSplit(idx int) int {
	for idx != u.parents[idx] {
		p := u.parents[idx]
		u.parents[idx] = u.parents[p]
		idx = p
	}
	return idx
}

// FindWithPathHalf 返回所有集合的父节点(路径减半)
func (u *QuickUnionV3) FindWithPathHalf(idx int) int {
	for idx != u.parents[idx] {
		g := u.parents[u.parents[idx]]
		u.parents[idx] = g
		idx = g
	}
	return idx
}

// Union 合并 idx1 与 idx2 所在的集合
func (u *QuickUnionV3) Union(idx1, idx2 int) {
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
func (u *QuickUnionV3) IsSameUnion(idx1, idx2 int) bool {
	return u.FindWithPathHalf(idx1) == u.FindWithPathHalf(idx2)
}
