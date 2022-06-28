package main

import "fmt"

func main() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(1))
	fmt.Println(isUgly(14))
}

func isUgly(n int) bool {
	// check if n can be divided by 2, 3, 5

	for {
		if n%2 == 0 {
			n = n / 2
			continue
		}
		if n%3 == 0 {
			n = n / 3
			continue
		}
		if n%5 == 0 {
			n = n / 5
			continue
		}
		break
	}
	return n == 1

}
