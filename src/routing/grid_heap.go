package routing

import "container/heap"

// 实现 container/heap 要求的各种接口
// https://pkg.go.dev/container/heap?tab=doc#example-package-IntHeap
type gridElements []*grid

func (l gridElements) Len() int           { return len(l) }
func (l gridElements) Less(i, j int) bool { return l[i].getF() < l[j].getF() }
func (l gridElements) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func (l *gridElements) Push(x interface{}) {
	g := x.(*grid)
	*l = append(*l, g)
}

func (l *gridElements) Pop() interface{} {
	oldList := *l
	length := len(oldList)

	grid := oldList[length-1]
	oldList[length-1] = nil
	*l = oldList[:length-1]

	return grid
}

type gridHeap struct {
	elements gridElements
}

func newGridHeap() *gridHeap {
	grids := make(gridElements, 0)
	heap.Init(&grids)

	result := &gridHeap{
		elements: grids,
	}
	return result
}

func (l *gridHeap) Pop() *grid {
	if l.Size() == 0 {
		return nil
	}

	ret := heap.Pop(&l.elements)
	return ret.(*grid)
}

func (l *gridHeap) Push(g *grid) {
	heap.Push(&l.elements, g)
}

func (l *gridHeap) Size() int {
	return len(l.elements)
}
