package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 2, 3, 5}, 3))
	fmt.Println(kthSmallestPrimeFraction([]int{1, 7}, 1))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	combinations := [][]int{}
	// list all conbinations of prime fractions
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			combinations = append(combinations, []int{arr[i], arr[j]})
		}
	}
	sort.Slice(combinations, func(i, j int) bool {
		// make sure the second element is equal
		first := combinations[i][0] * combinations[j][1]
		second := combinations[j][0] * combinations[i][1]
		return first < second
	})
	return combinations[k-1]

}
