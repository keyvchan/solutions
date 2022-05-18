package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
}

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	// make sure last not equal to first
	last := nums[0] - 1
	for _, num := range nums {
		if last == num {
			return true
		}
		last = num
	}

	return false
}
