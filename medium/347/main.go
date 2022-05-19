package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
	fmt.Println(topKFrequent([]int{-1, -1}, 1))
	fmt.Println(topKFrequent([]int{3, 0, 3, 0}, 1))
}

func topKFrequent(nums []int, k int) []int {
	count_map := make(map[int]int)

	for _, v := range nums {

		if _, ok := count_map[v]; ok {
			count_map[v] = count_map[v] + 1
		} else {
			count_map[v] = 1
		}
	}

	var vals []int
	for _, val := range count_map {
		vals = append(vals, val)
	}
	sort.Ints(vals)
	vals = vals[len(vals)-k:]

	var results []int
	for kk, v := range count_map {
		if len(results) >= k {
			break
		}
		for _, v1 := range vals {
			if v == v1 {
				results = append(results, kk)
				break
			}
		}
	}

	return results
}
