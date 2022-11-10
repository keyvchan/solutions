package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates("abcd", 2))
	fmt.Println(removeDuplicates("deeedbbcccbdaa", 3))
	fmt.Println(removeDuplicates("pbbcggttciiippooaais", 2))
}

func removeDuplicates(s string, k int) string {
	bytes := []byte(s)

	stack := list.New()

	for _, b := range bytes {
		if stack.Len() < k-1 {
			// we don't have enough elements to check for duplicates
			stack.PushBack(b)
			continue
		}
		matches := 0
		// we have enough elements to check for duplicates
		// create k elements slice
		// check the last K - 1 elements
		last := stack.Back()
		for i := 0; i < k-1; i++ {
			if last.Value != b {
				// we have a mismatch
				stack.PushBack(b)
				break
			} else {
				// we have a match
				matches++
				last = last.Prev()
			}
		}
		if matches == k-1 {
			// we have a match
			// remove the last k-1 elements
			for i := 0; i < k-1; i++ {
				stack.Remove(stack.Back())
			}
		}

	}
	// convert the stack to a string
	var result string
	for e := stack.Front(); e != nil; e = e.Next() {
		result += string(e.Value.(byte))
	}
	return result
}
