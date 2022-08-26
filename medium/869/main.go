package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reorderedPowerOf2(1))
	fmt.Println(reorderedPowerOf2(10))
	fmt.Println(reorderedPowerOf2(61))
	fmt.Println(reorderedPowerOf2(4609))
}

func backtracking(digits []byte, max int, current string) bool {
	if len(digits) == 0 {
		s, _ := strconv.Atoi(current)
		if max%s == 0 {
			return true
		} else {
			return false
		}
	}

	result := false

	for i, digit := range digits {
		if current == "" && digit == '0' {
			continue
		}
		new_digits := make([]byte, len(digits)-1)
		copy(new_digits, digits[:i])
		copy(new_digits[i:], digits[i+1:])

		result = result || backtracking(new_digits, max, current+string(digit))
	}

	return result
}

func reorderedPowerOf2(n int) bool {

	// the maxium of power of 2
	max := 1 << 30

	if n > 0 && n < 10 {
		return max%n == 0
	}

	n_str := strconv.FormatInt(int64(n), 10)
	digits := []byte(n_str)

	// try out every combination of digits

	return backtracking(digits, max, "")
}
