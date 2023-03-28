package main

import (
	"fmt"
)

func main() {
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
	fmt.Println(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := make([][]int, 0, k)
	i, j := 0, 0
	// generate the pairs
	for {
		fmt.Println(i, j)
		if i == len(nums1) && j == len(nums2) {
			break
		}
		if len(result) == k {
			break
		}
		result = append(result, []int{nums1[i], nums2[j]})
		// we determine the procceeding order
		if i < len(nums1)-1 && j < len(nums2)-1 {
			if nums1[i+1]-nums1[i] <= nums2[j+1]-nums2[j] {
				i++
			} else {
				j++
			}
		}
	}

	return result
}
