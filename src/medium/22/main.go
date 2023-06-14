package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func backtracking(open_counts int, close_counts int, target int, current string, result *[]string) {
	if close_counts == target {
		*result = append(*result, current)
		return
	}

	if open_counts == close_counts {
		// since open_counts == close_counts, we can only add one more open
		backtracking(open_counts+1, close_counts, target, current+"(", result)
	} else if open_counts > close_counts {
		// since open_counts > close_counts, we can add close and open, but we need to check open_counts and target
		if open_counts == target {
			// open_counts == target, we can only add one more close
			backtracking(open_counts, close_counts+1, target, current+")", result)
		} else {
			// we can add close and open
			backtracking(open_counts, close_counts+1, target, current+")", result)
			backtracking(open_counts+1, close_counts, target, current+"(", result)
		}
	}
}

func generateParenthesis(n int) []string {

	result := []string{}
	backtracking(0, 0, n, "", &result)

	return result

}
