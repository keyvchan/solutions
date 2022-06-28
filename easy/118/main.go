package main

import "fmt"

func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(1))
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		if i > 1 {
			for j := 1; j < i; j++ {
				temp[j] = result[i-1][j-1] + result[i-1][j]
			}
		}
		result[i] = temp
	}
	return result
}
