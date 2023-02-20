package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}

func searchInsert(nums []int, target int) int {

	// use binary search to find the target

	left, right := 0, len(nums)-1

	mid := 0
	for left <= right {
		mid = (left + right) / 2

		// check mid
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return left
}
