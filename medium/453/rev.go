package main

import "sort"

func main() {

}

func minMoves(nums []int) int {
	sort.Ints(nums)

	min := nums[0]

	count := 0
	for _, num := range nums {
		count += num - min
	}

	return count
}
