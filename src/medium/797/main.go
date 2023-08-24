package main

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{
		{1, 2},
		{3},
		{3},
		{},
	}))
	fmt.Println(allPathsSourceTarget([][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}))
}
func dfs(start int, end int, graph [][]int, current_path []int, result *[][]int) {
	if start == end {
		*result = append(*result, current_path)
		return
	}

	for _, next_node := range graph[start] {
		// copy current path
		new_current_path := make([]int, len(current_path))
		copy(new_current_path, current_path)
		new_current_path = append(new_current_path, next_node)
		dfs(next_node, end, graph, new_current_path, result)
	}

}

func allPathsSourceTarget(graph [][]int) [][]int {
	dest := len(graph) - 1
	src := 0

	result := [][]int{}
	dfs(src, dest, graph, []int{src}, &result)

	return result
}
