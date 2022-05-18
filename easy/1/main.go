package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j, num2 := range nums[i+1:] {
			if num+num2 == target {
				return []int{i, i + j + 1}
			}
		}
	}
	return []int{0, 0}
}
