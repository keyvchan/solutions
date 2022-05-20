package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

func twoSum(numbers []int, target int) []int {

	// create a map
	m := map[int]int{}

	for i, number := range numbers {
		if _, ok := m[target-number]; ok {
			return []int{m[target-number] + 1, i + 1}
		} else {
			m[number] = i
		}
	}

	return []int{}
}
