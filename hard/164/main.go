package main

import "sort"

func main() {

}

func maximumGap(nums []int) int {
	// sort nums
	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)

	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > max {
			max = nums[i] - nums[i-1]
		}
	}
	return max

}
