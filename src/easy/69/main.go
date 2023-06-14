package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(2147395599))
	fmt.Println(mySqrt(1))
}

func mySqrt(x int) int {

	if x == 1 {
		return 1
	}
	if x == 0 {
		return 0
	}

	last := x
	current := x / 2
	for {
		fmt.Println(current, last)

		if current*current > x {
			last = current
			current = current / 2
		} else {
			current = (current + last) / 2
			if current*current > x {
				continue
			}
			if current-last == 1 || current-last == -1 || current-last == 0 {
				break
			}
		}

	}

	return current
}
