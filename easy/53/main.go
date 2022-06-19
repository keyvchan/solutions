package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	// base case
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	new_sum := 0

	for _, num := range nums {
		// check current number and max
		new_sum = num + new_sum
		if new_sum > max {
			// this is a contiguous subarray
			max = new_sum
		}
		if new_sum < 0 {
			// start from next number
			new_sum = 0
		}
	}

	return max

}
