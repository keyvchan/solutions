package main

import (
	"fmt"
)

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
}

func canBeShiped(weights []int, days int, capcity int) bool {
	// can current capacity ship all weights in days
	// if can, return true
	// if can't, return false

	current := capcity

	for i := 0; i < len(weights); i++ {
		if current < weights[i] {
			// we need one more day
			days--
			current = capcity
		}
		current -= weights[i]
	}

	return days > 0

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func shipWithinDays(weights []int, days int) int {

	// max of weights
	max := 0
	sum := 0
	for _, w := range weights {
		sum += w
		if w > max {
			max = w
		}
	}

	upper := sum
	lower := max

	result := sum

	for lower <= upper {
		mid := (upper + lower) / 2
		if canBeShiped(weights, days, mid) {
			result = min(result, mid)
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}
	return result
}
