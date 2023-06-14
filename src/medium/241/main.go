package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))

}
func backtracking(expression string, memo map[string][]int) []int {
	if v, ok := memo[expression]; ok {
		return v
	}

	result := []int{}
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			lefts := backtracking(expression[:i], memo)
			rights := backtracking(expression[i+1:], memo)
			result_temp := 0
			for _, left := range lefts {
				for _, right := range rights {
					switch expression[i] {
					case '+':
						result_temp = left + right
					case '-':
						result_temp = left - right
					default:
						result_temp = left * right
					}
					result = append(result, result_temp)
				}
			}
		}
	}
	if len(result) == 0 {
		num, _ := strconv.Atoi(expression)
		result = append(result, num)
	}
	memo[expression] = result
	return result
}

func diffWaysToCompute(expression string) []int {
	memo := map[string][]int{}

	return backtracking(expression, memo)
}
