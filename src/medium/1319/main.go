package main

func main() {

}
func dfs(node int, adjMap map[int][]int, visited map[int]bool) {
	visited[node] = true
	for _, neighbor := range adjMap[node] {
		if !visited[neighbor] {
			dfs(neighbor, adjMap, visited)
		}
	}
}

func makeConnected(n int, connections [][]int) int {
	connectedNodes := map[int]bool{}

	// construct the graph from connections
	adjMap := map[int][]int{}

	for _, connection := range connections {
		adjMap[connection[0]] = append(adjMap[connection[0]], connection[1])
		connectedNodes[connection[0]] = true
		adjMap[connection[1]] = append(adjMap[connection[1]], connection[0])
		connectedNodes[connection[1]] = true
	}

	if len(connections) < n-1 {
		return -1
	}

	// we find the number of connected components
	connectedComponents := 0
	globalVisitedNodes := map[int]bool{}

	for i := 0; i < n; i++ {
		if !globalVisitedNodes[i] {
			dfs(i, adjMap, globalVisitedNodes)
			// we merge the tempVisited nodes with the globalVisitedNodes
			connectedComponents++
		}
	}

	return connectedComponents - 1
}
