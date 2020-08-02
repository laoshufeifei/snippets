package routing

import "math"

// grid 每一个小方格
// 's' 表示 start, 'e' 表示 end, 'x' 表示不可走，'*' 表示路径
// 为了简单，上下左右一个格代价是 10，对角线走一个格是 14
type grid struct {
	value   byte
	x, y    int
	gWeight int // g(n)
	hWeight int // h(n)
	fWeight int // f(n) = g(n) + h(n)
	parent  *grid
}

func newGrid(x, y int, value byte, parent *grid) *grid {
	return &grid{
		x:      x,
		y:      y,
		value:  value,
		parent: parent,
	}
}

func (g *grid) getG() int {
	return g.gWeight
}

func (g *grid) setG(weight int) {
	g.gWeight = weight
	g.fWeight = g.gWeight + g.hWeight
}

func (g *grid) getH() int {
	return g.hWeight
}

func (g *grid) setH(weight int) {
	g.hWeight = weight
	g.fWeight = g.gWeight + g.hWeight
}

func (g *grid) getF() int {
	return g.fWeight
}

func (g *grid) Equals(ref *grid) bool {
	return g.x == ref.x && g.y == ref.y
}

func (g *grid) isNotAvailable() bool {
	return g.value == 'x'
}

// 曼哈顿距离，只能上下左右的走
func (g *grid) manhattanDistance(other *grid) int {
	dx := int(math.Abs(float64(other.x - g.x)))
	dy := int(math.Abs(float64(other.y - g.y)))
	return (dx + dy) * 10
}
