package main

import "fmt"

func main() {
	fmt.Println(maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}))
}

func dfs(grid [][]int, i, j int, visited [][]bool, counts *int) {
	if grid[i][j] == 0 {
		return
	}
	*counts += 1
	visited[i][j] = true

	// left
	if j > 0 && !visited[i][j-1] {
		dfs(grid, i, j-1, visited, counts)
	}
	// up
	if i > 0 && !visited[i-1][j] {
		dfs(grid, i-1, j, visited, counts)
	}
	// right
	if j < len(grid[0])-1 && !visited[i][j+1] {
		dfs(grid, i, j+1, visited, counts)
	}
	// down
	if i < len(grid)-1 && !visited[i+1][j] {
		dfs(grid, i+1, j, visited, counts)
	}

}

func maxAreaOfIsland(grid [][]int) int {

	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			if visited[i][j] {
				continue
			}
			counts := 0
			dfs(grid, i, j, visited, &counts)
			if counts > max {
				max = counts
			}
		}
	}

	return max
}
