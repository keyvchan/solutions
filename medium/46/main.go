package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func remove_index(nums []int, i int) []int {
	new_nums := []int{}
	for index, v := range nums {
		if index != i {
			new_nums = append(new_nums, v)
		}
	}
	return new_nums
}

func rest(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	if len(nums) == 0 {
		return [][]int{}
	}
	result := [][]int{}
	for i, v := range nums {
		new_nums := remove_index(nums, i)
		sub_permutations := rest(new_nums)
		for _, sub_permutation := range sub_permutations {
			result = append(result, append([]int{v}, sub_permutation...))
		}
	}
	return result
}

func permute(nums []int) [][]int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return rest(nums)
}
