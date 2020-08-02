package routing

// GridSet ...
type GridSet struct {
	Items []*grid
}

func newGridSet() *GridSet {
	return &GridSet{
		Items: make([]*grid, 0),
	}
}

// Clone ...
func (set *GridSet) Clone() *GridSet {
	newSet := &GridSet{
		Items: make([]*grid, set.Size()),
	}

	copy(newSet.Items, set.Items)
	return newSet
}

// Contains ...
func (set *GridSet) Contains(g *grid) bool {
	for _, ref := range set.Items {
		if ref.Equals(g) {
			return true
		}
	}

	return false
}

// Add ...
func (set *GridSet) Add(g *grid) {
	if set.Contains(g) {
		set.Remove(g)
	}

	set.Items = append(set.Items, g)
}

// Remove ...
func (set *GridSet) Remove(g *grid) {
	count := len(set.Items)
	for i := 0; i < count; i++ {
		edge := set.Items[i]
		if edge.Equals(g) {
			set.Items[i] = nil
			copy(set.Items[i:], set.Items[i+1:count])
			set.Items = set.Items[:count-1]
			return
		}
	}
}

// Size ...
func (set *GridSet) Size() int {
	return len(set.Items)
}
