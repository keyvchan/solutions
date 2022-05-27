package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{4, 4, 4, 1, 4}))
}

func subsetsWithDup(nums []int) [][]int {

	sort.Ints(nums)
	result := [][]int{}

	backtracking(nums, []int{}, -20, &result)

	return result
}

func backtracking(nums []int, current []int, last int, result *[][]int) {

	if len(nums) == 0 {
		*result = append(*result, current)
		return
	}

	new_current := make([]int, len(current), len(current)+1)
	copy(new_current, current)
	new_current = append(new_current, nums[0])
	backtracking(nums[1:], new_current, nums[0], result)

	if nums[0] != last || last < -10 {
		backtracking(nums[1:], current, last, result)
	}

}
