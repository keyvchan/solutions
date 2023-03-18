package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
	fmt.Println(subsets([]int{1, 2, 3, 4, 5, 6, 7, 8, 10, 0}))
}

func backtracking(nums []int, current []int, all_subsets *[][]int) {
	// fmt.Println(current)

	if len(nums) == 0 {
		*all_subsets = append(*all_subsets, current)
		return
	}

	// choose
	choose_current := make([]int, len(current))
	copy(choose_current, current)
	backtracking(nums[1:], append(choose_current, nums[0]), all_subsets)

	// not choose
	backtracking(nums[1:], current, all_subsets)

}

func subsets(nums []int) [][]int {

	all_subsets := make([][]int, 0, 100)
	backtracking(nums, []int{}, &all_subsets)

	return all_subsets
}
