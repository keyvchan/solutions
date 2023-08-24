package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maximumValue([]string{"alic3"}))
}

func maximumValue(strs []string) int {

	// use atoi
	var largest int
	for _, str := range strs {
		if i, err := strconv.Atoi(str); err == nil {
			if i > largest {
				largest = i
			}
		} else {
			if len(str) > largest {
				// err != nil
				largest = len(str)
			}
		}
	}

	return largest
}
