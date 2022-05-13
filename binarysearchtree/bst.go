package binarysearchtree

import (
	"container/list"
	"fmt"
)

type BST struct {
	Root *Node
}

type Node struct {
	Value       int
	Left, Right *Node
}

func NewBST() *BST {
	return &BST{}
}

func (bst *BST) Add(value int) {
	if bst.Root == nil {
		bst.Root = &Node{Value: value}
	} else {
		insertNode(bst.Root, &Node{Value: value})
	}
}

func insertNode(root, node *Node) {
	if node.Value < root.Value {
		if root.Left == nil {
			root.Left = node
		} else {
			insertNode(root.Left, node)
		}
	} else if node.Value > root.Value {
		if root.Right == nil {
			root.Right = node
		} else {
			insertNode(root.Right, node)
		}
	}
}

func (bst *BST) Remove(value int) bool {
	if bst.Root == nil {
		return false
	}
	_, ok := remove(bst.Root, value)
	return ok
}

func remove(root *Node, value int) (*Node, bool) {
	var ok bool
	if value < root.Value {
		root.Left, ok = remove(root.Left, value)
		return root, ok
	} else if value > root.Value {
		root.Right, ok = remove(root.Right, value)
		return root, ok
	}
	if root.Left == nil && root.Right == nil {
		return nil, true
	} else if root.Left == nil {
		return root.Right, true
	} else if root.Right == nil {
		return root.Left, true
	}
	root.Value, _ = min(root.Right)
	root.Right, _ = remove(root.Right, root.Value)
	return root, true
}

func min(root *Node) (int, bool) {
	if root == nil {
		return 0, false
	}
	for {
		if root.Left == nil {
			return root.Value, true
		}
		root = root.Left
	}
}

func (bst *BST) Search(target int) bool {
	return search(bst.Root, target)
}

func search(root *Node, target int) bool {
	if root == nil {
		return false
	}
	if target < root.Value {
		return search(root.Left, target)
	} else if target > root.Value {
		return search(root.Right, target)
	}
	return true
}

func print(i int) {
	fmt.Printf("%d ", i)
}

func (bst *BST) PreOrderInRecursive(f func(int)) {
	if f == nil {
		f = print
	}
	preOrder(bst.Root, f)
}

func preOrder(root *Node, f func(int)) {
	if root == nil {
		return
	}
	f(root.Value)
	preOrder(root.Left, f)
	preOrder(root.Right, f)
}

func (bst *BST) PostOrderInRecursive(f func(int)) {
	if f == nil {
		f = print
	}
	postOrder(bst.Root, f)
}

func postOrder(root *Node, f func(int)) {
	if root == nil {
		return
	}
	postOrder(root.Left, f)
	postOrder(root.Right, f)
	f(root.Value)
}

func (bst *BST) InOrderInRecursive(f func(int)) {
	if f == nil {
		f = print
	}
	inOrder(bst.Root, f)
}

func inOrder(root *Node, f func(int)) {
	if root == nil {
		return
	}
	inOrder(root.Left, f)
	f(root.Value)
	inOrder(root.Right, f)
}

func (bst *BST) PreOrder(f func(int)) {
	if f == nil {
		f = print
	}
	if bst.Root == nil {
		return
	}
	root := bst.Root
	queen := list.New()
	for queen.Len() != 0 || root != nil {
		if root != nil {
			f(root.Value)
			queen.PushBack(root)
			root = root.Left
		} else {
			root = queen.Remove(queen.Back()).(*Node)
			root = root.Right
		}
	}
}

func (bst *BST) PostOrder(f func(int)) {
	if f == nil {
		f = print
	}
	if bst.Root == nil {
		return
	}
	root := bst.Root
	tmp := make([]int, 0)
	stack := make([]*Node, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			tmp = append(tmp, root.Value)
			stack = append(stack, root)
			root = root.Right
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Left
		}
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		f(tmp[i])
	}
}

func (bst *BST) InOrder(f func(int)) {
	if f == nil {
		f = print
	}
	stack := make([]*Node, 0)
	root := bst.Root
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			f(root.Value)
			root = root.Right
		}
	}
}
