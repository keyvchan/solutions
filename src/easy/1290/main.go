package main

import "math"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {

	array := []int{}

	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}

	sum := 0
	counts := 0
	for i := len(array) - 1; i >= 0; i-- {
		sum += int(math.Pow(2, float64(counts)) * float64(array[i]))
		counts++
	}

	return sum
}
