package main

import (
	"fmt"
	"io"
)

func main() {

	var n, m, k int

	for {
		_, err := fmt.Scan(&n, &m, &k)
		if err == io.EOF {
			break
		}
		var command string
		fmt.Scan(&command)

		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			grid[i] = make([]int, m)
		}

		result, num := solution(grid, command)
		fmt.Println(result, num)

	}

}

func solution(grid [][]int, commands string) (string, int) {
	grid[0][0] = 1
	cleaned := 1
	total := len(grid) * len(grid[0])
	x, y := 0, 0
	for i, command := range commands {
		switch command {
		case 'W':
			{
				x--
			}
		case 'A':
			{
				y--
			}
		case 'S':
			{
				x++
			}
		case 'D':
			{
				y++
			}
		}
		// we start at (0.0), then we execute the command
		if grid[x][y] == 0 {
			// we have not cleaned this cell yet
			cleaned++
			grid[x][y] = 1
			if cleaned >= total {
				return "Yes", i + 1
			}
		} else {
			if cleaned >= total {
				return "Yes", i
			}
		}
	}
	return "No", total - cleaned

}
