package main

import "fmt"

func main() {
	fmt.Println(getMaximumGold([][]int{
		{0, 6, 0},
		{5, 8, 7},
		{0, 9, 0},
	}))
	fmt.Println(getMaximumGold([][]int{
		{1, 0, 7},
		{2, 0, 6},
		{3, 4, 5},
		{0, 3, 0},
		{9, 0, 20},
	}))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func backtracking(grid [][]int, position Position) int {
	// out of bounds
	if position.x < 0 || position.x >= len(grid) || position.y < 0 || position.y >= len(grid[position.x]) {
		return 0
	}
	if grid[position.x][position.y] == 0 {
		return 0
	}
	temp := grid[position.x][position.y]
	grid[position.x][position.y] = 0

	// get the max of the 4 directions
	// we can only go up, down, left, right
	m := 0
	m = max(m, backtracking(grid, Position{position.x + 1, position.y}))
	m = max(m, backtracking(grid, Position{position.x - 1, position.y}))
	m = max(m, backtracking(grid, Position{position.x, position.y + 1}))
	m = max(m, backtracking(grid, Position{position.x, position.y - 1}))

	grid[position.x][position.y] = temp

	return m + grid[position.x][position.y]
}

type Position struct {
	x int
	y int
}

func getMaximumGold(grid [][]int) int {
	result := 0
	// traverse the grid
	// we can start from any cell that has gold
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 {
				result = max(result, backtracking(grid, Position{i, j}))
			}
		}
	}

	return result
}
