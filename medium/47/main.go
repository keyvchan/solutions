package main

import "sort"

func main() {

}

func backtracking(nums []int, current []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, current)
	}

	for i, num := range nums {
		// skip duplicate
		if i > 0 && nums[i-1] == num {
			continue
		}
		// backtracking with current + num
		// create a new slice to avoid mutating the original slice
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		newNums = append(newNums[:i], newNums[i+1:]...)
		backtracking(newNums, append(current, num), result)

	}

}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	backtracking(nums, []int{}, &result)

	return result
}
