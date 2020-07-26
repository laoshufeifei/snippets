package graph

import (
	"fmt"
)

// vertex 图的定点
type vertex struct {
	name string
	// key 是另一个端点的名字，value 是 edge
	inMaps  map[string]*edge
	outMaps map[string]*edge
}

func newVertex(name string) *vertex {
	return &vertex{
		name:    name,
		inMaps:  make(map[string]*edge),
		outMaps: make(map[string]*edge),
	}
}

func (v *vertex) equals(other *vertex) bool {
	return v.name == other.name
}

func (v *vertex) addInEdge(e *edge) {
	// e.toVertex == v
	v.inMaps[e.fromVertex.name] = e
}

// addOutEdge 如果以前存在会被删掉
func (v *vertex) addOutEdge(e *edge) {
	// e.outVertex == v
	v.outMaps[e.toVertex.name] = e
}

func (v *vertex) String() string {
	return fmt.Sprintf("%s", v.name)
}
