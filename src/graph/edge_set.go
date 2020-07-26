package graph

import "sort"

// EdgeItems ...
type EdgeItems []*edge

func (s EdgeItems) Len() int           { return len(s) }
func (s EdgeItems) Less(i, j int) bool { return s[i].weight < s[j].weight }
func (s EdgeItems) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// EdgeSet ...
type EdgeSet struct {
	Items EdgeItems
}

func newEdgeSet() *EdgeSet {
	return &EdgeSet{
		Items: make(EdgeItems, 0),
	}
}

// Contains ...
func (set *EdgeSet) Contains(e *edge) bool {
	for _, ref := range set.Items {
		if ref.equals(e) {
			return true
		}
	}

	return false
}

// Add ...
func (set *EdgeSet) Add(e *edge) {
	if set.Contains(e) {
		set.Remove(e)
	}

	set.Items = append(set.Items, e)
}

// Remove ...
func (set *EdgeSet) Remove(e *edge) {
	count := len(set.Items)
	for i := 0; i < count; i++ {
		edge := set.Items[i]
		if edge.equals(e) {
			set.Items[i] = nil
			copy(set.Items[i:], set.Items[i+1:count])
			set.Items = set.Items[:count-1]
			return
		}
	}
}

// Size ...
func (set *EdgeSet) Size() int {
	return len(set.Items)
}

// Sort ...
func (set *EdgeSet) Sort() {
	sort.Sort(set.Items)
}
