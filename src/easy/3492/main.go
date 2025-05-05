package main

import "fmt"

func main() {
	fmt.Println(maxContainers(2, 3, 15))
	fmt.Println(maxContainers(3, 5, 20))
}

func maxContainers(n int, w int, maxWeight int) int {
	rem := maxWeight / w

	if rem >= n*n {
		return n * n
	} else {
		return rem
	}
}
