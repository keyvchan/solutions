package main

func main() {

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(init int, graph map[int][]int, start int, visited map[int]bool, longest *int, length int) {
	// dfs
	if visited[start] {
		// cycle
		if start != init {
			return
		}
		*longest = max(*longest, length)
		return
	}
	visited[start] = true
	for _, v := range graph[start] {
		dfs(init, graph, v, visited, longest, length+1)
	}
}

func longestCycle(edges []int) int {
	// compose graph
	graph := map[int][]int{}

	for i, v := range edges {
		if v == -1 {
			continue
		}
		graph[i] = append(graph[i], v)
	}

	// find longest cycle
	longest := -1
	visited := map[int]bool{}
	// dfs
	for i := 0; i < len(edges); i++ {
		dfs(i, graph, i, visited, &longest, 0)
	}
	if longest == 0 {
		return -1
	}

	return longest
}
