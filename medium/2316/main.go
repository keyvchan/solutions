package main

func main() {

}

func dfs(graph map[int][]int, visited map[int]bool, node int) {
	if visited[node] {
		return
	}
	visited[node] = true

	for _, neighbor := range graph[node] {
		dfs(graph, visited, neighbor)
	}
}

func countPairs(n int, edges [][]int) int64 {

	// build graph
	graph := map[int][]int{}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	// count connected components
	components := []int{}

	visited := map[int]bool{}

	lastTotalNodes := 0
	for i := 0; i < n; i++ {
		// if node is visited, skip
		if visited[i] {
			continue
		}
		dfs(graph, visited, i)

		// newly visited nodes are in a new connected component
		components = append(components, len(visited)-lastTotalNodes)
		lastTotalNodes = len(visited)
	}

	if len(components) == 1 {
		return 0
	}

	result := int64(0)
	for i := 0; i < len(components); i++ {
		for j := i + 1; j < len(components); j++ {
			result += int64(components[i]) * int64(components[j])
		}
	}

	return result
}
