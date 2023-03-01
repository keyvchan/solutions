package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortArray([]int{3, 2, 1}))
	fmt.Println(sortArray([]int{1, 2, 3}))
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}

func partition(nums []int, start, end int) int {
	// use the last element as pivot
	pivot := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(nums, start, end)
	quickSort(nums, start, pivot-1)
	quickSort(nums, pivot+1, end)
}

func sortArray(nums []int) []int {
	// use quick sort
	quickSort(nums, 0, len(nums)-1)

	return nums
}
