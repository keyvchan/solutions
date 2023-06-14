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

type Record struct {
	Char  byte
	Count int
}

func removeDuplicates(s string, k int) string {
	bytes := []byte(s)

	stack := list.New()

	for _, b := range bytes {
		if stack.Len() == 0 {
			stack.PushBack(Record{Char: b, Count: 1})
			continue
		} else {
			// check last element
			last := stack.Back()
			if last.Value.(Record).Char == b {
				// check count
				if last.Value.(Record).Count == k-1 {
					stack.Remove(last)
				} else {
					last.Value = Record{Char: b, Count: last.Value.(Record).Count + 1}
				}
			} else {
				stack.PushBack(Record{Char: b, Count: 1})
			}
		}

	}
	// build result
	result := make([]byte, 0)
	for e := stack.Front(); e != nil; e = e.Next() {
		for i := 0; i < e.Value.(Record).Count; i++ {
			result = append(result, e.Value.(Record).Char)
		}
	}
	return string(result)
}
