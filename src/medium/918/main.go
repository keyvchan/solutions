package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}))
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))
	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))
}

func maxSubarraySumCircular(nums []int) int {
	// max(i+1, i)
	// get max sum of non-circular array
	max_non_circular := maxSubArray(nums, true)

	total_sum := 0
	for _, num := range nums {
		total_sum += num
	}
	max_circular := total_sum - maxSubArray(nums, false)
	fmt.Println(max_non_circular, max_circular)

	if max_non_circular >= 0 {
		if max_non_circular > max_circular {
			return max_non_circular
		}
		return max_circular
	} else {
		return max_non_circular
	}

}

func maxSubArray(nums []int, is_max bool) int {
	// base case
	if len(nums) == 0 {
		return 0
	}
	best_sum := 0
	new_sum := 0
	if is_max {
		best_sum = int(math.Inf(-1))
		new_sum = int(math.Inf(-1))
	} else {
		best_sum = 0
		new_sum = 0
	}

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	for _, num := range nums {
		// check current number and max
		if is_max {
			new_sum = num + max(0, new_sum)
			best_sum = max(best_sum, new_sum)
		} else {
			new_sum = num + min(0, new_sum)
			best_sum = min(best_sum, new_sum)
		}
	}
	fmt.Println(best_sum, is_max)

	return best_sum

}
