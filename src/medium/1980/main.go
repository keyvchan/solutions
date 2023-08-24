package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(findDifferentBinaryString([]string{"00", "01"}))
	fmt.Println(findDifferentBinaryString([]string{"0000000111", "0000001001", "0000000100", "0000000001", "0000000010", "1111111111", "0000000101", "0000000000", "0000001000", "0000000110"}))
	fmt.Println(findDifferentBinaryString([]string{"1"}))
}

func findDifferentBinaryString(nums []string) string {
	// first we use a hashmap to store all nums
	num_len := len(nums[0])

	nums_map := map[int64]bool{}

	for _, num := range nums {
		n, _ := strconv.ParseInt(num, 2, num_len+1)
		nums_map[n] = true
	}
	fmt.Println(nums_map)

	for i := 0; i < int(math.Pow(2, float64(num_len))); i++ {
		if _, ok := nums_map[int64(i)]; !ok {
			return fmt.Sprintf("%0*b", num_len, i)
		}
	}

	return ""

}
