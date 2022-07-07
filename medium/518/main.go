package main

import "fmt"

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(change(3, []int{2}))
	fmt.Println(change(10, []int{1, 5, 10}))
}

type Pair struct {
	amount int
	index  int
}

func backtracking(coins []int, current_index int, amount int, dp map[Pair]int) int {
	if amount == 0 {
		return 1
	}
	if amount < 0 {
		return 0
	}
	if _, ok := dp[Pair{amount, current_index}]; ok {
		return dp[Pair{amount, current_index}]
	}

	total := 0
	for i, coin := range coins[current_index:] {
		result := backtracking(coins, i+current_index, amount-coin, dp)
		total += result
	}
	dp[Pair{amount, current_index}] = total

	return total
}

func change(amount int, coins []int) int {
	dp := map[Pair]int{}

	return backtracking(coins, 0, amount, dp)
}
