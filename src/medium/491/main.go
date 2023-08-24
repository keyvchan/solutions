package main

import "fmt"

func main() {
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
}

// leetcode 491
// https://leetcode.com/problems/increasing-subsequences/

func backtracking(nums []int, start int, current []int, result *[][]int) {
	if len(current) >= 2 {
		*result = append(*result, append([]int{}, current...))
	}

	used := map[int]bool{}
	for i := start; i < len(nums); i++ {
		if used[nums[i]] {
			continue
		}

		if len(current) == 0 || nums[i] >= current[len(current)-1] {
			used[nums[i]] = true
			current = append(current, nums[i])
			backtracking(nums, i+1, current, result)
			current = current[:len(current)-1]
		}
	}
}

func findSubsequences(nums []int) [][]int {
	// for every number, we could either include it or not
	result := [][]int{}
	backtracking(nums, 0, []int{}, &result)

	return result

}
