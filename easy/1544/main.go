package main

func main() {

}

func makeGood(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1]+32 || s[i] == s[i+1]-32 {
			s = s[:i] + s[i+2:]
			i = -1
		}
	}
	return s
}
