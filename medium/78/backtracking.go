package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
	fmt.Println(subsets([]int{1, 2, 3, 4, 5, 6, 7, 8, 10, 0}))
}

func backtracking(nums []int, current []int, all_subsets *[][]int, index int, length int) {

	// early return when length > what we need
	if len(current) > length {
		return
	}

	if len(current) == length {
		new_subset := make([]int, len(current))
		copy(new_subset, current)
		*all_subsets = append(*all_subsets, new_subset)
		return
	}

	for i := index; i < len(nums); i++ {
		current = append(current, nums[i])
		backtracking(nums, current, all_subsets, i+1, length)
		current = current[:len(current)-1]
	}

}

func subsets(nums []int) [][]int {

	all_subsets := make([][]int, 0, 100)

	for k := 0; k < len(nums)+1; k++ {
		backtracking(nums, []int{}, &all_subsets, 0, k)
	}

	return all_subsets
}
