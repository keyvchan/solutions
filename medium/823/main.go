package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numFactoredBinaryTrees([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36}))
	fmt.Println(numFactoredBinaryTrees([]int{39, 7, 33, 27, 46, 30, 24, 6, 37, 11, 47, 44, 36, 31, 2, 48, 23, 32, 25, 86, 38, 9, 10, 14, 12, 26, 20, 34, 4, 28, 22, 18, 17, 45, 8, 21, 16, 13, 29, 42, 19, 41, 35, 49, 5, 43, 3, 144, 40, 15}))
}

func numFactoredBinaryTrees(arr []int) int {

	// brute force
	// we first take a number out of the array, and then we check if there are any other numbers that can be factored into it
	// create a map to store all numbers incase query its exists

	arrMap := map[int]bool{}
	for _, num := range arr {
		arrMap[num] = true
	}

	// we start from the smallest number
	sort.Ints(arr)

	dp := map[int]int{}
	total := 0
	for _, num := range arr {
		// single number can compose to a tree
		count := 1
		// TODO: for every number, we check if there are any other numbers that can be factored into it
		for _, otherNum := range arr {
			if otherNum >= num {
				break
			}
			if num%otherNum != 0 {
				// if the reminder is not 0, then we can't factorize it
				continue
			}

			// FIXME: integer division, should use float division
			// if _, ok := arrMap[num/otherNum]; ok {
			// we find a pair
			// we count the number of trees that can be composed from this pair
			count += (dp[num/otherNum] * dp[otherNum])
			// }
		}

		dp[num] = count
		total += count
	}
	fmt.Println(dp)

	return total % (1e9 + 7)
}
