package main

import "fmt"

func main() {
	fmt.Println(checkPossibility([]int{4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 1}))
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
}

func checkPossibility(nums []int) bool {
	flag := false
	for i, v := range nums {
		if i == 0 {
			continue
		}
		last := 0
		if i != 0 {
			last = nums[i-1]
		}
		if v < last {
			if flag {
				return false
			}
			flag = true
			nums[i] = nums[i-1]
		}
	}

	return true
}
