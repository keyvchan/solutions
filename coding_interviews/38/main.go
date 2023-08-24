package main

func main() {

}

func backtracking(s string, current string, result_map map[string]bool) {
	if len(s) == 0 {
		result_map[current] = true
		return
	}
	for i := 0; i < len(s); i++ {
		backtracking(s[:i]+s[i+1:], current+string(s[i]), result_map)
	}
}

func permutation(s string) []string {
	result_map := map[string]bool{}
	backtracking(s, "", result_map)

	result := []string{}
	for k := range result_map {
		result = append(result, k)
	}

	return result
}
