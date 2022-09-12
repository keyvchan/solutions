package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
	fmt.Println(checkPossibility([]int{5, 7, 1, 8}))
}

type Node struct {
	Value int
	Max   int
}

func checkPossibility(nums []int) bool {
	changed := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if changed {
				// we need to change again, not possible
				return false
			}
			changed = true
			if i == 1 || nums[i] >= nums[i-2] {
				// we change i-1
				nums[i-1] = nums[i]
			} else {
				// we change i
				nums[i] = nums[i-1]
			}
		}
	}
	return true
}
