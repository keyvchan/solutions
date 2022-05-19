package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{}))
}

func longestConsecutive(nums []int) int {
	// sort

	// make a map
	num_map := map[int]int{}

	for _, num := range nums {
		num_map[num] += 1
	}

	count := 0
	for _, num := range nums {
		// check
		if _, ok := num_map[num-1]; ok {
			continue
		}
		initial_num := num
		sub_count := 1
		for {
			if _, ok := num_map[initial_num+1]; ok {
				sub_count += 1
				initial_num += 1
			} else {
				if count < sub_count {
					count = sub_count
				}
				break
			}
		}
	}

	return count

}
