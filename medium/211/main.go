package main

func main() {

}

type Node struct {
	children map[rune]*Node
	isWord   bool
}

type WordDictionary struct {

	// so it's a trie
	// each node has a map of children
	// each node also has a boolean to indicate if it's a word

	node *Node
}

func Constructor() WordDictionary {

	// init the trie
	trie := WordDictionary{
		node: &Node{
			children: make(map[rune]*Node),
		},
	}

	// return the trie
	return trie

}

func (this *WordDictionary) AddWord(word string) {
	// add a word to the trie
	temp := this.node

	for _, char := range word {
		if _, ok := temp.children[char]; !ok {
			temp.children[char] = &Node{
				children: make(map[rune]*Node),
			}
		}

		temp = temp.children[char]
	}

	temp.isWord = true
}

func (node Node) Search(word string) bool {
	// search for a word in the trie

	// if the word is empty, return true if the node is a word
	if len(word) == 0 {
		return node.isWord
	}

	// if the word is a period, then we need to search all the children
	if word[0] == '.' {
		for _, child := range node.children {
			if child.Search(word[1:]) {
				return true
			}
		}
	} else {
		// if the word is not a period, then we need to search for the next character
		if _, ok := node.children[rune(word[0])]; ok {
			return node.children[rune(word[0])].Search(word[1:])
		}
	}

	return false
}

func (this *WordDictionary) Search(word string) bool {
	// search for a word in the trie

	// if the word is empty, return false
	if len(word) == 0 {
		return false
	}
	return this.node.Search(word)

}
