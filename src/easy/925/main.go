package main

import (
	"fmt"
)

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))
	fmt.Println(isLongPressedName("saeed", "ssaaedd"))
	fmt.Println(isLongPressedName("leelee", "lleeelee"))
	fmt.Println(isLongPressedName("vtkgn", "vttkgnn"))
}

type Section struct {
	Char  rune
	Count int
}

func isLongPressedName(name string, typed string) bool {

	// when the name == typed, its equal
	if name == typed {
		return true
	}

	if name[len(name)-1] != typed[len(typed)-1] {
		// if the last char is not equal, it must not be long pressed
		return false
	}

	// count the sections of the name
	nameSections := make([]Section, 0)
	lastChar := rune(0)

	currentSectionName := Section{}

	for _, char := range name {
		if lastChar == 0 {
			lastChar = char
			currentSectionName.Char = char
			currentSectionName.Count++
		} else if lastChar == char {
			currentSectionName.Count++
		} else {
			nameSections = append(nameSections, currentSectionName)
			currentSectionName = Section{}
			currentSectionName.Char = char
			currentSectionName.Count++
			lastChar = char
		}

	}

	nameSections = append(nameSections, currentSectionName)

	// count the sections of the typed
	typedSections := make([]Section, 0)
	lastChar = rune(0)
	currentSectionTyped := Section{}

	for _, char := range typed {
		if lastChar == 0 {
			lastChar = char
			currentSectionTyped.Char = char
			currentSectionTyped.Count++
		} else if lastChar == char {
			currentSectionTyped.Count++
		} else {
			typedSections = append(typedSections, currentSectionTyped)
			currentSectionTyped = Section{}
			currentSectionTyped.Char = char
			currentSectionTyped.Count++
			lastChar = char
		}

	}

	typedSections = append(typedSections, currentSectionTyped)

	// check if the sections are equal
	if len(nameSections) != len(typedSections) {
		return false
	}

	for i := 0; i < len(nameSections); i++ {
		if nameSections[i].Char != typedSections[i].Char || nameSections[i].Count > typedSections[i].Count {
			return false
		}
	}

	return true

}
