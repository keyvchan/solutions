package main

import "fmt"

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))

}

type Position struct {
	x int
	y int
}

func dfs(grid [][]byte, i int, j int, visited_places map[Position]bool) {
	// check every direction
	// left, right, up, down
	// left
	if grid[i][j] == '0' {
		return
	}
	visited_places[Position{i, j}] = true
	if j > 0 && !visited_places[Position{i, j - 1}] {
		// j must > 0, otherwise it will be out of range
		// current position is '1' and not visited yet, which means we can dfs this direction
		dfs(grid, i, j-1, visited_places)
	}
	// right
	if j < len(grid[0])-1 && !visited_places[Position{i, j + 1}] {
		dfs(grid, i, j+1, visited_places)
	}
	// up
	if i > 0 && !visited_places[Position{i - 1, j}] {
		dfs(grid, i-1, j, visited_places)
	}
	// down
	if i < len(grid)-1 && !visited_places[Position{i + 1, j}] {
		dfs(grid, i+1, j, visited_places)
	}
}

func numIslands(grid [][]byte) int {

	visited_places := map[Position]bool{}
	counts := 0
	// loop through all the grid
	for i, row := range grid {
		for j, column := range row {
			if column == '0' {
				// current position is water, skip
				continue
			}
			if _, ok := visited_places[Position{i, j}]; ok {
				// current position is already visited, skip
				continue
			}
			// we finally found a new island, start bfs
			counts++
			dfs(grid, i, j, visited_places)
		}

	}
	return counts
}
