package main

import "fmt"

func main() {
	fmt.Println(fib(4))
}

func fib(n int) int {
	f_map := map[int]int{0: 0, 1: 1}

	for i := 2; i <= n; i++ {
		f_map[i] = f_map[i-1] + f_map[i-2]
	}
	return f_map[n]

}
