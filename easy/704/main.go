package main

import "fmt"

func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
}

func search(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for {
		mid := left + (right-left)/2
		if left == right {
			if nums[mid] != target {
				return -1
			}
		}
		if right-left == 1 {
			if nums[left] == target {
				return left
			}
			if nums[right] == target {
				return right
			}
			return -1
		}
		if nums[mid] > target {
			// search left
			right = mid

		} else if nums[mid] < target {
			// search right
			left = mid
		} else {
			// found it
			return mid
		}
	}

}
