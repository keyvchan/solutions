package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
	fmt.Println(minEatingSpeed([]int{312884470}, 312884469))
}

func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)

	min := 0
	max := piles[len(piles)-1]
	mid := 0
	for {
		if min == max || max-min == 1 {
			mid = max
			break
		}
		mid = (min + max) / 2
		// check hours for mid
		counts := 0
		for _, pile := range piles {
			counts += pile / mid
			if pile%mid != 0 {
				counts += 1
			}
		}
		if counts > h {
			min = mid
		} else {
			max = mid
		}
	}

	return mid

}
