package main

import (
	"fmt"
	"io"
)

func main() {
	var n, m int
	var S string
	for {
		_, err := fmt.Scan(&n, &m)
		if err == io.EOF {
			break
		}
		// read s
		fmt.Scan(&S)

		x := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Scan(&x[i])
		}

		s := make([]string, m)

		// read m line
		for i := 0; i < m; i++ {
			fmt.Scan(&s[i])
		}

		fmt.Println(Solution(S, s))

	}

}

func dfs(S string, s []string, index int, m map[string]int) int {
	if index > len(S) {
		return 0
	}

	// we check the string before current index can be used or not
	if _, ok := m[S[:index+1]]; ok {
		// its a valid split point
		m[S[:index+1]]--
		if m[S[:index+1]] == 0 {
			delete(m, S[:index+1])
		}
		S = S[index+1:]
	} else {
		return 0
	}
	if len(m) == 0 {
		// we can say that we have done all the small string
		return 1
	}

	result := 0

	// for every position, we can choose to split here or not
	for i := 0; i < len(S); i++ {
		// we can split here
		// copy the map
		m2 := make(map[string]int)
		for k, v := range m {
			m2[k] = v
		}
		result += dfs(S, s, i, m2)
	}

	return result

}

func Solution(S string, s []string) int {
	// create a hash map to mark the small string is done or not
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	result := 0
	for i := 0; i < len(S); i++ {
		// copy the map
		m2 := make(map[string]int)
		for k, v := range m {
			m2[k] = v
		}
		result += dfs(S, s, i, m2)
	}

	return result
}
