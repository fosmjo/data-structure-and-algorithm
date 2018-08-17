package radixtrie

import (
	"reflect"
	"sort"
	"testing"
)

// keys := []string{"test", "toaster", "toasting", "slow", "slowly"}
// keys := []string{"中国", "中间", "中间人", "中国人", "中国之星", "中国蓝"}

func newRaidxTree(keys []string) *RadixTrie {
	rt := &RadixTrie{
		children: map[string]*RadixTrie{},
	}

	for i, k := range keys {
		rt.Insert(k, i)
	}

	return rt
}

func TestFind(t *testing.T) {
	keys := []string{"test", "toaster", "toasting", "slow", "slowly"}
	rt := newRaidxTree(keys)
	keys = append(keys, []string{"sloxly"}...)

	for i, k := range keys {
		switch k {
		case "sloxly", "x", "sloly", "slowlyx", "slowlx":
			if rt.Find(k) != nil {
				t.Errorf("find answer with key %s", k)
			}
		default:
			if i != rt.Find(k).Value {
				t.Errorf("can't find right answer with key %s", k)
			}
		}

	}
}

func TestDelete(t *testing.T) {
	keys := []string{"test", "toaster", "toasting", "slow", "slowly"}
	rt := newRaidxTree(keys)
	key := "toasting"
	if rt.Find(key) == nil {
		t.Errorf("can't find answer with key %s", key)
	}

	rt.Delete(key)
	if rt.Find(key) != nil {
		t.Errorf("failed to delete key %s", key)
	}
}

func TestAutoComplete(t *testing.T) {
	keys := []string{"test", "toaster", "toasting", "slow", "slowly"}
	predicate := map[string][]string{
		"test":     []string{"test"},
		"toaster":  []string{"toaster"},
		"toasting": []string{"toasting"},
		"slow":     []string{"slow", "slowly"},
		"slowly":   []string{"slowly"},
		"t":        []string{"test", "toaster", "toasting"},
		"to":       []string{"toaster", "toasting"},
		"tx":       []string{},
		"xy":       []string{},
	}
	rt := newRaidxTree(keys)

	for _, k := range keys {
		rs := rt.AutoComplete(k)
		completes := make([]string, 0)
		for _, r := range rs {
			completes = append(completes, r.Key)
		}
		sort.Strings(completes)

		if !reflect.DeepEqual(completes, predicate[k]) {
			t.Errorf("failed to auto complete %s, res:%s", k, completes)
		}
	}
}
