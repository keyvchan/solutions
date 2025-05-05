package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
}

func findComplement(num int) int {
	// conver num to string
	stra := strconv.FormatInt(int64(num), 2)
	straCopy := make([]byte, len(stra))
	copy(straCopy, stra)

	// convert all 1 to 0, 0 to 1
	for i := range len(stra) {
		if stra[i] == '1' {
			straCopy[i] = '0'
		} else {
			straCopy[i] = '1'
		}
	}
	// convert string to int
	res, _ := strconv.ParseInt(string(straCopy), 2, 32)
	return int(res)
}
