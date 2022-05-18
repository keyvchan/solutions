package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, target int) int {
	sum := []int{0}

	count := 0
	for _, num := range nums {
		for _, s := range sum {
			sum = append(sum, s+num)
			sum = append(sum, s-num)
			sum = sum[1:]
		}
	}

	for _, v := range sum {
		if v == target {
			count += 1
		}
	}
	return count
}
