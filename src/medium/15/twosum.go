package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2, 2}))
	fmt.Println(threeSum([]int{-3, -4, -2, 0, 2, 2, -2, 1, -1, 2, 3, -1, -5, -4, -5, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))

}

func twoSum(nums []int, target int, res *[][]int) {
	var i, j = 0, len(nums) - 1
	for i < j {
		if nums[i]+nums[j] == target {
			*res = append(*res, []int{nums[i], nums[j]})
			i++
			j--
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			i++
		}
	}
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		twoSum(nums[i+1:], -nums[i], &res)
	}
	return res
}
