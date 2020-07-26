package unionfind

// QuickUnion 在使用 rank 的基础上再加上路径减半，支持多种类型
type QuickUnion struct {
	parents []int
	rank    []int

	// key is data, value is indexes
	size    int
	indexes map[interface{}]int
}

// NewQuickUnion ...
func NewQuickUnion(size int) *QuickUnion {
	u := &QuickUnion{
		parents: make([]int, size),
		rank:    make([]int, size),
		indexes: make(map[interface{}]int),
	}

	for i := 0; i < size; i++ {
		u.parents[i] = i
		u.rank[i] = 1
	}
	return u
}

// Push ...
func (u *QuickUnion) Push(items ...interface{}) {
	for _, item := range items {
		u.indexes[item] = u.size
		u.size++
	}
}

// Find 返回所有集合的父节点(路径减半)
// 返回父节点所在的索引
func (u *QuickUnion) Find(item interface{}) int {
	idx, ok := u.indexes[item]
	if !ok {
		return -1
	}

	return u.findWithIndex(idx)
}

func (u *QuickUnion) findWithIndex(idx int) int {
	for idx != u.getParent(idx) {
		g := u.getGrandparent(idx)
		u.setParent(idx, g)
		idx = g
	}
	return idx
}

// Union 合并 idx1 与 idx2 所在的集合
func (u *QuickUnion) Union(item1, item2 interface{}) {
	idx1, ok1 := u.indexes[item1]
	idx2, ok2 := u.indexes[item2]
	if !ok1 || !ok2 {
		return
	}

	p1, p2 := u.findWithIndex(idx1), u.findWithIndex(idx2)
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

// IsSameUnion 判断 item1 和 item2 是否在同一集合里
func (u *QuickUnion) IsSameUnion(item1, item2 interface{}) bool {
	p1 := u.Find(item1)
	p2 := u.Find(item2)
	if p1 == -1 || p2 == -1 {
		return false
	}

	return p1 == p2
}

func (u *QuickUnion) getParent(idx int) int {
	return u.parents[idx]
}

func (u *QuickUnion) getGrandparent(idx int) int {
	return u.getParent(u.getParent(idx))
}

func (u *QuickUnion) setParent(idx, value int) {
	u.parents[idx] = value
}
