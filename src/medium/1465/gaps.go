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
	// get the maxium gaps in horizontalCuts and verticalCuts
	maxgap_h := 0
	maxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sort.Ints(horizontalCuts)
	for i := 0; i <= len(horizontalCuts); i++ {
		if i == 0 {
			maxgap_h = maxInt(maxgap_h, horizontalCuts[i])
			continue
		}
		if i == len(horizontalCuts) {
			maxgap_h = maxInt(maxgap_h, h-horizontalCuts[i-1])
			continue
		}
		maxgap_h = maxInt(maxgap_h, horizontalCuts[i]-horizontalCuts[i-1])
	}
	var maxgap_v int
	sort.Ints(verticalCuts)
	for i := 0; i <= len(verticalCuts); i++ {
		if i == 0 {
			maxgap_v = maxInt(maxgap_h, verticalCuts[i])
			continue
		}
		if i == len(verticalCuts) {
			maxgap_v = maxInt(maxgap_h, w-verticalCuts[i-1])
			continue
		}
		maxgap_v = maxInt(maxgap_h, verticalCuts[i]-verticalCuts[i-1])
	}
	fmt.Println(maxgap_h, maxgap_v)

	// get the maxium area
	return (maxgap_h * maxgap_v) % 1000000007
}
