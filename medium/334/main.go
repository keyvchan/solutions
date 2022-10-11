package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	fmt.Println(increasingTriplet([]int{0, 4, 2, 1, 0, -1, -3}))
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}

func increasingTriplet(nums []int) bool {

	i, j := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num <= i {
			i = num
		} else if num <= j {
			// num > i, check num < j
			j = num
		} else {
			// num > i, num > j
			return true
		}
	}
	return false
}
