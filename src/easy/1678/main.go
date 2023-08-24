package main

import (
	"strings"
)

func main() {
}

func interpret(command string) string {

	output := strings.Builder{}
	output.Grow(len(command))

	for i := 0; i < len(command); {
		if i >= len(command) {
			break
		}

		if command[i] == 'G' {
			// it's a G
			output.WriteByte('G')
			i++
		} else if command[i] == '(' && command[i+1] == ')' {
			// it's a ()
			output.WriteByte('o')
			i += 2
		} else if command[i] == '(' && command[i+1] == 'a' && command[i+2] == 'l' && command[i+3] == ')' {
			// it's a (al)
			output.WriteString("al")
			i += 4
		}
	}

	return output.String()

}
