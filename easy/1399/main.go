package main

import "fmt"

func main() {
	fmt.Println(countLargestGroup(13))
}

func countLargestGroup(n int) int {
	number := make([]int, n+1)
	ans := make([]int, n+1)

	largest := 0
	for i := 1; i < n+1; i++ {
		ans[i] = ans[i/10] + i%10
		if number[ans[i]]+1 > largest {
			largest = number[ans[i]] + 1
		}
		number[ans[i]] = number[ans[i]] + 1
	}

	counts := 0
	for _, v := range number {
		if v == largest {
			counts++
		}
	}
	return counts
}
