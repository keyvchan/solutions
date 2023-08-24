package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isStraight([]int{1, 2, 3, 4, 5}))
	fmt.Println(isStraight([]int{0, 0, 1, 2, 5}))
	fmt.Println(isStraight([]int{4, 7, 5, 9, 2}))
	fmt.Println(isStraight([]int{1, 2, 12, 0, 3}))
}

func isStraight(nums []int) bool {

	// count num frequency
	freq := map[int]bool{}
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if _, ok := freq[num]; ok {
			return false
		}
		freq[num] = true
	}

	sort.Ints(nums)
	// zeroCount
	zeroCount := 0
	i := 0
	for ; i < 4; i++ {
		if nums[i] == 0 {
			zeroCount++
		} else {
			break
		}
	}
	// start check from first non-zero number
	for j := nums[i]; j < nums[i]+5; j++ {
		if freq[j%14] {
			continue
		} else {
			zeroCount--
			if zeroCount < 0 {
				return false
			}
		}
	}

	return true

}
