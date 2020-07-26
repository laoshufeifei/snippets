package graph

import "fmt"

// edge 图的边
type edge struct {
	weight     float64
	fromVertex *vertex
	toVertex   *vertex
}

func newEdge(from *vertex, to *vertex) *edge {
	e := &edge{
		fromVertex: from,
		toVertex:   to,
	}

	from.addOutEdge(e)
	to.addInEdge(e)
	return e
}

func newEdgeByVerticeName(from, to string) *edge {
	fromVertex := newVertex(from)
	toVertex := newVertex(to)

	return &edge{
		fromVertex: fromVertex,
		toVertex:   toVertex,
	}
}

func (e *edge) equals(other *edge) bool {
	return e.fromVertex.equals(other.fromVertex) && e.toVertex.equals(other.toVertex)
}

func (e *edge) setWeight(w float64) {
	e.weight = w
}

func (e *edge) String() string {
	return fmt.Sprintf("from: %s, to: %s, weight: %f", e.fromVertex.name, e.toVertex.name, e.weight)
}
