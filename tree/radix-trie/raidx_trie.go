package radixtrie

import (
	"fmt"
	"strings"
)

// RadixTrie radix trie node
type RadixTrie struct {
	parent   *RadixTrie
	children map[string]*RadixTrie
	Value    int
}

// Result auto complete result
type Result struct {
	Key   string
	Value int
}

// Find find
func (rt *RadixTrie) Find(key string) *RadixTrie {
	edge, child, prefix := rt.findChildWithLongestPrefix(key)
	if child == nil || edge != prefix {
		return nil
	}
	// 叶子节点
	if len(child.children) == 0 {
		return child
	}

	return child.Find(key[len(edge):])
}

// AutoComplete auto complete
func (rt *RadixTrie) AutoComplete(prefix string) []*Result {
	rs := make([]*Result, 0)
	descendant := rt.findRadixTrie(prefix)

	if descendant == nil {
		return rs
	}

	if len(descendant.children) == 0 {
		rs = append(rs, &Result{Key: prefix, Value: descendant.Value})
		return rs
	}

	rs = descendant.Expand()
	// map
	for _, r := range rs {
		r.Key = prefix + r.Key
	}
	return rs
}

// Expand expand
func (rt *RadixTrie) Expand() []*Result {
	rs := make([]*Result, 0)

	for edge, child := range rt.children {
		if len(child.children) == 0 {
			rs = append(rs, &Result{Key: edge, Value: child.Value})
		} else {
			rets := child.Expand()
			// map
			for _, r := range rets {
				r.Key = edge + r.Key
			}
			rs = append(rs, rets...)
		}
	}

	return rs
}

// Insert insert
func (rt *RadixTrie) Insert(key string, value int) {
	edge, child, prefix := rt.findChildWithLongestPrefix(key)

	if child == nil {
		rt.children[key] = &RadixTrie{children: map[string]*RadixTrie{}, parent: rt, Value: value}
	} else {
		if strings.HasPrefix(key, edge) {
			// 叶子节点
			if len(child.children) == 0 {
				node := &RadixTrie{
					children: map[string]*RadixTrie{},
					parent:   rt,
				}

				child.parent = node
				node.children[""] = child

				node.children[key[len(edge):]] = &RadixTrie{
					children: map[string]*RadixTrie{},
					parent:   node,
					Value:    value,
				}

				rt.children[edge] = node
			} else {
				child.Insert(key[len(edge):], value)
			}
		} else if strings.HasPrefix(edge, key) {
			// 断旧边
			delete(rt.children, edge)

			// 新节点
			node := &RadixTrie{
				children: map[string]*RadixTrie{},
				parent:   rt,
			}

			// 旧节点
			child.parent = node
			node.children[edge[len(key):]] = child

			// 新插入节点，必然会引入这个叶子节点
			node.children[""] = &RadixTrie{
				children: map[string]*RadixTrie{},
				parent:   node,
				Value:    value,
			}

			rt.children[key] = node
		} else {
			node := &RadixTrie{
				children: map[string]*RadixTrie{},
				parent:   rt,
			}
			child.parent = node
			node.children[edge[len(prefix):]] = child

			node.children[key[len(prefix):]] = &RadixTrie{
				children: map[string]*RadixTrie{},
				Value:    value,
				parent:   node,
			}

			rt.children[prefix] = node
			delete(rt.children, edge)
		}
	}
}

func (rt *RadixTrie) findChildWithLongestPrefix(key string) (string, *RadixTrie, string) {
	var edge string
	var raidxTrie *RadixTrie
	var prefix string
	keyr := []rune(key)

	for k, v := range rt.children {
		if len(k) == 0 && len(key) == 0 {
			return "", v, ""
		}
		if len(k) == 0 || len(key) == 0 {
			continue
		}
		kr := []rune(k)
		if kr[0] == keyr[0] {
			edge, raidxTrie = k, v

			// find longest prefix
			index := 0
			for i := 1; i < len(kr) && i < len(keyr); i++ {
				if kr[i] != keyr[i] {
					break
				}
				index = i
			}
			prefix = string(keyr[0 : index+1])
		}
	}

	return edge, raidxTrie, prefix
}

// key对应一个值且又是另一个key的前缀时，返回内部节点，不返回空边的叶子节点=》优先返回内部节点
func (rt *RadixTrie) findRadixTrie(key string) *RadixTrie {
	edge, child, prefix := rt.findChildWithLongestPrefix(key)
	if child == nil || edge != prefix {
		return nil
	}
	// 叶子节点 || 内部节点
	if len(child.children) == 0 || prefix == key {
		return child
	}

	return child.findRadixTrie(key[len(edge):])
}

// Display display
func (rt *RadixTrie) Display() {
	queue := []*RadixTrie{rt}

	for len(queue) > 0 {
		children := make([]*RadixTrie, 0)

		for _, node := range queue {
			for k, v := range node.children {
				fmt.Printf("%s ", string(k))
				children = append(children, v)
			}
		}
		fmt.Println()

		queue = children
	}
}

// Delete delete
func (rt *RadixTrie) Delete(key string) {
	edge, child, _ := rt.findChildWithLongestPrefix(key)
	if child == nil {
		return
	}

	if len(child.children) == 0 {
		// delete
		delete(rt.children, key)

		// merge
		if len(rt.children) == 1 && rt.parent != nil {
			var key string
			var node *RadixTrie
			for k, v := range rt.children {
				key, node = k, v
			}
			// update rt's value
			delete(rt.children, key)
			rt.Value = node.Value

			// update rt's incoming edge
			parent := rt.parent
			var edge string
			for k, v := range parent.children {
				if v == rt {
					edge = k
					break
				}
			}

			delete(parent.children, edge)
			parent.children[edge+key] = rt
		}
		return
	}

	child.Delete(key[len(edge):])
}
