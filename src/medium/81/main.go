package main

import "fmt"

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
}

func search(nums []int, target int) bool {
	// get the shift value
	last := -1
	shift := 0
	for i, num := range nums {
		if num < last {
			// we found the shift value
			shift = i
			break
		} else {
			last = num
		}
	}

	// binary search the value
	low := 0
	high := len(nums) - 1
	mid := (low + high) / 2
	for low <= high {
		fmt.Println(low, high, mid)
		mid = (low + high) / 2
		if nums[(mid+shift)%len(nums)] == target {
			return true
		}
		if nums[(mid+shift)%len(nums)] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
