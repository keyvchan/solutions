package main

import "fmt"

func main() {
	fmt.Println(findPeakGrid([][]int{
		{1, 4},
		{3, 2},
	}))

}

func findPeakGrid(mat [][]int) []int {
	// find peak in every row
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			left := false
			right := false
			check := false
			if j == 0 {
				left = true
			} else {
				left = mat[i][j-1] < mat[i][j]
			}
			if j == len(mat[i])-1 {
				right = true
			} else {
				right = mat[i][j+1] < mat[i][j]
			}
			if left && right {
				check = true
			}

			if check {
				// check 8 directions
				up := false
				down := false
				if i == 0 {
					up = true
				} else {
					up = mat[i-1][j] < mat[i][j]
				}
				if i == len(mat)-1 {
					down = true
				} else {
					down = mat[i+1][j] < mat[i][j]
				}
				if up && down {
					return []int{i, j}
				}
			}
		}
	}

	return []int{}
}
