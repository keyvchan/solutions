package main

import (
	"fmt"
)

func main() {
	fmt.Println(shiftingLetters("abc", []int{3, 5, 9}))
	fmt.Println(shiftingLetters("aaa", []int{1, 2, 3}))
	fmt.Println(shiftingLetters("a", []int{1}))
}

func shiftLetter(letter byte, shift int) byte {
	// convert the letter to a number
	letterNum := int(letter) - 97
	// apply the shift
	letterNum = (letterNum + shift) % 26
	// convert the number back to a letter
	letter = byte(letterNum + 97)
	return letter
}

func shiftingLetters(s string, shifts []int) string {
	// calcuate the sum of all shifts
	shiftsFinal := make([]int, len(shifts))

	// backwards loop through the shifts
	for i := len(shifts) - 1; i >= 0; i-- {
		// add the shift to the final shifts
		if i == len(shifts)-1 {
			shiftsFinal[i] = shifts[i]
		} else {
			shiftsFinal[i] = shifts[i] + shiftsFinal[i+1]
		}
	}
	for i, shift := range shiftsFinal {
		shiftsFinal[i] = shift % 26
	}
	// apply the shifts to the string
	new_string := make([]byte, len(s))
	for i, shift := range shiftsFinal {
		new_string[i] = shiftLetter(s[i], shift)
	}

	return string(new_string)
}
