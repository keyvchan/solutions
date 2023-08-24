package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}

type Position struct {
	x int
	y int
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	ways := map[Position]int{}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 0 {
				if i == 0 && j == 0 {
					// base case
					ways[Position{i, j}] = 1
					continue
				}
				ways[Position{i, j}] = ways[Position{i - 1, j}] + ways[Position{i, j - 1}]
			}
		}
	}

	return ways[Position{len(obstacleGrid) - 1, len(obstacleGrid[0]) - 1}]
}
