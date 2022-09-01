package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getHappyString(1, 3))
	fmt.Println(getHappyString(1, 4))
	fmt.Println(getHappyString(3, 9))
}

func backtracking(n int, base []string, last string, current string, result *[]string) {
	if len(current) == n {
		*result = append(*result, current)
		return
	}

	for _, b := range base {
		if b == last {
			continue
		}
		backtracking(n, base, b, current+b, result)
	}

}

func getHappyString(n int, k int) string {
	// base set
	base := []string{"a", "b", "c"}
	result := []string{}
	backtracking(n, base, "", "", &result)
	sort.Strings(result)

	if k > len(result) {
		return ""
	}

	return result[k-1]
}
