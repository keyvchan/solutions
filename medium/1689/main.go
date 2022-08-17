package main

import (
	"fmt"
)

func main() {
	fmt.Println(minPartitions("32"))
	fmt.Println(minPartitions("82734"))
	fmt.Println(minPartitions("27346209830709182346"))
}

func minPartitions(n string) int {
	// convert n to rune array

	max := '0'
	// linear scan to find the max number of partitions
	for _, r := range n {
		if r > max {
			max = r
		}
		if max == '9' {
			break
		}
	}
	return int(max - '0')
}
