package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4, 5}}))
	fmt.Println(spiralOrder([][]int{{1}, {2}, {3}, {4}, {5}}))
}

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0, len(matrix)*len(matrix[0]))
	first_row := 0
	last_row := len(matrix) - 1
	first_col := 0
	last_col := len(matrix[0]) - 1
	i, j := 0, 0
	for {
		if first_row > last_row || first_col > last_col {
			// fmt.Println(i, j)
			break
		}
		if i == first_row {
			// go right
			for j <= last_col {
				result = append(result, matrix[i][j])
				if j == last_col {
					first_row++
					break
				}
				j++
			}
			i++
		}
		if first_row > last_row || first_col > last_col {
			// fmt.Println(i, j)
			break
		}
		if j == last_col {
			// go down
			for i <= last_row {
				result = append(result, matrix[i][j])
				if i == last_row {
					last_col--
					break
				}
				i++
			}
			j--
		}
		if first_row > last_row || first_col > last_col {
			// fmt.Println(i, j)
			break
		}
		if i == last_row {
			// go left
			for j >= first_col {
				result = append(result, matrix[i][j])
				if j == first_col {
					last_row--
					break
				}
				j--
			}
			i--
		}
		if first_row > last_row || first_col > last_col {
			// fmt.Println(i, j)
			break
		}
		if j == first_col {
			// go up
			for i >= first_row {
				result = append(result, matrix[i][j])
				if i == first_row {
					first_col++
					break
				}
				i--
			}
			j++
		}
		if first_row > last_row || first_col > last_col {
			break
		}
	}

	return result
}
