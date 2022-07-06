package main

import "fmt"

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
}

func backtracking(coins []int, amount int, dp map[int]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		// this is not a valid solution
		return -1
	}
	if _, ok := dp[amount]; ok {
		return dp[amount]
	}

	max := 214783646
	for _, coin := range coins {
		steps := backtracking(coins, amount-coin, dp)
		if steps == -1 {
			continue
		} else {
			max = minInt(max, steps+1)
		}

	}
	dp[amount] = max

	return max
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp := map[int]int{}
	result := backtracking(coins, amount, dp)
	if result == 214783646 {
		return -1
	} else {
		return result
	}
}
