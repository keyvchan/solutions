package main

import (
	"fmt"
)

func main() {
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(2, -2))
	fmt.Println(myPow(1.00000, -2147483648))
}

func subPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}

	// check n
	if n%2 == 0 {
		result := subPow(x, n/2)
		return result * result
	} else {
		result := subPow(x, n/2)
		return x * result * result
	}

}

func myPow(x float64, n int) float64 {
	if x == 1 {
		return 1
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	result := subPow(x, n)

	if neg {
		return 1 / result
	}
	return result

}
