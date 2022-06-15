package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	// fmt.Println(maxProduct([]int{-2, 3, -4}))
	// fmt.Println(maxProduct([]int{-2, -3, -4}))
	// fmt.Println(maxProduct([]int{-2, -3, -4, -5}))
}

func maxProduct(nums []int) int {
	cur_max := 1
	cur_min := 1

	max := func(a, b, c int) int {
		if a > b {
			if a > c {
				return a
			}
			return c
		}
		if b > c {
			return b
		}
		return c
	}
	min := func(a, b, c int) int {
		if a < b {
			if a < c {
				return a
			}
			return c
		}
		if b < c {
			return b
		}
		return c
	}
	actually_max := 0

	for i := 0; i < len(nums); i++ {
		// keep tracking the max and min
		new_cur_max := cur_max
		cur_max = max(cur_max*nums[i], cur_min*nums[i], nums[i])
		cur_min = min(new_cur_max*nums[i], cur_min*nums[i], nums[i])
		if cur_max > actually_max {
			actually_max = cur_max
		}
	}

	return actually_max
}
