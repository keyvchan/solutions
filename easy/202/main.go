package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	record := map[int]bool{}
	for {
		str := strconv.Itoa(n)
		sum := 0
		for _, char := range str {
			num, _ := strconv.Atoi(string(char))
			sum += num * num
		}
		if sum == 1 {
			return true
		} else {
			if record[sum] {
				return false
			}
			record[sum] = true
			n = sum
		}
	}
}
