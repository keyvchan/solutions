package main

import (
	"fmt"
	"io"
)

func main() {

	var n, m int
	for {
		_, err := fmt.Scan(&n, &m)
		if err == io.EOF {
			break
		}
		if n == 0 && m == 0 {
			break
		}

		matrix := make([][]int, n)
		for i := 0; i < n; i++ {
			matrix[i] = make([]int, m)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Scan(&matrix[i][j])
			}
		}

		fmt.Println(Solution(matrix))
	}

}

type Position struct {
	x int
	y int
}

func dfs(matrix [][]int, p Position, visited map[Position]bool) int {
	// check the boundary
	if p.x < 0 || p.x >= len(matrix) || p.y < 0 || p.y >= len(matrix[0]) {
		return 2147483647
	}
	if matrix[p.x][p.y] == 3 {
		return 0
	}

	if _, ok := visited[p]; ok {
		return 2147483647
	}
	// check current position
	if matrix[p.x][p.y] == 1 {
		return 2147483647
	}
	costs := 1
	if matrix[p.x][p.y] == 4 {
		costs = 3
	}
	if matrix[p.x][p.y] == 6 {
		// alter the matrix
		// create a new matrix
		newMatrix := make([][]int, len(matrix))
		for i := 0; i < len(matrix); i++ {
			newMatrix[i] = make([]int, len(matrix[0]))
		}
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				newMatrix[i][j] = matrix[i][j]
			}
		}
		// up
		if p.x-1 >= 0 && newMatrix[p.x-1][p.y] == 1 {
			newMatrix[p.x-1][p.y] = 0
		}
		// down
		if p.x+1 < len(matrix) && newMatrix[p.x+1][p.y] == 1 {
			newMatrix[p.x+1][p.y] = 0
		}
		// left
		if p.y-1 >= 0 && newMatrix[p.x][p.y-1] == 1 {
			newMatrix[p.x][p.y-1] = 0
		}
		// right
		if p.y+1 < len(matrix[0]) && newMatrix[p.x][p.y+1] == 1 {
			newMatrix[p.x][p.y+1] = 0
		}

		costs = 1
		matrix = newMatrix
	}
	visited[p] = true
	// create a new map
	newVisited := map[Position]bool{}
	for k, v := range visited {
		newVisited[k] = v
	}
	// up
	up := dfs(matrix, Position{p.x - 1, p.y}, newVisited) + costs
	// down
	// create a new map
	newVisited = map[Position]bool{}
	for k, v := range visited {
		newVisited[k] = v
	}

	down := dfs(matrix, Position{p.x + 1, p.y}, newVisited) + costs
	// left
	// create a new map
	newVisited = map[Position]bool{}
	for k, v := range visited {
		newVisited[k] = v
	}
	left := dfs(matrix, Position{p.x, p.y - 1}, newVisited) + costs
	// right
	// create a new map
	newVisited = map[Position]bool{}
	for k, v := range visited {
		newVisited[k] = v
	}
	right := dfs(matrix, Position{p.x, p.y + 1}, newVisited) + costs

	// check the min
	min := 2147483647
	if up < min {
		min = up
	}
	if down < min {
		min = down
	}
	if left < min {
		min = left
	}
	if right < min {
		min = right
	}
	return min

}

func Solution(matrix [][]int) int {
	visited := map[Position]bool{}
	// search the soldier Position
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 2 {
				return dfs(matrix, Position{i, j}, visited)
			}
		}
	}

	return 0
}
