package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	fmt.Println(evalRPN([]string{""}))
}

func evalRPN(tokens []string) int {
	stack := list.New()

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			// pop top two elements
			e1 := stack.Remove(stack.Back())
			e2 := stack.Remove(stack.Back())

			switch token {
			case "+":
				stack.PushBack(e2.(int) + e1.(int))
			case "-":
				stack.PushBack(e2.(int) - e1.(int))
			case "*":
				stack.PushBack(e2.(int) * e1.(int))
			case "/":
				stack.PushBack(e2.(int) / e1.(int))
			}
		} else {
			i, err := strconv.Atoi(token)
			if err != nil {
				return 0
			}
			stack.PushBack(i)
		}
	}

	return stack.Remove(stack.Back()).(int)
}
