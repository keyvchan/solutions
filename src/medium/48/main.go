package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	// transpose
	for i := 0; i < len(matrix); i++ {
		// first rot should be len(matrix)
		for j := i + 1; j < len(matrix); j++ {
			// swap
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// mirror
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}
	}

}
