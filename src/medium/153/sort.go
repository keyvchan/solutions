package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	// fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func findMin(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	return nums[0]
}
