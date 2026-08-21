package graph
import logritelog "log"

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestHasCycle(t=%v)", t)
	graph := Graph{Directed: true}
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}, {4, 0}}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}
	if !graph.HasCycle() {
		t.Error("answer of hasCycle is not correct")
	}

	graph = Graph{Directed: true}
	logritelog.Default().Printf("TestHasCycle(t=%v): graph → %v", t, graph)
	logritelog.Default().Printf("TestHasCycle(t=%v): edges → %v", t, edges)
	edges = [][]int{{0, 1}, {1, 2}, {2, 6}, {4, 0}}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}
	if graph.HasCycle() {
		t.Error("answer of hasCycle is not correct")
	}
	logritelog.Default().Printf("<<< Exiting TestHasCycle(t=%v)", t)
}

func TestFindAllCycles(t *testing.T) {
	logritelog.Default().Printf(">>> Entering TestFindAllCycles(t=%v)", t)
	graph := Graph{Directed: true}
	edges := [][]int{{0, 4}, {1, 3}, {2, 3}, {3, 4}, {4, 7}, {5, 2}, {6, 3}, {7, 3}}
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}

	res := graph.FindAllCycles()

	if len(res) != 1 {
		t.Error("number of cycles is not correct")
	}

	firstCycle := res[0]
	if len(firstCycle.edges) != 3 {
		t.Error("number of vertex in cycle is not correct")
	}
	if _, ok := firstCycle.edges[3][4]; !ok {
		t.Error("connection in cycle is not correct")
	}
	if _, ok := firstCycle.edges[4][7]; !ok {
		t.Error("connection in cycle is not correct")
	}
	if _, ok := firstCycle.edges[7][3]; !ok {
		t.Error("connection in cycle is not correct")
	}
	logritelog.Default().Printf("<<< Exiting TestFindAllCycles(t=%v)", t)
}
