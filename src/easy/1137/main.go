package main

import "fmt"

func main() {
	fmt.Println(tribonacci(25))
}

func tribonacci(n int) int {
	f_map := map[int]int{0: 0, 1: 1, 2: 1}

	for i := 3; i <= n; i++ {
		f_map[i] = f_map[i-1] + f_map[i-2] + f_map[i-3]
	}

	return f_map[n]

}
