package main

import (
	"fmt"
)

func main() {
	fmt.Println(movingCount(2, 3, 1))
	fmt.Println(movingCount(3, 1, 0))
}

func dfs(m int, n int, x int, y int, k int, visited map[Position]bool) int {
	if x < 0 || x >= m || y < 0 || y >= n || (x%10+x/10+y%10+y/10) > k {
		return 0
	}
	if _, ok := visited[Position{x, y}]; ok {
		return 0
	}
	result := 1
	visited[Position{x, y}] = true

	// up
	result += dfs(m, n, x-1, y, k, visited)
	// down
	result += dfs(m, n, x+1, y, k, visited)
	// left
	result += dfs(m, n, x, y-1, k, visited)
	// right
	result += dfs(m, n, x, y+1, k, visited)

	return result
}

type Position struct {
	x int
	y int
}

func movingCount(m int, n int, k int) int {
	visited := map[Position]bool{}
	return dfs(m, n, 0, 0, k, visited)

}
