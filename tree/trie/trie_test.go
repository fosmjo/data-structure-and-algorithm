package trie

import (
	"strconv"
	"testing"
)

func newTrie(keys []string) *Trie {
	trie := &Trie{children: make(map[rune]*Trie, 0)}
	for i, k := range keys {
		trie.Insert(k, strconv.Itoa(i))
	}
	return trie
}

// positive case
func TestExistKey(t *testing.T) {
	keys := []string{"A", "to", "tea", "ted", "ten", "i", "in", "inn"}
	trie := newTrie(keys)

	for i, k := range keys {
		if v, err := trie.Find(k); err != nil {
			t.Errorf("cant find value with key %s", k)
		} else {
			if v.value != strconv.Itoa(i) {
				t.Errorf("find wrong value %s with key %s, it should be %s", v.value, k, strconv.Itoa(i))
			}
		}
	}
}

// negative case
func TestNonExistKey(t *testing.T) {
	keys := []string{"A", "to", "tea", "ted", "ten", "i", "in", "inn"}
	trie := newTrie(keys)

	nonExistKeys := []string{"hello", "world"}
	for _, k := range nonExistKeys {
		if _, err := trie.Find(k); err == nil {
			t.Errorf("find value with a non-exist key %s", k)
		}
	}
}
