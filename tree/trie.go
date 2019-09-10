package main

import (
	"fmt"
)

type TrieTreeNode struct {
	chld map[rune]*TrieTreeNode
	end  bool
}

type TrieTree struct {
	trie *TrieTreeNode
}

func (ttn *TrieTreeNode) Add(c []rune, index int) {
	if len(c) == index {
		ttn.end = true
		return
	}

	node := ttn.chld[c[index]]
	if node == nil {
		node = &TrieTreeNode{chld: make(map[rune]*TrieTreeNode), end: false}
		ttn.chld[c[index]] = node
	}

	node.Add(c, index+1)
}

func (ttn *TrieTreeNode) Search(c []rune, index int) bool {
	if index == len(c) {
		return ttn.end
	}

	node := ttn.chld[c[index]]
	if node == nil {
		return false
	}

	return node.Search(c, index+1)
}

func (ttn *TrieTreeNode) Delete(c []rune, index int) bool {
	if index == len(c) {
		if !ttn.end {
			return false
		}

		ttn.end = false
		return len(ttn.chld) == 0
	}

	node := ttn.chld[c[index]]
	if node == nil {
		return false
	}

	shouldDelete := node.Delete(c, index+1)
	if shouldDelete {
		delete(ttn.chld, c[index])
		return len(ttn.chld) == 0
	}

	return false
}

func (tt *TrieTree) Add(c []rune) {
	if tt.trie == nil {
		tt.trie = &TrieTreeNode{chld: make(map[rune]*TrieTreeNode), end: false}
	}

	tt.trie.Add(c, 0)
}

func (tt *TrieTree) Search(c string) bool {
	if tt.trie == nil {
		return false
	}

	return tt.trie.Search([]rune(c), 0)
}

func (tt *TrieTree) Delete(c string) bool {
	if tt.trie == nil {
		return false
	}

	return tt.trie.Delete([]rune(c), 0)
}

func main() {
	trie := &TrieTree{}
	keys := []string{"the", "a", "there",
		"answer", "any", "by",
		"bye", "their"}

	for _, val := range keys {
		trie.Add([]rune(val))
	}

	fmt.Println("Search for the word 'the':", trie.Search("the"))
	fmt.Println("Search for the word 'these':", trie.Search("these"))
	fmt.Println("Search for the word 'an':", trie.Search("an"))

	fmt.Println("Delete 'the':", trie.Delete("the"))
	fmt.Println("After delete search for the word 'the':", trie.Search("the"))
}
