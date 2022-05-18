package main

import (
	"fmt"
)

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 1, 2, 2}))
}

// dp[i] = dp[i-1] + v
// dp[rest] = dp[i] - v

func canPartition(nums []int) bool {

	sum := 0
	// sum of all nums
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}

	dp := map[int]bool{0: true}

	for _, num := range nums {
		newdp := map[int]bool{}
		for k := range dp {
			if k+num == sum/2 {
				return true
			} else {
				newdp[k+num] = true
			}
		}
		// merge two dp
		for k, v := range newdp {
			dp[k] = v
		}
	}

	return false
}
