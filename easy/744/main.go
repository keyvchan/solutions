package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd')))
	fmt.Println(string(nextGreatestLetter([]byte{'a', 'b'}, 'z')))
}

func nextGreatestLetter(letters []byte, target byte) byte {

	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	for _, v := range letters {
		if v > target {
			return v
		}
	}

	return letters[0]
}
