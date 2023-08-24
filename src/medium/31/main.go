package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)
	nums = []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
	nums = []int{5, 1, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}
func non_increasing(nums []int) int {
	// start from back
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			// we found our point
			return i
		}
	}
	return -1
}
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextPermutation(nums []int) {
	// find the longest non-increasing part of the array
	inverse_point := non_increasing(nums)
	fmt.Println(inverse_point)

	if inverse_point != -1 {

		index := len(nums) - 1
		// find the next largest number
		for i := inverse_point + 1; i < len(nums)-1; i++ {
			if nums[i] > nums[inverse_point] && nums[i+1] <= nums[inverse_point] {
				index = i
				break
			}
		}
		// swap the two numbers
		nums[inverse_point], nums[index] = nums[index], nums[inverse_point]
	}
	// reverse the array
	reverse(nums[inverse_point+1:])

}
