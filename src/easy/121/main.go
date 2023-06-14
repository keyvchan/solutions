package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	// fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	// fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
}

func maxProfit(prices []int) int {
	least := prices[0]
	max_profit := 0
	for _, price := range prices {
		if price < least {
			least = price
		}
		if price-least > max_profit {
			max_profit = price - least
		}
	}

	return max_profit
}
