package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Solution([]int{-3, -2, 1, 0, 8, 7, 1}, 3))
	fmt.Println(Solution([]int{7, 1, 11, 8, 4, 10}, 8))
}

func Solution(A []int, M int) int {
	// write your code in Go (Go 1.4)
	// check every pair
	sort.Sort(sort.IntSlice(A))
	dp := map[int]int{}

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if (A[j]-A[i])%M == 0 {
				fmt.Println(i, j)
				dp[j]++
				dp[i]++
			}
		}
		dp[i]++
	}

	// find max value
	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}

	return max
}
