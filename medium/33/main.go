package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1, 3, 5}, 0))
	fmt.Println(search([]int{1, 3, 5}, 1))
}

func search(nums []int, target int) int {
	// restore order
	min := 0
	max := len(nums) - 1
	mid := 0

	shift := 0
	for {
		if min == max || max-min == 1 {
			if nums[min] < nums[max] {
				shift = min
			} else {
				shift = max
			}
			break
		}
		mid = (max + min) / 2

		if nums[mid] > nums[len(nums)-1] {
			min = mid
		}
		if nums[mid] < nums[len(nums)-1] {
			max = mid
		}

	}

	// binary search for the target
	min = 0
	max = len(nums) - 1
	mid = 0

	do_shift := func(index int) int {

		index = (index + shift) % len(nums)

		return index
	}

	// 4 5 6 7 0 1 2
	// 0 1 2 3 4 5 6
	// 3 4 5 6 0 1 2
	// 3 1
	// 0 1
	// 1 0
	// index < shift, index + shift - 1
	// index > shift, index - shift
	for {
		if min == max || max-min == 1 {
			if nums[do_shift(min)] == target {
				return do_shift(min)
			} else if nums[do_shift(max)] == target {
				return do_shift(max)
			} else {
				return -1
			}
		}
		mid = (min + max) / 2
		if nums[do_shift(mid)] < target {
			min = mid
		} else {
			max = mid
		}
	}

}
