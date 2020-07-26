package graph

import (
	"gossips/src/hashset"
	"gossips/src/queue"
	"gossips/src/stack"
	"gossips/src/unionfind"
	"strings"
)

// DirectedGraph 有向图
type DirectedGraph struct {
	vertices map[string]*vertex
	edgeSet  *EdgeSet
}

func newDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		vertices: make(map[string]*vertex),
		edgeSet:  newEdgeSet(),
	}
}

func (g *DirectedGraph) addVertex(v *vertex) {
	_, ok := g.vertices[v.name]
	if ok {
		return
	}

	g.vertices[v.name] = v
}

func (g *DirectedGraph) addVertexByName(name string) {
	_, ok := g.vertices[name]
	if ok {
		return
	}

	vertex := newVertex(name)
	g.vertices[name] = vertex
}

func (g *DirectedGraph) removeVertex(name string) {
	vertex, ok := g.vertices[name]
	if !ok {
		return
	}
	delete(g.vertices, name)

	for _, e := range vertex.inMaps {
		g.edgeSet.Remove(e)
	}
	for _, e := range vertex.outMaps {
		g.edgeSet.Remove(e)
	}
}

// addEdge ...
func (g *DirectedGraph) addEdge(fromName, toName string, weight float64) {
	fromName = strings.TrimSpace(fromName)
	from, ok := g.vertices[fromName]
	if !ok {
		from = newVertex(fromName)
		g.vertices[fromName] = from
	}

	toName = strings.TrimSpace(toName)
	to, ok := g.vertices[toName]
	if !ok {
		to = newVertex(toName)
		g.vertices[toName] = to
	}

	e := newEdge(from, to)
	e.setWeight(weight)
	g.edgeSet.Add(e)
}

func (g *DirectedGraph) breadthFirstSearch(name string) [][]*vertex {
	begin, ok := g.vertices[name]
	if !ok {
		return nil
	}

	results := make([][]*vertex, 0)
	oneLevel := make([]*vertex, 0)

	searchList := queue.New()
	hasPushSet := hashset.New()

	searchList.Push(begin)
	hasPushSet.Push(begin)

	// 标记这一层遍历了多少个了
	index := 0
	nextLevelSize := searchList.Size()
	for !searchList.IsEmpty() {
		iter, _ := searchList.Poll()
		v := iter.(*vertex)

		index++
		oneLevel = append(oneLevel, v)

		for _, e := range v.outMaps {
			if hasPushSet.Contains(e.toVertex) {
				continue
			}

			searchList.Push(e.toVertex)
			hasPushSet.Push(e.toVertex)
		}

		if index == nextLevelSize {
			index = 0
			nextLevelSize = searchList.Size()
			results = append(results, oneLevel)
			oneLevel = make([]*vertex, 0)
		}
	}

	return results
}

func (g *DirectedGraph) depthFirstSearch(name string) []*vertex {
	begin, ok := g.vertices[name]
	if !ok {
		return nil
	}

	results := make([]*vertex, 0)
	searchList := stack.New()
	hasPushSet := hashset.New()

	searchList.Push(begin)
	hasPushSet.Push(begin)

	for !searchList.IsEmpty() {
		iter, _ := searchList.Pop()
		vertex := iter.(*vertex)
		results = append(results, vertex)

		for _, e := range vertex.outMaps {
			to := e.toVertex
			if hasPushSet.Contains(to) {
				continue
			}

			searchList.Push(to)
			hasPushSet.Push(to)
		}
	}

	return results
}

func (g *DirectedGraph) depthFirstSearch2(name string) []*vertex {
	begin, ok := g.vertices[name]
	if !ok {
		return nil
	}

	results := make([]*vertex, 0)
	hadSearch := make(map[string]bool)
	g.dfs(begin, &results, hadSearch)
	return results
}

func (g *DirectedGraph) dfs(begin *vertex, results *[]*vertex, hadSearch map[string]bool) {
	_, ok := hadSearch[begin.name]
	if ok {
		return
	}

	*results = append(*results, begin)
	hadSearch[begin.name] = true
	for _, e := range begin.outMaps {
		g.dfs(e.toVertex, results, hadSearch)
	}
}

// kruskal 最小生成树
func (g *DirectedGraph) kruskal() []*edge {
	results := make([]*edge, 0)

	vSize := len(g.vertices)
	union := unionfind.NewQuickUnion(vSize)
	for _, v := range g.vertices {
		union.Push(v)
	}

	g.edgeSet.Sort()
	items := g.edgeSet.Items
	for _, e := range items {
		if len(results) >= vSize {
			break
		}

		if union.IsSameUnion(e.fromVertex, e.toVertex) {
			continue
		}

		results = append(results, e)
		union.Union(e.fromVertex, e.toVertex)
	}

	return results
}