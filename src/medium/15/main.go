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

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	for i, num := range nums {
		if num > 0 {
			break
		}
		if i > 0 && num == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			sum := num + nums[j] + nums[k]
			if sum > 0 {
				k -= 1
			} else if sum < 0 {
				j += 1
			} else {
				result = append(result, []int{num, nums[j], nums[k]})
				for k > 0 {
					if nums[k] == nums[k-1] {
						k -= 1
					} else {
						break
					}
				}
				k -= 1
				j += 1
			}
		}
	}

	return result
}
