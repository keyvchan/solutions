package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}

func partitionLabels(s string) []int {
	count := 0
	return_array := []int{}

	had_array := []rune{}
	for i, c := range s {
		had_array = append(had_array, c)
		count += 1
		for j, cc := range had_array {
			if strings.Contains(s[i+1:], string(cc)) {
				break
			}
			if j == len(had_array)-1 {
				return_array = append(return_array, count)
				count = 0
				had_array = []rune{}
			}
		}
	}

	return return_array

}
