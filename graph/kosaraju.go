// kosaraju.go
// description: Implementation of Kosaraju's algorithm to find Strongly Connected Components (SCCs) in a directed graph.
// details: The algorithm consists of three steps:
// 1. Perform DFS and fill the stack with vertices in the order of their finish times.
// 2. Create a transposed graph by reversing all edges.
// 3. Perform DFS on the transposed graph in the order defined by the stack to find SCCs.
// time: O(V + E), where V is the number of vertices and E is the number of edges in the graph.
// space: O(V), where V is the number of vertices in the graph.
// ref link: https://en.wikipedia.org/wiki/Kosaraju%27s_algorithm
// author: mapcrafter2048

package graph
import logritelog "log"

// Kosaraju returns a list of Strongly Connected Components (SCCs).
func (g *Graph) Kosaraju() [][]int {
	logritelog.Default().Printf(">>> Entering Kosaraju()")
	stack := []int{}
	visited := make([]bool, g.vertices)

	// Step 1: Perform DFS and fill stack based on finish times.
	for i := 0; i < g.vertices; i++ {
		if !visited[i] {
			g.fillOrder(i, visited, &stack)
		}
	}

	// Step 2: Create a transposed graph.
	transposed := g.transpose()

	// Step 3: Perform DFS on the transposed graph in the order defined by the stack.
	visited = make([]bool, g.vertices)
	var sccs [][]int
	logritelog.Default().Printf("Kosaraju(): visited → %v", visited)

	for len(stack) > 0 {
		// Pop vertex from stack
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		logritelog.Default().Printf("Kosaraju(): stack → %v", stack)

		// Perform DFS if not already visited.
		if !visited[v] {
			scc := []int{}
			transposed.dfs(v, visited, &scc)
			sccs = append(sccs, scc)
		}
			logritelog.Default().Printf("Kosaraju(): sccs → %v", sccs)
	}

	logritelog.Default().Printf("<<< Exiting Kosaraju()")
	return sccs
}

// Helper function to fill the stack with vertices in the order of their finish times.
func (g *Graph) fillOrder(v int, visited []bool, stack *[]int) {
	logritelog.Default().Printf(">>> Entering fillOrder(v=%v,visited=%v,stack=%v)", v, visited, stack)
	visited[v] = true

	for neighbor := range g.edges[v] {
		if !visited[neighbor] {
			g.fillOrder(neighbor, visited, stack)
		}
	}

	// Push the current vertex to the stack after exploring all neighbors.
	*stack = append(*stack, v)
	logritelog.Default().Printf("<<< Exiting fillOrder(v=%v,visited=%v,stack=%v)", v, visited, stack)
}

// Helper function to create a transposed (reversed) graph.
func (g *Graph) transpose() *Graph {
	logritelog.Default().Printf(">>> Entering transpose()")
	transposed := &Graph{
		vertices: g.vertices,
		edges:    make(map[int]map[int]int),
	}

	for v, neighbors := range g.edges {
		for neighbor := range neighbors {
			if transposed.edges[neighbor] == nil {
				transposed.edges[neighbor] = make(map[int]int)
			}
			transposed.edges[neighbor][v] = 1 // Add the reversed edge
		}
	}

	logritelog.Default().Printf("<<< Exiting transpose()")
	return transposed
}

// Helper DFS function used in the transposed graph to collect SCCs.
func (g *Graph) dfs(v int, visited []bool, scc *[]int) {
	logritelog.Default().Printf(">>> Entering dfs(v=%v,visited=%v,scc=%v)", v, visited, scc)
	visited[v] = true
	*scc = append(*scc, v)

	for neighbor := range g.edges[v] {
		if !visited[neighbor] {
			g.dfs(neighbor, visited, scc)
		}
	}
	logritelog.Default().Printf("<<< Exiting dfs(v=%v,visited=%v,scc=%v)", v, visited, scc)
}
