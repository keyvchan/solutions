package main

func main() {

}

type Position struct {
	x int
	y int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(grid [][]int, position Position, memo map[Position]int) int {
	if position.x >= len(grid) || position.y >= len(grid[0]) {
		return 0
	}
	if val, ok := memo[position]; ok {
		return val
	}

	result := grid[position.x][position.y]
	result += max(dfs(grid, Position{position.x + 1, position.y}, memo), dfs(grid, Position{position.x, position.y + 1}, memo))

	memo[position] = result

	return result
}

func maxValue(grid [][]int) int {
	memo := map[Position]int{}

	return dfs(grid, Position{0, 0}, memo)

}
