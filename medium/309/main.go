package main

import "fmt"

func main() {
	// fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

type Stock struct {
	index int
	buy   bool
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(prices []int, index int, buy bool, caching map[Stock]int) int {

	// base case
	if index > len(prices)-1 {
		// return
		return 0
	}
	if _, ok := caching[Stock{index, buy}]; ok {
		return caching[Stock{index, buy}]
	}

	profit := -2147483648
	if buy {
		// we have stock, so we can sell it, then enter a cooldown period
		profit = dfs(prices, index+1, !buy, caching) - prices[index]
		profit = max(profit, dfs(prices, index+1, buy, caching))
		caching[Stock{index, buy}] = profit
	} else {
		// not buy
		profit = dfs(prices, index+2, !buy, caching) + prices[index]
		profit = max(profit, dfs(prices, index+1, buy, caching))
		caching[Stock{index, buy}] = profit
	}

	// add current index to the cache

	return profit
}

func maxProfit(prices []int) int {
	// create a map caching the current stock status
	caching := map[Stock]int{}

	profit := dfs(prices, 0, true, caching)

	return profit

}
