package main

import (
	"math/big"
	"strconv"
)

func main() {

}

func plusOne(digits []int) []int {

	// convert digits to int
	var num_str string
	for _, digit := range digits {
		num_str += strconv.Itoa(digit)
	}

	num, is_success := big.NewInt(0).SetString(num_str, 10)
	if !is_success {
		return []int{}
	}

	add_result := big.NewInt(0).Add(num, big.NewInt(1))

	// convert back to digits
	var result []int
	for _, char := range add_result.String() {
		digit, _ := strconv.Atoi(string(char))
		result = append(result, digit)
	}
	return result

}
