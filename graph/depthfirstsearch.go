// depthfirstsearch.go
// description: this file contains the implementation of the depth first search algorithm
// details: Depth-first search (DFS) is an algorithm for traversing or searching tree or graph data structures. The algorithm starts at the root node and explores as far as possible along each branch before backtracking.
// time complexity: O(n)
// space complexity: O(n)

package graph
import logritelog "log"

func GetIdx(target int, nodes []int) int {
	logritelog.Default().Printf("[ERROR] >>> Entering GetIdx(target=%v,nodes=%v)", target, nodes)
	for i := 0; i < len(nodes); i++ {
		if nodes[i] == target {
			return i
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting GetIdx(target=%v,nodes=%v)", target, nodes)
	return -1
}

func NotExist(target int, slice []int) bool {
	logritelog.Default().Printf("[ERROR] >>> Entering NotExist(target=%v,slice=%v)", target, slice)
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return false
		}
	}
	logritelog.Default().Printf("[WARN] <<< Exiting NotExist(target=%v,slice=%v)", target, slice)
	return true
}

func DepthFirstSearchHelper(start, end int, nodes []int, edges [][]bool, showroute bool) ([]int, bool) {
	logritelog.Default().Printf("[ERROR] >>> Entering DepthFirstSearchHelper(start=%v,end=%v,nodes=%v,edges=%v,showroute=%v)", start, end, nodes, edges, showroute)
	logritelog.Default().Printf("[ERROR] >>> Entering DepthFirstSearchHelper(start=%v,end=%v,nodes=%v,edges=%v,showroute=%v)", start, end, nodes, edges, showroute)
	var route []int
	var stack []int
	startIdx := GetIdx(start, nodes)
	stack = append(stack, startIdx)
	for len(stack) > 0 {
	logritelog.Default().Printf("[DEBUG] DepthFirstSearchHelper(start=%v,end=%v,nodes=%v,edges=%v,showroute=%v): stack → %v", start, end, nodes, edges, showroute, stack)
		now := stack[len(stack)-1]
		route = append(route, nodes[now])
		if len(stack) > 1 {
		logritelog.Default().Printf("[DEBUG] DepthFirstSearchHelper(start=%v,end=%v,nodes=%v,edges=%v,showroute=%v): route → %v", start, end, nodes, edges, showroute, route)
			stack = stack[:len(stack)-1]
		} else {
			stack = stack[:len(stack)-1]
		}
		for i := 0; i < len(edges[now]); i++ {
			if edges[now][i] && NotExist(i, stack) {
				stack = append(stack, i)
			}
			edges[now][i] = false
			edges[i][now] = false
		}
		if route[len(route)-1] == end {
			return route, true
		}
	}

	if showroute {
		return route, false
	} else {
		logritelog.Default().Printf("[WARN] <<< Exiting DepthFirstSearchHelper(start=%v,end=%v,nodes=%v,edges=%v,showroute=%v)", start, end, nodes, edges, showroute)
		return nil, false
	}
}

func DepthFirstSearch(start, end int, nodes []int, edges [][]bool) ([]int, bool) {
	logritelog.Default().Printf("[ERROR] >>> Entering DepthFirstSearch(start=%v,end=%v,nodes=%v,edges=%v)", start, end, nodes, edges)
	logritelog.Default().Printf("[ERROR] >>> Entering DepthFirstSearch(start=%v,end=%v,nodes=%v,edges=%v)", start, end, nodes, edges)
	return DepthFirstSearchHelper(start, end, nodes, edges, false)
}

// func main() {
// 	nodes := []int{
// 		1, 2, 3, 4, 5, 6,
// 	}
// 	/*
// 		sample graph
// 		①-②
// 		|  |
// 		③-④-⑤-⑥
// 	*/
// 	edges := [][]bool{
// 		{false, true, true, false, false, false},
// 		{true, false, false, true, false, false},
// 		{true, false, false, true, false, false},
// 		{false, true, true, false, true, false},
// 		{false, false, false, true, false, true},
// 		{false, false, false, false, true, false},
// 	}
// 	start := 1
// 	end := 6
// 	route, _ := dfs(start, end, nodes, edges)
// 	fmt.Println(route)
// }
