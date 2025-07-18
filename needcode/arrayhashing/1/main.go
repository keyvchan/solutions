package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	// use a set to emunerate all possible pairs

	numMap := map[int]int{}

	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{i, j}
		}
		numMap[num] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
