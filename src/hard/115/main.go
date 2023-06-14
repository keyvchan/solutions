package main

import "fmt"

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("babgbag", "bag"))
	fmt.Println(numDistinct("dbaaadcddccdddcadacbadbadbabbbcad", "dadcccbaab"))
	fmt.Println(numDistinct("aabdbaabeeadcbbdedacbbeecbabebaeeecaeabaedadcbdbcdaabebdadbbaeabdadeaabbabbecebbebcaddaacccebeaeedababedeacdeaaaeeaecbe", "bddabdcae"))
}

func dfs(s string, t string, index int, current string, memo map[Key]int) int {
	if index == len(s) {
		// check current
		if current == t {
			return 1
		}
		return 0
	}
	if len(current) == len(t) {
		// check current
		if current == t {
			return 1
		}
		return 0
	}
	if len(current) > len(t) {
		// current is longer than t, no need to continue
		return 0
	}
	if len(s)-index < len(t)-len(current) {
		// not enough characters left to match t
		return 0
	}

	// check memo
	if val, ok := memo[Key{index, current}]; ok {
		return val
	}
	// we could delete the current character or not

	result := 0
	// check current index
	if s[index] != t[len(current)] {
		// if current character does not match, we can only delete it
		result = dfs(s, t, index+1, current, memo)
	} else {
		result = dfs(s, t, index+1, current, memo) + dfs(s, t, index+1, current+string(s[index]), memo)
	}

	// delete
	memo[Key{index, current}] = result

	return result
}

type Key struct {
	index   int
	current string
}

func numDistinct(s string, t string) int {
	memo := map[Key]int{}
	return dfs(s, t, 0, "", memo)
}
