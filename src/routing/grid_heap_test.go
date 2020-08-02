package routing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridHeaper(t *testing.T) {
	test := assert.New(t)
	test.Equal(1, 1)

	heaper := newGridHeap()
	test.Equal(heaper.Size(), 0)

	g1 := newGrid(1, 2, 'x', nil)
	g1.setG(14)
	heaper.Push(g1)

	g2 := newGrid(2, 2, 'x', nil)
	g2.setG(10)
	heaper.Push(g2)

	g3 := newGrid(3, 2, 'x', nil)
	g3.setG(40)
	heaper.Push(g3)
	test.Equal(heaper.Size(), 3)

	g := heaper.Pop()
	test.Equal(g.getF(), 10)
	test.Equal(heaper.Size(), 2)

	g = heaper.Pop()
	test.Equal(g.getF(), 14)

	g = heaper.Pop()
	test.Equal(g.getF(), 40)

	g = heaper.Pop()
	test.True(g == nil)
}
