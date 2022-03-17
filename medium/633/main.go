package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(3))
	fmt.Println(judgeSquareSum(2))
	fmt.Println(judgeSquareSum(0))
	fmt.Println(judgeSquareSum(2147483645))
}

func judgeSquareSum(c int) bool {

	mid := int(math.Sqrt(float64(c)))

	i, j := 0, mid
	for {
		if i > j {
			break
		}
		if i*i+j*j == c {
			return true
		}
		if i*i+j*j > c {
			j--
		} else {
			i++
		}
	}

	return false
}
