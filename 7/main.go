package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}

// 123
// 1*10^2 + 2*10^1 + 3*10^0
// 3*10^2 + 2*10^1 + 1*10^0
// 1111011
// 321
// 101000001

func reverse(x int) int {
	result := strconv.Itoa(x)
	fmt.Println(result)
	sum := 0
	false_flag := false
	if result[0] == '-' {
		false_flag = true
		result = result[1:]
	}
	for i := len(result) - 1; i >= 0; i-- {
		int_result, _ := strconv.Atoi(string(result[i]))
		if int_result*int(math.Pow10(i)) > math.MaxInt32-sum {
			return 0
		}
		sum = sum + int_result*int(math.Pow10(i))
	}

	if false_flag {
		sum = -sum
	}
	return sum
}
