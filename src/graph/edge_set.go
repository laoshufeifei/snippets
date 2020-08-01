package graph

import (
	"math"
	"sort"
)

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

// Clone ...
func (set *EdgeSet) Clone() *EdgeSet {
	newSet := &EdgeSet{
		Items: make(EdgeItems, set.Size()),
	}

	copy(newSet.Items, set.Items)
	return newSet
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

// find ...
func (set *EdgeSet) find(from, to string) *edge {
	for _, e := range set.Items {
		if e.fromVertex.name == from && e.toVertex.name == to {
			return e
		}
	}
	return nil
}

// findByFrom ...
func (set *EdgeSet) findByFrom(from string) *EdgeSet {
	edges := make([]*edge, 0)
	for _, e := range set.Items {
		if e.fromVertex.name == from {
			edges = append(edges, e)
		}
	}

	newSet := &EdgeSet{
		Items: make(EdgeItems, len(edges)),
	}

	for i, e := range edges {
		newSet.Items[i] = e
	}

	return newSet
}

// findByTo ...
func (set *EdgeSet) findByTo(to string) *EdgeSet {
	edges := make([]*edge, 0)
	for _, e := range set.Items {
		if e.toVertex.name == to {
			edges = append(edges, e)
		}
	}

	newSet := &EdgeSet{
		Items: make(EdgeItems, len(edges)),
	}

	for _, e := range edges {
		newSet.Items = append(newSet.Items, e)
	}

	return newSet
}

func (set *EdgeSet) findMinWeight() *edge {
	set.Sort()
	return set.Items[0]
}

func (set *EdgeSet) getWeight(fromName, toName string) float64 {
	if fromName == toName {
		return 0.
	}

	e := set.find(fromName, toName)
	if e == nil {
		return math.MaxFloat64
	}

	return e.weight
}

func (set *EdgeSet) updateWeight(fromName, toName string, weight float64) {
	e := set.find(fromName, toName)
	if e == nil {
		e = newEdgeByVerticeName(fromName, toName)
		e.setWeight(weight)
		set.Add(e)
		return
	}

	e.setWeight(weight)
}

func (set *EdgeSet) getWeightWithPassby(fromName, toName, passby string) float64 {
	e1 := set.find(fromName, passby)
	if e1 == nil {
		return math.MaxFloat64
	}

	e2 := set.find(passby, toName)
	if e2 == nil {
		return math.MaxFloat64
	}

	return e1.weight + e2.weight
}
