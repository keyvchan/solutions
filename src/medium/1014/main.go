package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
	fmt.Println(maxScoreSightseeingPair([]int{1, 2}))
	fmt.Println(maxScoreSightseeingPair([]int{2, 2, 2}))
	// 2 3 4
	// 2 1 0
}

func maxScoreSightseeingPair(values []int) int {

	max_left := -2147483648
	max_right := -2147483648
	max := 0

	for i := 1; i < len(values); i++ {
		// check i-1 can be used to make the left max
		if max_left < values[i-1]+i-1 {
			max_left = values[i-1] + i - 1
		}
		max_right = values[i] - i
		// check if the max is the max of the left and right
		if max < max_left+max_right {
			max = max_left + max_right
		}
	}

	return max
}
