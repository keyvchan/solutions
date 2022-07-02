package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxArea(5, 4, []int{1, 2, 4}, []int{1, 3}))
	fmt.Println(maxArea(5, 4, []int{3, 1}, []int{1}))
	fmt.Println(maxArea(5, 4, []int{3}, []int{3}))
	fmt.Println(maxArea(1000000000, 1000000000, []int{2}, []int{2}))
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	// sort horizontalCuts and verticalCuts
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	max := 0
	last_h := 0

	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// for every horizontalCuts, compare all the areas
	for i := 0; i <= len(horizontalCuts); i++ {
		var current_h int
		if i == len(horizontalCuts) {
			current_h = h
		} else {
			current_h = horizontalCuts[i]
		}
		last_v := 0
		for j := 0; j <= len(verticalCuts); j++ {
			var current_v int
			if j == len(verticalCuts) {
				current_v = w
			} else {
				current_v = verticalCuts[j]
			}

			max = maxInt(max, (current_h-last_h)*(current_v-last_v))

			if j != len(verticalCuts) {
				last_v = verticalCuts[j]
			}

		}
		if i != len(horizontalCuts) {
			last_h = horizontalCuts[i]
		}
	}

	// max modulo 1000000007
	return max % 1000000007

}
