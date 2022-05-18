package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
}

func containsDuplicate(nums []int) bool {

	set := map[int]bool{}

	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = true
	}

	return false
}
