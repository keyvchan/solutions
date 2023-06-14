package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(firstBadVersion(5))
	fmt.Println(firstBadVersion(1))
}

// test purpose
// Forward declaration of isBadVersion API.
// @param   version   your guess about first bad version
// @return 	 	      true if current version is bad
//  		          false if current version is good
// func isBadVersion(version int) bool;
func isBadVersion(version int) bool {
	return version == 4
}

func binarySearch(left int, right int) int {
	fmt.Println(left, right)
	if left == right || math.Abs(float64(left-right)) == 1 {
		if isBadVersion(left) {
			return left
		} else {
			return right
		}
	}
	result := 0
	if isBadVersion((left + right) / 2) {
		// search left
		result = binarySearch(left, (left+right)/2)
	} else {
		// search right
		result = binarySearch((left+right)/2, right)
	}
	return result
}

func firstBadVersion(n int) int {
	return binarySearch(0, n)
}
