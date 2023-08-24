package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMaxLen([]int{1, -2, -3, 4}))
	fmt.Println(getMaxLen([]int{0, 1, -2, -3, -4}))
	fmt.Println(getMaxLen([]int{-1, -2, -3, 0, 1}))
	fmt.Println(getMaxLen([]int{1, 2, 3, 5, -6, 4, 0, 10}))
	fmt.Println(getMaxLen([]int{5, -20, -20, -39, -5, 0, 0, 0, 36, -32, 0, -7, -10, -7, 21, 20, -12, -34, 26, 2}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxLen(nums []int) int {

	// split nums into two parts by 0
	splited := [][]int{}
	last_index := 0
	neg_counts := 0
	pos_counts := 0
	total := 0
	for i, num := range nums {
		if num == 0 {
			if last_index == i {
				last_index++
				continue
			}
			if neg_counts%2 != 0 {
				splited = append(splited, nums[last_index:i])
			} else {
				total = max(total, len(nums[last_index:i]))
			}
			last_index = i + 1
			neg_counts = 0
			pos_counts = 0
		} else {
			if num < 0 {
				neg_counts++
			} else {
				pos_counts++
			}
		}
		if i == len(nums)-1 {
			if neg_counts%2 != 0 {
				splited = append(splited, nums[last_index:])
			} else {
				total = max(total, len(nums[last_index:]))
			}
		}
	}

	for _, nums := range splited {
		// search from front to first negative number
		for i, num := range nums {
			if num < 0 {
				total = max(total, len(nums[i+1:]))
				break
			}
		}
		// search from back to first negative number
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] < 0 {
				total = max(total, len(nums[:i]))
				break
			}
		}

	}

	return total

}
