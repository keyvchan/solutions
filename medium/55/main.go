package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{1, 2, 0, 1}))
}

func canJump(nums []int) bool {

	require := 1

	for i := len(nums) - 2; i >= 0; i-- {
		fmt.Println(i, require)
		if nums[i] < require {
			if i == 0 {
				return false
			}
			require += 1
		} else {
			require = 1
		}
	}

	return true
}
