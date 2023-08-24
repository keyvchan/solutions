package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}))
}

func orangesRotting(grid [][]int) int {
	// rotten counts
	var rottenCount int

	queue := list.New()

	// find all rotten oranges in a grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			switch grid[i][j] {
			case 2:
				queue.PushBack([]int{i, j, 0})
			case 1:
				rottenCount++
			}
		}
	}
	// do first iteration
	counts := 0
	for queue.Len() > 0 {
		// get the first element
		element := queue.Remove(queue.Front())
		// get the coordinates
		// set left right top bottom
		coordinates := element.([]int)
		if coordinates[2] > counts {
			counts = coordinates[2]
		}
		if coordinates[0] > 0 && grid[coordinates[0]-1][coordinates[1]] == 1 {
			grid[coordinates[0]-1][coordinates[1]] = 2
			queue.PushBack([]int{coordinates[0] - 1, coordinates[1], coordinates[2] + 1})
			rottenCount--
		}
		if coordinates[0] < len(grid)-1 && grid[coordinates[0]+1][coordinates[1]] == 1 {
			grid[coordinates[0]+1][coordinates[1]] = 2
			queue.PushBack([]int{coordinates[0] + 1, coordinates[1], coordinates[2] + 1})
			rottenCount--
		}
		if coordinates[1] > 0 && grid[coordinates[0]][coordinates[1]-1] == 1 {
			grid[coordinates[0]][coordinates[1]-1] = 2
			queue.PushBack([]int{coordinates[0], coordinates[1] - 1, coordinates[2] + 1})
			rottenCount--
		}
		if coordinates[1] < len(grid[0])-1 && grid[coordinates[0]][coordinates[1]+1] == 1 {
			grid[coordinates[0]][coordinates[1]+1] = 2
			queue.PushBack([]int{coordinates[0], coordinates[1] + 1, coordinates[2] + 1})
			rottenCount--
		}

	}

	if rottenCount > 0 {
		return -1
	}
	return counts
}
