package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}

func findMin(nums []int) int {

	min := 0
	max := len(nums) - 1
	mid := 0
	for {
		if min == max || max-min == 1 {
			// check max and min
			if nums[max] < nums[min] {
				return nums[max]
			} else {
				return nums[min]
			}
		}
		mid = (min + max) / 2

		if nums[mid] > nums[len(nums)-1] {
			min = mid
		} else {
			max = mid
		}

	}

}
