package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
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

func dfs(prices []int, index int, fee int, buy bool, caching map[Stock]int) int {
	if index >= len(prices) {
		return 0
	}
	if val, ok := caching[Stock{index, buy}]; ok {
		return val
	}

	if buy {
		// we can buy
		profit1 := dfs(prices, index+1, fee, !buy, caching) - prices[index]

		// we not buy, the max profit of rest donest change
		profit2 := dfs(prices, index+1, fee, buy, caching)

		caching[Stock{index, buy}] = max(profit1, profit2)
	} else {
		// we can sell
		profit1 := dfs(prices, index+1, fee, !buy, caching) + prices[index] - fee

		// we doesn't sell the max profit of rest donest change
		profit2 := dfs(prices, index+1, fee, buy, caching)

		caching[Stock{index, buy}] = max(profit1, profit2)

	}

	return caching[Stock{index, buy}]

}

func maxProfit(prices []int, fee int) int {

	caching := map[Stock]int{}

	return dfs(prices, 0, fee, true, caching)

}
