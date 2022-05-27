package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
	// fmt.Println(subsets([]int{1, 2, 3, 4, 5, 6, 7, 8, 10, 0}))
}

func backtracking(nums []int, current_subset []int, all_subsets *[][]int) {
	temp := make([]int, len(current_subset))
	copy(temp, current_subset)
	*all_subsets = append(*all_subsets, temp)
	if len(nums) == 1 {
		current_subset = append(current_subset, nums[0])
		*all_subsets = append(*all_subsets, current_subset)
		return
	}

	for i, v := range nums {
		current_subset_temp := make([]int, len(current_subset)+1)
		copy(current_subset_temp, append(current_subset, v))
		backtracking(nums[i+1:], current_subset_temp, all_subsets)
	}

}

func subsets(nums []int) [][]int {

	all_subsets := [][]int{}

	backtracking(nums, []int{}, &all_subsets)

	return all_subsets
}
