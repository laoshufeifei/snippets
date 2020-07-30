package graph

import (
	"fmt"
)

// vertex 图的定点
type vertex struct {
	name     string
	inEdges  *EdgeSet
	outEdges *EdgeSet
}

func newVertex(name string) *vertex {
	return &vertex{
		name:     name,
		inEdges:  newEdgeSet(),
		outEdges: newEdgeSet(),
	}
}

func (v *vertex) equals(other *vertex) bool {
	return v.name == other.name
}

func (v *vertex) addInEdge(e *edge) {
	v.inEdges.Add(e)
}

// addOutEdge 如果以前存在会被删掉
func (v *vertex) addOutEdge(e *edge) {
	v.outEdges.Add(e)
}

func (v *vertex) String() string {
	return fmt.Sprintf("%s", v.name)
}
