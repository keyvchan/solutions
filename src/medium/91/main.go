package main

import "fmt"

func main() {
	fmt.Println(numDecodings("1201234"))
	fmt.Println(numDecodings("111111111111111111111111111111111111111111111"))
}

func backtracking(s string, current string, counts *int) {
	if len(s) == 0 {
		fmt.Println(current)
		*counts++
		return
	}
	if s[0] == '0' {
		// s[0] == 0, it's invalid
		return
	}
	// choose 1 or choose two
	if len(s) == 1 {
		// we can only choose 1, 1 always valid
		backtracking(s[1:], current+" "+string(s[0]), counts)
		return
	}
	// len(s) > 1
	// we can choose 1 or choose 2
	// take two digits
	if s[0] <= '2' && s[1] == '0' {
		// s[1] == 0, we can only choose 2
		backtracking(s[2:], current+" "+string(s[:2]), counts)
		return
	}
	if s[0] == '2' && s[1] > '6' {
		// s[0] >= 2, s[1] > 6, it's invalid
		// we can only choose 1
		backtracking(s[1:], current+" "+string(s[0]), counts)
		return
	}
	if s[0] > '2' {
		// len(s) >= 2 and s[0] >= 2, we can only choose 1
		backtracking(s[1:], current+" "+string(s[0]), counts)
		return

	}

	// other cases, we can choose 1 or choose 2
	backtracking(s[1:], current+" "+string(s[0]), counts)
	backtracking(s[2:], current+" "+string(s[:2]), counts)

}

func numDecodings(s string) int {
	counts := 0
	backtracking(s, "", &counts)

	return counts
}
