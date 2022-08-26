package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution("bdaaadadb"))
	fmt.Println(Solution("abacb"))
	fmt.Println(Solution("zthtzh"))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(s string, current string, index int, wanted map[string]int, memo map[Pair]int) int {
	if len(s) == 0 {
		if len(wanted) == 0 {
			return len(current)
		}
		return 0
	}

	wanted[string(s[0])]++
	if wanted[string(s[0])]%2 == 0 {
		delete(wanted, string(s[0]))
	}
	current += string(s[0])
	if val, ok := memo[Pair{current, index}]; ok {
		return val
	}

	l := -9223372036854775807
	if len(wanted) == 0 {
		l = len(current)
	}
	// in a position, we can include this char or start from current location
	l = max(l, dfs(s[1:], current, index+1, wanted, memo))

	// create a empty map
	new_wanted := map[string]int{}
	l = max(l, dfs(s[1:], "", index+1, new_wanted, memo))

	memo[Pair{current, index}] = l

	return l
}

type Pair struct {
	current string
	index   int
}

func Solution(S string) int {
	wanted := map[string]int{}
	memo := map[Pair]int{}

	result := dfs(S, "", 0, wanted, memo)

	return result

}
