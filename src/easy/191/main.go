package main

import "strconv"

func main() {

}

func hammingWeight(num uint32) int {
	// convert it to binary string
	st := strconv.FormatUint(uint64(num), 2)

	count := 0
	for _, v := range st {
		if v == '1' {
			count++
		}
	}
	return count

}
