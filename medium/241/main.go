package main

import (
	"errors"
	"strconv"
)

func main() {

}

func evalute(expression string, last_result int) int {
	for i := 0; i < len(expression); i++ {
			// fork the result
			not_add := evalute(expression[:i], last_result)
			add_result := evalute(expression[i+1:], last_result)

		}
	}
	return 0
}

func diffWaysToCompute(expression string) []int {

	evalute(expression, 0)

	return nil
}
