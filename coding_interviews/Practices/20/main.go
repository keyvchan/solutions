package main

import "strings"

func main() {

}

func isNumber(s string) bool {

	new_s := strings.TrimSpace(s)

	if len(new_s) == 0 {
		return false
	}

	if new_s[0] == '-' || new_s[0] == '+' {
		// remove the sign
		new_s = new_s[1:]
	}
	if len(new_s) == 0 {
		return false
	}

	// the next should be a number or .
	if new_s[0] == '.' {
		// it must be a float
		// the next part should be a number
		new_s = new_s[1:]
		if len(new_s) == 0 {
			return false
		}
		for i, c := range new_s {
			if c < '0' || c > '9' {
				if c == 'e' || c == 'E' {
					new_s = new_s[i:]
					break
				}
				// we already have a dot, it must not have another dot
				return false
			}
		}
	} else if new_s[0] >= '0' && new_s[0] <= '9' {
		dot := false
		// it must be a number
		for i, c := range new_s {
			if c < '0' || c > '9' {
				if c == '.' {
					if dot {
						// we already have a dot, it must not have another dot
						return false
					}
					// we check if its a float
					dot = true
				}
				if c == 'e' || c == 'E' {
					new_s = new_s[i:]
					break
				}
				return false
			}
		}
	} else {
		return false
	}

	if len(new_s) == 0 {
		return true
	}
	if new_s[0] == 'e' || new_s[0] == 'E' {
		// the next part should be a number
		for _, c := range new_s {
			if c < '0' || c > '9' {
				return false
			}
		}
	}

	return true

}
