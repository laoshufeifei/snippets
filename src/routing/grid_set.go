package routing

import "sort"

// GridItems ...
type GridItems []*Grid

// GridSet ...
type GridSet struct {
	items GridItems
}

func (gs GridItems) Len() int           { return len(gs) }
func (gs GridItems) Less(i, j int) bool { return gs[i].fWeight < gs[j].fWeight }
func (gs GridItems) Swap(i, j int)      { gs[i], gs[j] = gs[j], gs[i] }

func newGridSet() *GridSet {
	return &GridSet{
		items: make([]*Grid, 0),
	}
}

// Contains ...
func (set *GridSet) Contains(g *Grid) bool {
	for _, ref := range set.items {
		if ref.Equals(g) {
			return true
		}
	}

	return false
}

// Push ...
func (set *GridSet) Push(g *Grid) {
	if set.Contains(g) {
		set.Remove(g)
	}

	set.items = append(set.items, g)
}

// Remove ...
func (set *GridSet) Remove(g *Grid) {
	count := len(set.items)
	for i := 0; i < count; i++ {
		edge := set.items[i]
		if edge.Equals(g) {
			set.items[i] = nil
			copy(set.items[i:], set.items[i+1:count])
			set.items = set.items[:count-1]
			return
		}
	}
}

// Size ...
func (set *GridSet) Size() int {
	return len(set.items)
}

// Sort ...
func (set *GridSet) Sort() {
	sort.Sort(set.items)
}

// Pop ...
func (set *GridSet) Pop() *Grid {
	if set.Size() <= 0 {
		return nil
	}

	set.Sort()
	old := set.items[0]
	set.Remove(old)
	return old
}
