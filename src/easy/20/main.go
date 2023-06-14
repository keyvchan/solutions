package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := list.New()
	for _, v := range s {
		if v == ')' || v == ']' || v == '}' {
			// pop
			if stack.Len() == 0 {
				return false
			}
			e := stack.Remove(stack.Back())
			if v == ')' && e.(rune) != '(' {
				return false
			}
			if v == ']' && e.(rune) != '[' {
				return false
			}
			if v == '}' && e.(rune) != '{' {
				return false
			}
		} else {
			// push
			stack.PushBack(v)
		}
	}
	return stack.Len() == 0
}
