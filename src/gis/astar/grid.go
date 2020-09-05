package astar

import (
	"fmt"
	"math"
)

// Grid 每一个小方格
// 's' 表示 start, 'e' 表示 end, 'x' 表示不可走，'*' 表示路径
// 为了简单，上下左右一个格代价是 10，对角线走一个格是 14
type Grid struct {
	value   byte
	x, y    int
	gWeight int // g(n)
	hWeight int // h(n)
	fWeight int // f(n) = g(n) + h(n)
	parent  *Grid
}

func newGrid(x, y int, value byte, parent *Grid) *Grid {
	return &Grid{
		x:      x,
		y:      y,
		value:  value,
		parent: parent,
	}
}

func (g *Grid) getG() int {
	return g.gWeight
}

func (g *Grid) setG(weight int) {
	g.gWeight = weight
	g.fWeight = g.gWeight + g.hWeight
}

func (g *Grid) getH() int {
	return g.hWeight
}

func (g *Grid) setH(weight int) {
	g.hWeight = weight
	g.fWeight = g.gWeight + g.hWeight
}

func (g *Grid) getF() int {
	return g.fWeight
}

// Equals ...
func (g *Grid) Equals(ref *Grid) bool {
	return g.x == ref.x && g.y == ref.y
}

func (g *Grid) isNotAvailable() bool {
	return g.value == 'x'
}

// 曼哈顿距离，只能上下左右的走
func (g *Grid) manhattanDistance(other *Grid) int {
	dx := int(math.Abs(float64(other.x - g.x)))
	dy := int(math.Abs(float64(other.y - g.y)))
	return (dx + dy) * 10
}

func (g *Grid) String() string {
	return fmt.Sprintf("[%d,%d] %d = %d + %d", g.x, g.y, g.fWeight, g.gWeight, g.hWeight)
}
