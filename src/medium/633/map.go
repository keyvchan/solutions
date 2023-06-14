package main

import (
	"fmt"
)

func main() {
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(3))
	fmt.Println(judgeSquareSum(2))
	fmt.Println(judgeSquareSum(0))
	fmt.Println(judgeSquareSum(2147483645))
}

func judgeSquareSum(c int) bool {
	number := map[int]bool{}

	for i := 0; i <= c; i++ {
		if i*i <= c {
			number[i*i] = true
			if number[c-i*i] {
				return true
			}
		} else {
			break
		}
	}

	return false
}
