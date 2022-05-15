package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{2, 3, 2}))
}

// 3 1 1 5 4 4 3 7 4 6

func rob(nums []int) int {
	counts_map := map[int]int{}

	counts_map[0] = nums[0]

	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}

	for i, num := range nums {
		counts_map[i] = max(counts_map[i-2]+num, counts_map[i-1])
	}

	return counts_map[len(nums)-1]
}
