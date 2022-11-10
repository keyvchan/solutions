package main

import "fmt"

func main() {
	fmt.Println(executeInstructions(5, []int{1, 2}, "LLRLR"))
	fmt.Println(executeInstructions(3, []int{0, 1}, "RRDDLU"))
	fmt.Println(executeInstructions(2, []int{1, 1}, "LURD"))
	fmt.Println(executeInstructions(1, []int{0, 0}, "LRUD"))
}

func executeInstruction(n int, currentPos []int, s string) int {
	// if the currentPos out of the grid, return -1
	if currentPos[0] < 0 || currentPos[0] >= n || currentPos[1] < 0 || currentPos[1] >= n {
		return -1
	}
	if s == "" {
		return 0
	}
	result := 1
	switch s[0] {
	case 'U':
		result += executeInstruction(n, []int{currentPos[0] - 1, currentPos[1]}, s[1:])
	case 'D':
		result += executeInstruction(n, []int{currentPos[0] + 1, currentPos[1]}, s[1:])
	case 'L':
		result += executeInstruction(n, []int{currentPos[0], currentPos[1] - 1}, s[1:])
	case 'R':
		result += executeInstruction(n, []int{currentPos[0], currentPos[1] + 1}, s[1:])
	}
	return result
}

func executeInstructions(n int, startPos []int, s string) []int {
	// start from every index
	result := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		result = append(result, executeInstruction(n, startPos, s[i:]))
	}

	return result
}
