package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Solution([]int{4, 8, 2, 2, 1, 4}, []int{3, 6, 5, 2, 4, 1}, 3))
}

func Solution(X []int, Y []int, W int) int {
	// only sort X
	sort.Sort(sort.IntSlice(X))
	// start from 1
	counts := 0
	left := 0
	for i, x := range X {
		fmt.Println(i, x, left, counts)
		if i == 0 {
			left = x
			counts++
		}
		if x-left > W {
			left = x
			counts++
		}

	}
	return counts

}
