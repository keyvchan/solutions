package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(prices []int, index int, buyed bool, counts int, memo map[State]int) int {
	if index == len(prices) {
		return 0
	}
	if val, ok := memo[State{index, buyed, counts}]; ok {
		return val
	}
	profit := 0
	if !buyed {
		// we are not hold any stock, we can buy or do nothing
		// buy
		if counts < 2 {
			// we can only buy 2 stocks
			profit = max(profit, dfs(prices, index+1, true, counts+1, memo)-prices[index])
		}
		// do nothing
		profit = max(profit, dfs(prices, index+1, buyed, counts, memo))

	} else {
		// when buyed is true, we can only sell or not sell
		// sell
		profit = max(profit, dfs(prices, index+1, false, counts, memo)+prices[index])
		// not sell
		profit = max(profit, dfs(prices, index+1, buyed, counts, memo))

	}
	memo[State{index, buyed, counts}] = profit

	return profit

}

type State struct {
	index  int
	buyed  bool
	counts int
}

func maxProfit(prices []int) int {
	memo := map[State]int{}

	return dfs(prices, 0, false, 0, memo)
}
