package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	counts_map := make(map[int]int)
	for _, v := range nums {
		counts_map[v] += 1
	}
	counts := counts_map[0]
	i := 0
	for ; i < counts; i++ {
		nums[i] = 0
	}
	counts = counts_map[1] + i
	for ; i < counts; i++ {
		nums[i] = 1
	}
	counts = counts_map[2] + i
	for ; i < counts; i++ {
		nums[i] = 2
	}
}
