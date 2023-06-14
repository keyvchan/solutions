package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3))
	fmt.Println(searchMatrix([][]int{{0}}, 0))
	fmt.Println(searchMatrix([][]int{{1}}, 1))
	fmt.Println(searchMatrix([][]int{{1, 3}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 11))
}

func searchMatrix(matrix [][]int, target int) bool {

	for i, row := range matrix {
		if target <= row[len(row)-1] {
			return binary_search(matrix[i], target)
		}
	}
	return false
}

func binary_search(a []int, target int) bool {
	if a[0] == target {
		return true
	}

	if len(a)/2 == 0 {
		return false
	}
	left := false
	right := false
	if a[len(a)/2] <= target {
		left = binary_search(a[len(a)/2:], target)
	} else {
		right = binary_search(a[:len(a)/2], target)
	}

	return left || right
}
