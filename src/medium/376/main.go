package main

import "fmt"

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	fmt.Println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(wiggleMaxLength([]int{1, 1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(wiggleMaxLength([]int{3, 3, 3, 2, 5}))
}

func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return len(nums)
	}
	if len(nums) == 2 {
		if nums[0] == nums[1] {
			return 1
		} else {
			return 0
		}
	}
	dp := map[int]int{}

	dp[0] = 1
	if nums[0] == nums[1] {
		dp[1] = 1
	} else {
		dp[1] = 2
	}

	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	secondLast := nums[0]
	last := nums[1]
	for i := 2; i < len(nums); i++ {
		if nums[i] == last {
			// dp[i] don't change
			dp[i] = maxInt(1, dp[i-1])
		} else {
			// if nums[i] is the peek, then dp[i] = dp[i-1] + 1
			// if nums[i] is not the peek, we will skip nums[i]
			// check if nums[i] is the peek
			if (nums[i]-last)*(last-secondLast) < 0 {
				// its the peek
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = maxInt(2, dp[i-1])
			}
			secondLast = last
			last = nums[i]
		}

	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}
