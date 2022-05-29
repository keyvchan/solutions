package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for j, rest := range temperatures[i+1:] {
			if rest > temp {
				res[i] = j + 1
				break
			}
			if j == len(temperatures[i+1:])-1 {
				res[i] = 0
			}
		}
	}
	return res
}
