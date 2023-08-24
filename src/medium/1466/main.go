package main

func main() {

}

func dfs(node int, graph map[int][]int, visited map[int]bool, counts *int) {
	if visited[node] {
		return
	}
	visited[node] = true

	for _, next := range graph[node] {
		if visited[next] {
			continue
		}
		if next > 0 {
			*counts++
			dfs(next, graph, visited, counts)
		} else {
			dfs(-next, graph, visited, counts)
		}
	}

}

func minReorder(n int, connections [][]int) int {
	// construct the graph
	graph := map[int][]int{}

	for _, c := range connections {
		graph[c[0]] = append(graph[c[0]], c[1])
		graph[c[1]] = append(graph[c[1]], -c[0])
	}

	var counts int
	visited := map[int]bool{}
	dfs(0, graph, visited, &counts)

	return counts
}
