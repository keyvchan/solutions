package main

import (
	"fmt"
)

func main() {
	fmt.Println(lastRemaining(10, 17))
}

func lastRemaining(n int, m int) int {
	// 0 1 2 3 4
	// 0 1 3 4
	// 1 3 4
	// 1 3
	// 3

}
