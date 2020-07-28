package graph

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphBFS(t *testing.T) {
	test := assert.New(t)
	g := newDirectedGraph()
	edges := [][]string{
		{"0", "1"},
		{"0", "4"},
		{"1", "2"},
		{"2", "0"},
		{"2", "4"},
		{"2", "5"},
		{"3", "1"},
		{"4", "6"},
		{"4", "7"},
		{"5", "3"},
		{"5", "7"},
		{"6", "2"},
		{"6", "7"},
	}
	for _, vs := range edges {
		g.addEdge(vs[0], vs[1], 0.)
	}

	results := g.breadthFirstSearch("0")
	test.Equal(len(results), 5)

	allString := make([]string, 0)
	for _, oneLevel := range results {
		tmpString := make([]string, 0)
		for _, v := range oneLevel {
			tmpString = append(tmpString, v.name)
		}
		sort.Strings(tmpString)
		allString = append(allString, tmpString...)
	}
	test.Equal(allString, []string{"0", "1", "4", "2", "6", "7", "5", "3"})
}

func TestGraphDFS(t *testing.T) {
	test := assert.New(t)

	g := newDirectedGraph()
	edges := [][]string{
		{"a", "b"},
		{"a", "e"},
		{"b", "e"},
		{"c", "b"},
		{"d", "a"},
		{"e", "c"},
		{"e", "f"},
		{"f", "c"},
	}
	for _, vs := range edges {
		g.addEdge(vs[0], vs[1], 0)
	}

	vertices := g.depthFirstSearch("a")
	results := make([]string, 0)
	for _, v := range vertices {
		results = append(results, v.name)
	}

	s := strings.Join(results, ", ")
	s1 := "a, b, e, c, f"
	s2 := "a, b, e, f, c"
	s3 := "a, e, f, c, b"
	s4 := "a, e, c, b, f"
	fmt.Println(s)
	test.True(s == s1 || s == s2 || s == s3 || s == s4)
}

func TestGraphDFS2(t *testing.T) {
	test := assert.New(t)

	g := newDirectedGraph()
	edges := [][]string{
		{"a", "b"},
		{"a", "e"},
		{"b", "e"},
		{"c", "b"},
		{"d", "a"},
		{"e", "c"},
		{"e", "f"},
		{"f", "c"},
	}
	for _, vs := range edges {
		g.addEdge(vs[0], vs[1], 0)
	}

	vertices := g.depthFirstSearch2("a")
	results := make([]string, 0)
	for _, v := range vertices {
		results = append(results, v.name)
	}

	s := strings.Join(results, ", ")
	s1 := "a, b, e, c, f"
	s2 := "a, b, e, f, c"
	s3 := "a, e, f, c, b"
	s4 := "a, e, c, b, f"
	test.True(s == s1 || s == s2 || s == s3 || s == s4)
}

func TestGraphKruskal(t *testing.T) {
	test := assert.New(t)

	g := newDirectedGraph()
	edges := []string{
		"A, B, 4",
		"A, H, 8",
		"B, C, 8",
		"B, H, 11",
		"C, D, 7",
		"C, F, 4",
		"C, I, 2",
		"D, E, 9",
		"D, F, 14",
		"E, F, 10",
		"F, G, 2",
		"G, H, 1",
		"G, I, 6",
		"H, I, 7",
	}
	for _, s := range edges {
		parts := strings.Split(s, ",")
		parts[2] = strings.TrimSpace(parts[2])
		weight, _ := strconv.ParseFloat(parts[2], 64)

		g.addEdge(parts[0], parts[1], weight)
		g.addEdge(parts[1], parts[0], weight)
	}

	totalWeight := 0.
	results := g.kruskal()
	for _, e := range results {
		totalWeight += e.weight
	}
	test.Equal(totalWeight, 37.)
}

func TestGraphDijkstra(t *testing.T) {
	test := assert.New(t)

	g := newDirectedGraph()
	edges := []string{
		"A, B, 10",
		"A, D, 30",
		"A, E, 100",
		"B, C, 50",
		"C, E, 10",
		"D, C, 20",
		"D, E, 60",
	}
	for _, s := range edges {
		parts := strings.Split(s, ",")
		parts[2] = strings.TrimSpace(parts[2])
		weight, _ := strconv.ParseFloat(parts[2], 64)

		g.addEdge(parts[0], parts[1], weight)
	}

	results := g.dijkstra("A")
	weightB, _ := results["B"]
	test.Equal(weightB, 10.)

	weightC, _ := results["C"]
	test.Equal(weightC, 50.)

	weightD, _ := results["D"]
	test.Equal(weightD, 30.)

	weightE, _ := results["E"]
	test.Equal(weightE, 60.)
}
