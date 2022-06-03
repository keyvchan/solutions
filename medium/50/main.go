package main

import (
	"fmt"
)

func main() {
	fmt.Println(myPow(2, 10))
	fmt.Println(myPow(2, -2))
	fmt.Println(1.00000,
		-2147483648)
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	y := 1.0
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	for i := 0; i < n; i++ {
		y = y * x
	}
	if !neg {
		return y
	}

	return 1 / y

}
