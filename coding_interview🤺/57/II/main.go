package main

import (
	"fmt"
)

func main() {
	fmt.Println(findContinuousSequence(9))
	fmt.Println(findContinuousSequence(15))
}

func sum(i, j int) int {
	// sum of range i to j
	return j*(j+1)/2 - (i-1)*(i)/2
}

func findContinuousSequence(target int) [][]int {

	result := [][]int{}

	// start from 1 to target
	i, j := 1, 1

	for i < target && j < target {

		// check if sum of range i to j is equal to target
		if sum(i, j) == target {
			temp := []int{}
			for k := i; k <= j; k++ {
				temp = append(temp, k)
			}
			result = append(result, temp)
			i++
			continue
		}
		if sum(i, j) > target {
			i++
			continue
		}
		if sum(i, j) < target {
			j++
			continue
		}

	}

	return result

}
