package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 1))
}

func backtracking(current string, n int, k int, result *[]string) {
	if n == 0 {
		*result = append(*result, current)
		return
	}

	// check last digit
	lastDigit, _ := strconv.Atoi(string(current[len(current)-1]))
	// if lastDigit - k < 0 && lastDigit + k > 9, impossible to make a number
	if lastDigit-k >= 0 {
		backtracking(current+strconv.Itoa(lastDigit-k), n-1, k, result)
	}
	// if lastDigit-k == lastDigit+k, it will be duplicated
	if lastDigit-k != lastDigit+k && lastDigit+k <= 9 {
		backtracking(current+strconv.Itoa(lastDigit+k), n-1, k, result)
	}
}

func numsSameConsecDiff(n int, k int) []int {
	// use string to store the number

	result := []string{}
	// the first digit is not zero
	for i := 1; i <= 9; i++ {
		backtracking(strconv.Itoa(i), n-1, k, &result)
	}

	// convert string to int
	intResult := []int{}
	for _, v := range result {
		temp_v, _ := strconv.Atoi(v)
		intResult = append(intResult, temp_v)
	}
	return intResult
}
