package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(3, 7))
}

type Cell struct {
	x int
	y int
}

func uniquePaths(m int, n int) int {
	dp := map[Cell]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				// base case
				dp[Cell{x: 0, y: 0}] = 1
				continue
			}
			dp[Cell{i, j}] = dp[Cell{i - 1, j}] + dp[Cell{i, j - 1}]
		}
	}

	return dp[Cell{x: m - 1, y: n - 1}]
}
