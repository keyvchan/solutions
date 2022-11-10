package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(s string) string {
	// s to byte array
	bytes := []byte(s)
	// stack
	stack := list.New()
	for _, b := range bytes {
		// keep check current char and last char in stack
		if stack.Len() == 0 {
			stack.PushBack(b)
			continue
		}
		// if current char is equal to last char in stack, pop it
		if b == stack.Back().Value {
			stack.Remove(stack.Back())
		} else {
			stack.PushBack(b)
		}
	}

	// stack to string
	var res string
	for e := stack.Front(); e != nil; e = e.Next() {
		res += string(e.Value.(byte))
	}
	return res
}
