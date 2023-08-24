package main

import (
	"fmt"
)

func main() {
	fmt.Println(nthUglyNumber(10))
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
func nthUglyNumber(n int) int {

	// use a map to store the ugly number
	i2, i3, i5 := 0, 0, 0
	ugly := make([]int, n)
	ugly[0] = 1
	for i := 1; i < n; i++ {
		ugly[i] = min(ugly[i2]*2, ugly[i3]*3, ugly[i5]*5)
		if ugly[i] == ugly[i2]*2 {
			i2++
		}
		if ugly[i] == ugly[i3]*3 {
			i3++
		}
		if ugly[i] == ugly[i5]*5 {
			i5++
		}
	}

	return ugly[n-1]
}
