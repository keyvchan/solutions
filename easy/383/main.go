package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {

	// count word frequency
	ransomNoteFreq := make(map[rune]int)
	for _, c := range ransomNote {
		ransomNoteFreq[c]++
	}
	magazineFreq := make(map[rune]int)
	for _, c := range magazine {
		magazineFreq[c]++
	}

	// for every character in ransomNote, the frequency shoule always be less than or equal to the frequency in magazine
	for c, freq := range ransomNoteFreq {
		if magazineFreq[c] < freq {
			return false
		}
	}
  return true

}
