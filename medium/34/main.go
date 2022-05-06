package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{1}, 1))
	fmt.Println(searchRange([]int{2, 2}, 2))
	fmt.Println(searchRange([]int{1, 4}, 4))
	fmt.Println(searchRange([]int{-3, -2, -1}, 0))
}

func searchRange(nums []int, target int) []int {

	// check nums len is 0
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := 0
	right := len(nums) - 1
	mid := 0

	// find first
	for {
		mid = left + (right-left)/2
		if left == right {
			if nums[left] == target {
				break
			} else {
				return []int{-1, -1}
			}
		}
		if mid == 0 {
			if len(nums) > 1 {
				if nums[mid] == target {
					break
				}
				if nums[mid+1] != target {
					mid = -1
				} else {
					mid = mid + 1
				}
				break
			}
			// check mid + 1
			if nums[mid] != target {
				mid = -1
			}
			break
		}
		if nums[mid] < target {
			// position must be right
			left = mid + 1
		}
		if nums[mid] > target {
			// position must be left
			right = mid - 1
		}

		if nums[mid] == target {

			// check mid - 1
			if nums[mid-1] == target {
				right = mid - 1
			}
			if nums[mid-1] != target {
				// mid is the first position
				break
			}
		}
	}

	// we got the mid
	if mid == -1 {
		return []int{-1, -1}
	}

	right = mid
	// search from mid to right
	for i := mid; i < len(nums); i++ {
		if nums[i] == target {
			right = i
			continue
		}
		break
	}

	return []int{mid, right}
}
