package main

import "fmt"

func main() {
	array := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(array.SumRange(0, 2))
}

type NumArray struct {
	array   []int
	sum_map map[int]int
}

func Constructor(nums []int) NumArray {
	sum_map := make(map[int]int)
	for i, v := range nums {
		sum_map[i] = sum_map[i-1] + v
	}
	fmt.Println(sum_map)
	return NumArray{nums, sum_map}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sum_map[right] - this.sum_map[left-1]
}
