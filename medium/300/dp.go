package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	// fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	// fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7}))
	// fmt.Println(lengthOfLIS([]int{1, 2, 3}))
	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	// start from back of the array
	dp := map[int]int{len(nums) - 1: 1}
	max := 1

	for i := len(nums) - 2; i >= 0; i-- {

		result := 1
		// if current number > last number, then we get max of rest
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				result = maxInt(result, dp[j]+1)
			}
		}
		if result > max {
			max = result
		}
		dp[i] = result
	}
	fmt.Println(dp)

	return max
}
