package main

import (
	"fmt"
)

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
	fmt.Println(finalPrices([]int{1, 2, 3, 4, 5}))
	fmt.Println(finalPrices([]int{10, 1, 1, 6}))
	fmt.Println(finalPrices([]int{8, 7, 4, 2, 8, 1, 7, 7, 10, 1}))

}

type MyNode struct {
	prices int
	max    int
}

func finalPrices(prices []int) []int {
	// compose result
	result := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
			result[i] = prices[i]
		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				result[i] = prices[i] - prices[j]
				break
			}
		}
	}

	return result
}
