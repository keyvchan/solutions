package main

func main() {

}

func shiftGrid(grid [][]int, k int) [][]int {
	for i := 0; i < k; i++ {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				grid[i][j+1] = grid[i][j]
			}
			grid[i+1][0] = grid[i][len(grid[0])-1]
			grid[0][0] = grid[len(grid)-1][len(grid[0])-1]
		}
	}

	return grid
}
