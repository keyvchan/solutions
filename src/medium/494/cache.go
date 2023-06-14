package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func dfs(nums []int, target int, index int, sum int, cache map[pair]int) int {
	if index == len(nums) {
		if sum == target {
			return 1
		}
		return 0
	}
	if _, ok := cache[pair{index, sum}]; ok {
		return cache[pair{index, sum}]
	}

	result := dfs(nums, target, index+1, sum+nums[index], cache) + dfs(nums, target, index+1, sum-nums[index], cache)

	cache[pair{index, sum}] = result
	return result
}

type pair struct {
	index int
	sum   int
}

func findTargetSumWays(nums []int, target int) int {
	cache := map[pair]int{}
	return dfs(nums, target, 0, 0, cache)
}
