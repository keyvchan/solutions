package main

import "fmt"

func main() {
	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
	// fmt.Println(zeroFilledSubarray([]int{1, 0, 0, 1, 0, 1, 1, 0, 1, 0}))
}

func zeroFilledSubarray(nums []int) int64 {

	consecutiveZeros := int64(0)

	sum := int64(0)

	for _, num := range nums {
		// check current num is zero
		if num == 0 {
			// we can add to the sum
			sum += (consecutiveZeros + 1)
			consecutiveZeros++
		} else {
			// reset the consecutiveZeros
			consecutiveZeros = 0
		}
	}

	return int64(sum)
}
