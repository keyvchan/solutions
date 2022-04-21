package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {

	max_profit := 0
	for i, price := range prices {
		if i == 0 {
			continue
		}
		if price > prices[i-1] {
			max_profit += price - prices[i-1]
		}
	}

	return max_profit
}
