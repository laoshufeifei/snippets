package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdgeSet(t *testing.T) {
	test := assert.New(t)
	set := newEdgeSet()

	e1 := newEdgeByVerticeName("A", "B")
	e1.setWeight(6.)
	set.Add(e1)

	e2 := newEdgeByVerticeName("A", "C")
	e2.setWeight(1.)
	set.Add(e2)

	e3 := newEdgeByVerticeName("C", "D")
	e3.setWeight(13.)
	set.Add(e3)
	test.True(set.Size() == 3)

	set.Sort()
	items := set.Items
	test.True(items[0].equals(e2))
	test.True(items[2].equals(e3))

	set.Remove(e3)
	test.True(set.Size() == 2)
}
