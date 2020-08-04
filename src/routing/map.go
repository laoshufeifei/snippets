package routing

import "fmt"

// gridMap 整个地图, 为了简单使用了二维数组
// https://www.gamedev.net/reference/articles/article2003.asp
type gridMap struct {
	grids           [100][100]*Grid
	rowSize         int
	colSize         int
	begin           *Grid
	end             *Grid
	openList        *GridSet
	closeList       *GridSet
	canCrossCorners bool // 是否可以穿墙角
}

func newGridMap(gridBytes [][]byte, rowSize, colSize int) *gridMap {
	obj := &gridMap{
		rowSize:   rowSize,
		colSize:   colSize,
		openList:  newGridSet(),
		closeList: newGridSet(),
	}

	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			value := gridBytes[i][j]
			g := newGrid(i, j, value, nil)
			obj.grids[i][j] = g

			if value == 's' {
				obj.begin = g
			} else if value == 'e' {
				obj.end = g
			}
		}
	}

	return obj
}

func (m *gridMap) enableCrossCorners() {
	m.canCrossCorners = true
}

func (m *gridMap) getGrid(x, y int) *Grid {
	if x < 0 || x >= m.rowSize {
		return nil
	}

	if y < 0 || y >= m.colSize {
		return nil
	}

	return m.grids[x][y]
}

// 不允许穿墙角
func (m *gridMap) isReachable(g1, g2 *Grid) bool {
	if g1.isNotAvailable() || g2.isNotAvailable() {
		return false
	}

	if m.canCrossCorners {
		return true
	}

	x1, y1 := g1.x, g1.y
	x2, y2 := g2.x, g2.y

	dx, dy := x2-x1, y2-y1
	if (dx == -1 || dx == 1) && (dy == -1 || dy == 1) {
		// 不用判断是否为 nil
		x1y2 := m.getGrid(x1, y2)
		x2y1 := m.getGrid(x2, y1)

		if x1y2.isNotAvailable() || x2y1.isNotAvailable() {
			return false
		}
	}

	return true
}

func (m *gridMap) processSurrounding(parent *Grid) []*Grid {
	results := make([]*Grid, 0)

	x, y := parent.x, parent.y
	dG := 0
	for i := 1; i <= 8; i++ {
		var child *Grid
		switch i {
		case 1:
			child = m.getGrid(x-1, y-1) // 左上角
			dG = 14
		case 2:
			child = m.getGrid(x-1, y) // 上
			dG = 10
		case 3:
			child = m.getGrid(x-1, y+1) // 右上角
			dG = 14
		case 4:
			child = m.getGrid(x, y-1) // 左
			dG = 10
		case 5:
			child = m.getGrid(x, y+1) // 右
			dG = 10
		case 6:
			child = m.getGrid(x+1, y-1) // 左下角
			dG = 14
		case 7:
			child = m.getGrid(x+1, y) // 下
			dG = 10
		case 8:
			child = m.getGrid(x+1, y+1) // 右下角
			dG = 14
		}

		if child == nil || child.isNotAvailable() || m.closeList.Contains(child) {
			continue
		}

		// 不允许穿墙角
		if !m.isReachable(parent, child) {
			continue
		}

		oldG := child.getG()
		newG := dG + parent.getG()
		if m.openList.Contains(child) && oldG <= newG {
			continue
		}

		child.parent = parent
		child.setG(newG)
		// 使用麦哈顿距离来估算
		child.setH(child.manhattanDistance(m.end))

		results = append(results, child)
	}
	return results
}

func (m *gridMap) astar() []*Grid {
	found := false
	m.openList.Push(m.begin)

	for m.openList.Size() > 0 {
		cur := m.openList.Pop()
		fmt.Println("open  list pop ", cur)

		if cur.Equals(m.end) {
			found = true
			break
		}

		if m.closeList.Contains(cur) {
			continue
		}

		surrounding := m.processSurrounding(cur)
		for _, child := range surrounding {
			m.openList.Push(child)
			// fmt.Println("open  list push", child)
		}

		m.closeList.Push(cur)
		// fmt.Println("close list push", cur)
	}

	if !found {
		fmt.Println("not found")
		return nil
	}

	results := make([]*Grid, 0)
	parent := m.end.parent
	for parent != nil && !parent.Equals(m.begin) {
		fmt.Println(parent)
		parent.value = '*'
		results = append(results, parent)
		parent = parent.parent
	}

	for i := 0; i < m.rowSize; i++ {
		for j := 0; j < m.colSize; j++ {
			fmt.Printf("%c ", m.grids[i][j].value)
		}
		fmt.Println()
	}

	return results
}
