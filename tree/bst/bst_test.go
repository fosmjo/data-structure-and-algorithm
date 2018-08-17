package bst

import (
	"reflect"
	"testing"
)

func TestBST(t *testing.T) {
	data := []int{6, 2, 1, 4, 3, 8}
	tree := NewTree(data)
	prefixOrder := []int{6, 2, 1, 4, 3, 8}
	infixOrder := []int{1, 2, 3, 4, 6, 8}
	postfixOrder := []int{1, 3, 4, 2, 8, 6}

	if !reflect.DeepEqual(tree.PrefixTraverse(), prefixOrder) {
		t.Error("PrefixTraverse failed")
	}
	if !reflect.DeepEqual(tree.PrefixTraverseIter(), prefixOrder) {
		t.Error("PrefixTraverseIter failed")
	}
	if !reflect.DeepEqual(tree.InfixTraverse(), infixOrder) {
		t.Error("InfixTraverse failed")
	}
	if !reflect.DeepEqual(tree.InfixTraverseIter(), infixOrder) {
		t.Error("InfixTraverseIter failed")
	}
	if !reflect.DeepEqual(tree.PostfixTraverse(), postfixOrder) {
		t.Error("PostfixTraverse failed")
	}
	if !reflect.DeepEqual(tree.PostfixTraverseIter(), postfixOrder) {
		t.Error("PrefixTraverPostfixTraverseIterse failed")
	}
}
