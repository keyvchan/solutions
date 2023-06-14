package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{}))
}

func dfs(num int, m map[int]bool) int {
	result := 0
	for m[num] {

		delete(m, num)
		result += 1
		num += 1
	}

	return result
}

func longestConsecutive(nums []int) int {
	// use a map to store all nums
	m := map[int]bool{}

	for _, num := range nums {
		m[num] = true
	}

	sort.Ints(nums)

	max := 0
	for _, num := range nums {
		result := dfs(num, m)
		if result > max {
			max = result
		}
	}
	return max
}
