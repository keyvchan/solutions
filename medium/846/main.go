package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(isNStraightHand([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(isNStraightHand([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(isNStraightHand([]int{1, 1, 2, 2, 3, 3}, 3))
	fmt.Println(isNStraightHand([]int{1, 2, 3, 3, 4, 4, 5, 6}, 4))
	fmt.Println(isNStraightHand([]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3))
}

func isNStraightHand(hand []int, groupSize int) bool {

	// check len of hand
	if len(hand)%groupSize != 0 {
		return false
	}

	// add hand to map
	handMap := map[int]int{}
	for _, v := range hand {
		handMap[v]++
	}

	// check if it can construct the consecutive hand
	sort.Ints(hand)

	for _, v := range hand {
		if handMap[v] == 0 {
			continue
		}
		for i := 0; i < groupSize; i++ {
			if handMap[v] <= 0 {
				// we can't find the number
				return false
			} else {
				handMap[v]--
			}
			v++
		}
	}

	return true
}
