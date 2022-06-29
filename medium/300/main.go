package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7}))
	fmt.Println(lengthOfLIS([]int{1, 2, 3}))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(nums []int, index int, max int) int {
	if index >= len(nums) {
		return 0
	}

	result := 0
	// we choose include the current number or not
	if nums[index] > max {
		result = dfs(nums, index+1, nums[index]) + 1
		result = maxInt(result, dfs(nums, index+1, max))
	} else {
		result = dfs(nums, index+1, max)
	}

	return result
}

func lengthOfLIS(nums []int) int {

	caching := map[int]int{}
	max := 0
	for i := range nums {
		var temp int
		if _, ok := caching[i]; ok {
			temp = caching[i]
		} else {
			temp = dfs(nums, i, -2147483648)
		}

		if temp > max {
			max = temp
		}
	}

	return max
}
