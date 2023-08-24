package main

import "fmt"

func main() {

	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
}

func maxProfit(prices []int) int {

	max := 0
	// sell out max
	sell_out := 0

	for i := len(prices) - 2; i >= 0; i-- {
		// i is the buy in day
		// check i+1 can be max sell out day
		if prices[i+1] > sell_out {
			sell_out = prices[i+1]
		}
		if sell_out-prices[i] > max {
			max = sell_out - prices[i]
		}
	}
	return max

}
