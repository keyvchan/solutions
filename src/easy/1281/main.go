package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {

	// convert int to string

	str := strconv.Itoa(n)

	sum := 0
	product := 1

	for _, v := range str {
		sum += int(v - '0')
		product *= int(v - '0')
	}

	return product - sum

}
