package main

import "sort"

func main() {

}

func minPairSum(nums []int) int {
	// sort nums
	sort.Ints(nums)

	left, right := 0, len(nums)-1

	// calcualte the max sum of each pair
	maxSum := 0
	for left < right {
		m := nums[left] + nums[right]
		if m > maxSum {
			maxSum = m
		}
		left++
		right--
	}

	return maxSum

}
