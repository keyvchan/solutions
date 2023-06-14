package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(largestOddNumber("1357"))
	fmt.Println(largestOddNumber("13579"))
	fmt.Println(largestOddNumber("52"))
	fmt.Println(largestOddNumber("4206"))
	fmt.Println(largestOddNumber("35427"))
	fmt.Println(largestOddNumber("0"))
}

func largestOddNumber(num string) string {
	// odd number should be end with 1 3 5 7 9

	right := len(num) - 1

	for right >= 0 {
		number, _ := strconv.Atoi(string(num[right]))
		if number%2 == 1 {
			// find the end of odd number, extend the sliding window
			break
		}
		right--
	}

	return num[:right+1]
}
