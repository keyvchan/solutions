package main

type Trie struct {
	chars       map[rune]*Trie
	end_of_word bool
}

func Constructor() Trie {
	return Trie{
		chars:       map[rune]*Trie{},
		end_of_word: false,
	}
}

func (this *Trie) Insert(word string) {
	temp := this
	for i, c := range word {
		// check if c in temp.chars, if not, create a new Trie
		if _, ok := temp.chars[c]; !ok {
			temp.chars[c] = &Trie{
				chars:       map[rune]*Trie{},
				end_of_word: false,
			}
		}
		if i == len(word)-1 {
			temp.chars[c].end_of_word = true
		}
		temp = temp.chars[c]

	}
}

func (this *Trie) Search(word string) bool {
	temp := this
	for i, c := range word {
		if _, ok := temp.chars[c]; !ok {
			return false
		}
		if i == len(word)-1 {
			return temp.chars[c].end_of_word
		}
		temp = temp.chars[c]
	}
	return true

}

func (this *Trie) StartsWith(prefix string) bool {
	temp := this
	for _, c := range prefix {
		if _, ok := temp.chars[c]; !ok {
			return false
		}
		temp = temp.chars[c]
	}
	return true
}
