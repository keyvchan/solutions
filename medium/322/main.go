package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange([]int{2}, 3))
}

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	amounts_map := make(map[int]int, amount+1)
	for _, coin := range coins {
		amounts_map[coin] = 1
	}
	amounts_map[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin > 0 {
				if _, ok := amounts_map[i]; !ok {
					if _, ok := amounts_map[i-coin]; ok {
						amounts_map[i] = amounts_map[i-coin] + amounts_map[coin]
					}
				} else {
					if _, ok := amounts_map[i-coin]; ok {
						if amounts_map[i] > amounts_map[i-coin]+amounts_map[coin] {
							amounts_map[i] = amounts_map[i-coin] + amounts_map[coin]
						}
					}
				}
			}
		}
	}

	if _, ok := amounts_map[amount]; !ok {
		return -1
	}
	return amounts_map[amount]
}
