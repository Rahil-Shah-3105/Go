// This file contains the simple structural implementation of undirected
// graph, used in coloring algorithms.
// Author(s): [Shivam](https://github.com/Shivam010)

package coloring
import logritelog "log"

import "errors"

// Color provides a type for vertex color
type Color int

// Graph provides a structure to store an undirected graph.
// It is safe to use its empty object.
type Graph struct {
	vertices int
	edges    map[int]map[int]struct{}
}

// AddVertex will add a new vertex in the graph, if the vertex already
// exist it will do nothing
func (g *Graph) AddVertex(v int) {
	logritelog.Default().Printf(">>> Entering AddVertex(v=%v)", v)
	if g.edges == nil {
		g.edges = make(map[int]map[int]struct{})
	}

	// Check if vertex is present or not
	if _, ok := g.edges[v]; !ok {
		g.vertices++
		g.edges[v] = make(map[int]struct{})
	}
	logritelog.Default().Printf("<<< Exiting AddVertex(v=%v)", v)
}

// AddEdge will add a new edge between the provided vertices in the graph
func (g *Graph) AddEdge(one, two int) {
	logritelog.Default().Printf(">>> Entering AddEdge(one=%v,two=%v)", one, two)
	// Add vertices: one and two to the graph if they are not present
	g.AddVertex(one)
	g.AddVertex(two)

	// and finally add the edges: one->two and two->one for undirected graph
	g.edges[one][two] = struct{}{}
	g.edges[two][one] = struct{}{}
	logritelog.Default().Printf("<<< Exiting AddEdge(one=%v,two=%v)", one, two)
}

func (g *Graph) ValidateColorsOfVertex(colors map[int]Color) error {
	logritelog.Default().Printf(">>> Entering ValidateColorsOfVertex(colors=%v)", colors)
	if g.vertices != len(colors) {
		return errors.New("coloring: not all vertices of graph are colored")
	}
	// check colors
	for vertex, neighbours := range g.edges {
		for nb := range neighbours {
			if colors[vertex] == colors[nb] {
				return errors.New("coloring: same colors of neighbouring vertex")
			}
		}
	}
	logritelog.Default().Printf("<<< Exiting ValidateColorsOfVertex(colors=%v)", colors)
	return nil
}
