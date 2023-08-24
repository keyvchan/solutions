package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, -1, 0, -3, -3}))
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	left_array := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			left_array[i] = 1
		} else {
			left_array[i] = left_array[i-1] * nums[i-1]
		}
	}

	right_array := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			right_array[i] = 1
		} else {
			right_array[i] = right_array[i+1] * nums[i+1]
		}
	}

	for i := range nums {
		result[i] = left_array[i] * right_array[i]
	}

	return result
}
