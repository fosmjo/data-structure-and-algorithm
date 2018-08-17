package trie

import (
	"errors"
	"fmt"
)

// Trie tree Trie
type Trie struct {
	children map[rune]*Trie
	value    string
}

// Find find value by key
func (t *Trie) Find(key string) (*Trie, error) {
	trie := t

	for _, r := range key {
		if _, exist := trie.children[r]; !exist {
			return nil, errors.New("not found")
		}
		trie = trie.children[r]
	}

	return trie, nil
}

// Insert insert key-value pair
func (t *Trie) Insert(key string, value string) {
	var index = -1
	rs := []rune(key)
	trie := t

	for i, r := range rs {
		if _, exist := trie.children[r]; !exist {
			break
		}
		index = i
		trie = trie.children[r]
	}

	for _, r := range rs[index+1:] {
		trie.children[r] = &Trie{children: map[rune]*Trie{}}
		trie = trie.children[r]
	}

	trie.value = value
}

// Display display trie instance
func (t *Trie) Display() {
	queue := []*Trie{t}

	for len(queue) > 0 {
		children := make([]*Trie, 0)

		for _, trie := range queue {
			for k, v := range trie.children {
				fmt.Printf("%s ", string(k))
				children = append(children, v)
			}
		}
		fmt.Println()
		queue = children
	}
}
