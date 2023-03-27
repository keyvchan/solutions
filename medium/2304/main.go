package main

import "math"

func main() {

}

type Position struct {
	x int
	y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(grid [][]int, moveCost [][]int, currentPos Position, memo map[Position]int) int {
	// check if we are at the end
	if currentPos.x == len(grid)-1 {
		// we can return the cost, last row have zero cost
		return grid[currentPos.x][currentPos.y]
	}

	if cost, ok := memo[currentPos]; ok {
		return cost
	}

	cost := math.MaxInt32
	// we can move to the next row
	for i := 0; i < len(grid[0]); i++ {
		// we can move to the next row
		cost = min(cost, dfs(grid, moveCost, Position{currentPos.x + 1, i}, memo)+moveCost[grid[currentPos.x][currentPos.y]][i]+grid[currentPos.x][currentPos.y])
	}
	memo[currentPos] = cost
	return cost

}

func minPathCost(grid [][]int, moveCost [][]int) int {

	// so we use dfs to try every options
	memo := map[Position]int{}

	cost := math.MaxInt32
	for i := 0; i < len(grid[0]); i++ {
		cost = min(cost, dfs(grid, moveCost, Position{0, i}, memo))
	}

	return cost
}
