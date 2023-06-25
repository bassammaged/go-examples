package trie

import "strings"

const AlphabetSize = 26

type Node struct {
	children [26]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{&Node{}}
}

// Insert() Add the word into the trie struct
//
// No returns
func (t *Trie) Insert(word string) {

	// Point to the currentNode, start with the root
	currentNode := t.root
	// Lowercase the letters
	word = strings.ToLower(word)

	// Iterate over the word
	for _, letter := range word {
		// Change from unicode char (Rune) to english index from 0-25
		charIndex := letter - 'a'

		// Check if the letter index has value, if it's nil add empty slice *Node
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		// Move the pointer to the next children node
		currentNode = currentNode.children[charIndex]
	}
	// Flag the last letter with isEnd true
	currentNode.isEnd = true
}

// Search() Lookup for the word in trie struct
//
// Returns true if trie includes the word. Otherwise, return false
func (t Trie) Search(word string) (result bool) {
	// Point to the currentNode, start with the root
	currentNode := t.root

	// Iterate over the word
	for _, letter := range word {
		// Change from unicode char (Rune) to english index from 0-25
		charIndex := letter - 'a'
		// Check if the letter index has value, if it's nil add empty slice *Node
		if currentNode.children[charIndex] == nil {
			return false
		}
		// Move the pointer to the next children node
		currentNode = currentNode.children[charIndex]
	}
	// Return true if isEnd is marked as true
	return currentNode.isEnd
}

// GetDict() Lookup for all possible word
//
// Returns: if trie includes the word return true, and all possible word in []string
func (t Trie) GetDict(word string) (result bool, dict []string) {
	var words []string

	// Point to the currentNode, start with the root
	currentNode := t.root
	// Lowercase the letters
	word = strings.ToLower(word)

	// Iterate over the word
	for index, letter := range word {
		// Change from unicode char (Rune) to english index from 0-25
		charIndex := letter - 'a'

		// Add the possible words into words slice
		if currentNode.isEnd {
			words = append(words, word[:index])
		}
		// Check if the letter index has value, if it's nil add empty slice *Node
		if currentNode.children[charIndex] == nil {
			return false, words
		}

		// Move the pointer to the next children node
		currentNode = currentNode.children[charIndex]
	}
	// Return true if isEnd is marked as true
	return currentNode.isEnd, words
}
