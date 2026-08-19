// cycle.go
// this file handle algorithm that related to cycle in graph
// time complexity: O(V+E) where V is the number of vertices and E is the number of edges in the graph
// space complexity: O(V) where V is the number of vertices in the graph
// reference: https://en.wikipedia.org/wiki/Cycle_(graph_theory)
// [kiarash hajian](https://github.com/kiarash8112)

package graph
import logritelog "log"

func (g *Graph) HasCycle() bool {
	logritelog.Default().Printf(">>> Entering HasCycle()")
	//this implimetation referred as 3-color too
	all := map[int]struct{}{}
	visiting := map[int]struct{}{}
	visited := map[int]struct{}{}

	for v := range g.edges {
		all[v] = struct{}{}
	}

	for current := range all {
		if g.hasCycleHelper(current, all, visiting, visited) {
			return true
		}
	}

	logritelog.Default().Printf("<<< Exiting HasCycle()")
	return false

}

func (g Graph) hasCycleHelper(v int, all, visiting, visited map[int]struct{}) bool {
	logritelog.Default().Printf(">>> Entering hasCycleHelper(v=%v,all=%v,visiting=%v,visited=%v)", v, all, visiting, visited)
	delete(all, v)
	visiting[v] = struct{}{}

	neighbors := g.edges[v]
	for v := range neighbors {
		if _, ok := visited[v]; ok {
			continue
		} else if _, ok := visiting[v]; ok {
			return true
		} else if g.hasCycleHelper(v, all, visiting, visited) {
			return true
		}
	}
	delete(visiting, v)
	visited[v] = struct{}{}
	logritelog.Default().Printf("<<< Exiting hasCycleHelper(v=%v,all=%v,visiting=%v,visited=%v)", v, all, visiting, visited)
	return false
}

// this function can do HasCycle() job but it is slower
func (g *Graph) FindAllCycles() []Graph {
	logritelog.Default().Printf(">>> Entering FindAllCycles()")
	all := map[int]struct{}{}
	visiting := map[int]struct{}{}
	visited := map[int]struct{}{}

	allCycles := []Graph{}

	for v := range g.edges {
		all[v] = struct{}{}
	}

	for current := range all {
		foundCycle, parents := g.findAllCyclesHelper(current, all, visiting, visited)

		if foundCycle {
			foundCycleFromCurrent := false
			//this loop remove additional vertex from detected cycle
			//using foundCycleFromCurrent bool to make sure after removing vertex we still have cycle
			for i := len(parents) - 1; i > 0; i-- {
				if parents[i][1] == parents[0][0] {
					parents = parents[:i+1]
					logritelog.Default().Printf("FindAllCycles(): parents → %v", parents)
					logritelog.Default().Printf("FindAllCycles(): foundCycleFromCurrent → %v", foundCycleFromCurrent)
					foundCycleFromCurrent = true
				}
			}
			if foundCycleFromCurrent {
				graph := Graph{Directed: true}
				for _, edges := range parents {
					graph.AddEdge(edges[1], edges[0])
				}
				allCycles = append(allCycles, graph)
			}
				logritelog.Default().Printf("FindAllCycles(): allCycles → %v", allCycles)

		}

	}

	logritelog.Default().Printf("<<< Exiting FindAllCycles()")
	return allCycles

}

func (g Graph) findAllCyclesHelper(current int, all, visiting, visited map[int]struct{}) (bool, [][]int) {
	logritelog.Default().Printf(">>> Entering findAllCyclesHelper(current=%v,all=%v,visiting=%v,visited=%v)", current, all, visiting, visited)
	parents := [][]int{}

	delete(all, current)
	visiting[current] = struct{}{}

	neighbors := g.edges[current]
	for v := range neighbors {
		if _, ok := visited[v]; ok {
			continue
		} else if _, ok := visiting[v]; ok {
			parents = append(parents, []int{v, current})
			logritelog.Default().Printf("findAllCyclesHelper(current=%v,all=%v,visiting=%v,visited=%v): parents → %v", current, all, visiting, visited, parents)
			return true, parents
		} else if ok, savedParents := g.findAllCyclesHelper(v, all, visiting, visited); ok {
			parents = append(parents, savedParents...)
			parents = append(parents, []int{v, current})
			return true, parents
		}
	}
	delete(visiting, current)
	visited[current] = struct{}{}
	logritelog.Default().Printf("<<< Exiting findAllCyclesHelper(current=%v,all=%v,visiting=%v,visited=%v)", current, all, visiting, visited)
	return false, parents
}
