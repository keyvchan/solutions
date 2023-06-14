package main

import (
	"fmt"
)

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 1))
}

func backtracking(current int, last int, n int, k int, result *[]int) {
	if n == 0 {
		*result = append(*result, current)
		return
	}

	// if lastDigit - k < 0 && lastDigit + k > 9, impossible to make a number
	if last-k >= 0 {

		temp := current*10 + last - k
		backtracking(temp, last-k, n-1, k, result)
	}
	// if lastDigit-k == lastDigit+k, it will be duplicated
	if last-k != last+k && last+k <= 9 {
		temp := current*10 + last + k
		backtracking(temp, last+k, n-1, k, result)
	}
}

func numsSameConsecDiff(n int, k int) []int {
	// use string to store the number

	result := []int{}
	// the first digit is not zero
	for i := 1; i <= 9; i++ {
		backtracking(i, i, n-1, k, &result)
	}

	return result
}
