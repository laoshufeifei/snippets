package routing

// gridMap 整个地图, 为了简单使用了二维数组
// https://www.gamedev.net/reference/articles/article2003.asp
type gridMap struct {
	grids    [100][100]*grid
	openList *gridHeap
	rowSize  int
	colSize  int
	begin    *grid
	end      *grid
}

func newGridMap(gridBytes [][]byte, rowSize, colSize int) *gridMap {
	obj := &gridMap{
		openList: newGridHeap(),
		rowSize:  rowSize,
		colSize:  colSize,
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

func (m *gridMap) getGrid(x, y int) *grid {
	if x < 0 || x >= m.rowSize {
		return nil
	}

	if y < 0 || y >= m.colSize {
		return nil
	}

	return m.grids[x][y]
}

func (m *gridMap) processSurrounding(parent *grid) []*grid {
	results := make([]*grid, 0)

	x, y := parent.x, parent.y
	dG := 0
	for i := 1; i <= 8; i++ {
		var child *grid
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

		if child == nil || child.isNotAvailable() {
			continue
		}

		child.parent = parent
		child.setG(dG + parent.getG())
		// 使用麦哈顿距离来估算
		child.setH(child.manhattanDistance(m.end))

		results = append(results, child)
	}
	return results
}
