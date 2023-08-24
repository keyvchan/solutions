package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMoves([]int{1, 2, 3}))
	fmt.Println(minMoves([]int{1, 1, 1}))
	fmt.Println(minMoves([]int{1, 1000000000}))
}

func minMoves(nums []int) int {
	// use a map to count nums frequency

	frequency := map[int]int{}
	for _, num := range nums {
		frequency[num]++
	}

	// increase all the nums by 1 except the largest one
	count := 0
	for len(frequency) != 1 {

		sort.Stable(sort.IntSlice(nums))
		for i := 0; i < len(nums)-1; i++ {
			frequency[nums[i]]--
			if frequency[nums[i]] == 0 {
				delete(frequency, nums[i])
			}
			nums[i]++
			frequency[nums[i]]++
		}
		count++
	}

	return count
}
