package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))

	fmt.Println(minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

type Point struct {
	x int
	y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(grid [][]int, currentPoint Point, memo map[Point]int) int {

	// check if out of bounds
	if currentPoint.x < 0 || currentPoint.x >= len(grid) || currentPoint.y < 0 || currentPoint.y >= len(grid[0]) {
		return math.MaxInt32
	}
	if val, ok := memo[currentPoint]; ok {
		return val
	}

	if currentPoint.x == len(grid)-1 && currentPoint.y == len(grid[0])-1 {
		return grid[currentPoint.x][currentPoint.y]
	}

	sum := math.MaxInt32

	sum = min(sum, dfs(grid, Point{currentPoint.x + 1, currentPoint.y}, memo)+grid[currentPoint.x][currentPoint.y])
	sum = min(sum, dfs(grid, Point{currentPoint.x, currentPoint.y + 1}, memo)+grid[currentPoint.x][currentPoint.y])

	memo[currentPoint] = sum

	return sum

}

func minPathSum(grid [][]int) int {

	memo := map[Point]int{}
	return dfs(grid, Point{0, 0}, memo)
}
