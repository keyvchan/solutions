package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
	fmt.Println(getRow(4))
}

func getRow(rowIndex int) []int {
	last_row := []int{}

	for i := 0; i <= rowIndex; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		if i > 1 {
			for j := 1; j < i; j++ {
				temp[j] = last_row[j-1] + last_row[j]
			}
		}
		last_row = temp

	}

	return last_row
}
