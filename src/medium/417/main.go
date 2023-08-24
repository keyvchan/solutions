package main

import "fmt"

func main() {
	fmt.Println([][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	})
}

func dfs(heights [][]int, i, j int, visited [][]bool, pacific bool) bool {
	// check if the point is in the boundary
	if pacific {
		if i == 0 || j == 0 {
			return true
		}
	} else {
		if i == len(heights)-1 || j == len(heights[0])-1 {
			return true
		}
	}
	result := false
	visited[i][j] = true
	// left
	if i > 0 && heights[i-1][j] <= heights[i][j] && !visited[i-1][j] {
		result = result || dfs(heights, i-1, j, visited, pacific)
	}
	if i < len(heights)-1 && heights[i+1][j] <= heights[i][j] && !visited[i+1][j] {
		result = result || dfs(heights, i+1, j, visited, pacific)
	}
	if j > 0 && heights[i][j-1] <= heights[i][j] && !visited[i][j-1] {
		result = result || dfs(heights, i, j-1, visited, pacific)
	}
	if j < len(heights[0])-1 && heights[i][j+1] <= heights[i][j] && !visited[i][j+1] {
		result = result || dfs(heights, i, j+1, visited, pacific)
	}
	return result
}

func pacificAtlantic(heights [][]int) [][]int {

	result := [][]int{}
	empty := make([][]bool, len(heights))
	for i := 0; i < len(heights); i++ {
		empty[i] = make([]bool, len(heights[0]))
	}
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			// copy empty
			empty := make([][]bool, len(heights))
			for u := 0; u < len(heights); u++ {
				empty[u] = make([]bool, len(heights[0]))
			}

			// check every point can be reached from pacific
			pacific := dfs(heights, i, j, empty, true)
			// check every point can be reached from atlantic
			empty = make([][]bool, len(heights))
			for v := 0; v < len(heights); v++ {
				empty[v] = make([]bool, len(heights[0]))
			}
			atlantic := dfs(heights, i, j, empty, false)
			if pacific && atlantic {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}
