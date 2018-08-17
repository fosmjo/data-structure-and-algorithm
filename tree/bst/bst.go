package bst

import (
	"errors"
)

// 二叉查找树，构造，遍历，查找，删除

// Tree 树节点
type Tree struct {
	Data    int
	Left    *Tree
	Right   *Tree
	Visited bool
}

// Find x
func (t *Tree) Find(x int) *Tree {
	if t == nil {
		return nil
	}

	if x < t.Data {
		return t.Left.Find(x)
	} else if x > t.Data {
		return t.Right.Find(x)
	} else {
		return t
	}
}

// FindMin min
func (t *Tree) FindMin() *Tree {
	if t == nil {
		return nil
	}

	for t.Left != nil {
		t = t.Left
	}

	return t
}

// Insert 往树中插入节点
func (t *Tree) Insert(x int) *Tree {
	if t == nil {
		t = &Tree{Data: x}
	} else if x < t.Data {
		t.Left = t.Left.Insert(x)
	} else if x > t.Data {
		t.Right = t.Right.Insert(x)
	} else {
		// pass
	}

	return t
}

// Delete 删除值为x的节点
func (t *Tree) Delete(x int) (*Tree, error) {
	if t == nil {
		return nil, errors.New("element not found")
	}

	if x < t.Data {
		t.Left, _ = t.Left.Delete(x)
	} else if x > t.Data {
		t.Right, _ = t.Right.Delete(x)
	} else if t.Left != nil && t.Right != nil {
		temp := t.Right.FindMin()
		t.Data = temp.Data
		t.Right, _ = t.Right.Delete(t.Data)
	} else {
		// one or zero child
		if t.Left == nil {
			t = t.Right
		} else if t.Right == nil {
			t = t.Left
		}
	}

	return t, nil
}

// PrefixTraverse 前序遍历
func (t *Tree) PrefixTraverse() []int {
	data := make([]int, 0)
	if t == nil {
		return data
	}

	data = append(data, t.Data)
	data = append(data, t.Left.PrefixTraverse()...)
	data = append(data, t.Right.PrefixTraverse()...)

	return data
}

// PrefixTraverseIter 前序遍历
func (t *Tree) PrefixTraverseIter() []int {
	data := make([]int, 0)
	stack := []*Tree{}

	for t != nil {
		// vistit
		data = append(data, t.Data)
		stack = append(stack, t)
		// advance
		t = t.Left

		// drawback
		for t == nil && len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			t = top.Right
		}
	}

	return data
}

// InfixTraverse 中序遍历
func (t *Tree) InfixTraverse() []int {
	data := make([]int, 0)
	if t == nil {
		return data
	}

	data = append(data, t.Left.InfixTraverse()...)
	data = append(data, t.Data)
	data = append(data, t.Right.InfixTraverse()...)

	return data
}

// InfixTraverseIter 中序遍历 迭代
func (t *Tree) InfixTraverseIter() []int {
	data := make([]int, 0)
	stack := []*Tree{}

	for t != nil {
		stack = append(stack, t)
		// advance
		t = t.Left

		for t == nil && len(stack) > 0 {
			// pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// visit
			data = append(data, top.Data)

			// drawback
			t = top.Right
		}
	}

	return data
}

// PostfixTraverse 后序遍历
func (t *Tree) PostfixTraverse() []int {
	data := make([]int, 0)
	if t == nil {
		return data
	}

	data = append(data, t.Left.PostfixTraverse()...)
	data = append(data, t.Right.PostfixTraverse()...)
	data = append(data, t.Data)

	return data
}

// PostfixTraverseIter 后序遍历
func (t *Tree) PostfixTraverseIter() []int {
	data := make([]int, 0)
	stack := []*Tree{}

	for t != nil {
		if !t.Visited {
			stack = append(stack, t)
		}

		t = t.Left

		for (t == nil || t.Visited) && len(stack) > 0 {
			// drawback
			top := stack[len(stack)-1]
			t = top.Right

			// both left and right of top are handled, we handle top now!
			if t == nil || t.Visited {
				// pop
				stack = stack[:len(stack)-1]
				// visit
				data = append(data, top.Data)
				top.Visited = true
			}
		}
	}

	return data
}

// LevelTraverse 层序遍历
func (t *Tree) LevelTraverse() []int {
	data := make([]int, 0)
	if t == nil {
		return data
	}

	queue := []*Tree{t}
	// 每次处理一层
	// 还有一种做法是每次处理一个元素
	for len(queue) > 0 {
		for _, elem := range queue {
			data = append(data, elem.Data)
		}
		oldLen := len(queue)
		for _, elem := range queue {
			if elem.Left != nil {
				queue = append(queue, elem.Left)
			}
			if elem.Right != nil {
				queue = append(queue, elem.Right)
			}
		}
		queue = queue[oldLen:]
	}

	return data
}

// NewTree Tree的构造函数
func NewTree(elems []int) *Tree {
	var t *Tree
	for _, elem := range elems {
		t = t.Insert(elem)
	}
	return t
}
