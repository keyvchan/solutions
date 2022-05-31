package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {

	sum := 0
	max := nums[0]
	for _, v := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += v
		if sum > max {
			max = sum
		}
	}
	return max
}
