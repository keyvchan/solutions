package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	fmt.Println(longestOnes([]int{0, 0, 0, 1}, 4))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestOnes(nums []int, k int) int {
	// create a sliding window
	left, right := 0, 0
	max := 0
	counts := k

	for right < len(nums) {
		if nums[right] == 1 {
			// nums[right] is 1, extend
			right++
		} else {
			// nums[right] is 0, check if we can extend
			if counts > 0 {
				// update max
				counts--
				right++
			} else {
				// counts < 0, we can't extend, shirk untail we can
				for nums[left] == 1 {
					left++
				}
				// we already know nums[left] is 0, so we can extend
				left++
				counts++
			}
		}
		if max < right-left {
			max = right - left
		}
	}

	return max
}
