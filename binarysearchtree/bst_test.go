package binarysearchtree

import (
	"fmt"
	"testing"
)

func _assert(condition bool, msg string, v ...interface{}) {
	if condition {
		panic(fmt.Sprintf("assertion failed: "+msg, v))
	}
}

func TestSearchAndRemove(t *testing.T) {
	bst := NewBST()
	bst.Add(3)
	bst.Add(1)
	bst.Add(5)
	bst.Add(4)
	bst.Add(6)
	_assert(!bst.Search(6), "Search 6: Need true; Got false")
	_assert(bst.Search(0), "Search 0: Need false; Got true")
	bst.Remove(5)
	_assert(bst.Search(5), "Search 5: Need false; Got true")
	pre := []int{3, 1, 6, 4}
	_pre := make([]int, 0)
	bst.PreOrderInRecursive(func(i int) {
		_pre = append(_pre, i)
	})
	_assert(!equals(pre, _pre), "Need: %v; Got: %v", pre, _pre)
}

func TestRecursive(t *testing.T) {
	bst := NewBST()
	//  	3
	//	 1		5
	// 0  2	  4	  6
	bst.Add(3)
	bst.Add(1)
	bst.Add(0)
	bst.Add(2)
	bst.Add(5)
	bst.Add(4)
	bst.Add(6)
	pre := []int{3, 1, 0, 2, 5, 4, 6}
	in := []int{0, 1, 2, 3, 4, 5, 6}
	post := []int{0, 2, 1, 4, 6, 5, 3}
	_pre := make([]int, 0)
	_in := make([]int, 0)
	_post := make([]int, 0)
	bst.PreOrderInRecursive(func(i int) {
		_pre = append(_pre, i)
	})
	bst.InOrderInRecursive(func(i int) {
		_in = append(_in, i)
	})
	bst.PostOrderInRecursive(func(i int) {
		_post = append(_post, i)
	})
	_assert(!equals(pre, _pre), "PreOrder: Need %v; Got %v", pre, _pre)
	_assert(!equals(in, _in), "InOrder: Need %v; Got %v", in, _in)
	_assert(!equals(post, _post), "PostOrder: Need %v; Got %v", post, _post)
}

func equals(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestTraverseNonRecursive(t *testing.T) {
	bst := NewBST()
	bst.Add(3)
	bst.Add(1)
	bst.Add(0)
	bst.Add(2)
	bst.Add(5)
	bst.Add(4)
	bst.Add(6)
	pre := []int{3, 1, 0, 2, 5, 4, 6}
	in := []int{0, 1, 2, 3, 4, 5, 6}
	post := []int{0, 2, 1, 4, 6, 5, 3}
	_pre := make([]int, 0)
	_in := make([]int, 0)
	_post := make([]int, 0)
	bst.PreOrder(func(i int) {
		_pre = append(_pre, i)
	})
	bst.InOrder(func(i int) {
		_in = append(_in, i)
	})
	bst.PostOrder(func(i int) {
		_post = append(_post, i)
	})
	_assert(!equals(pre, _pre), "PreOrder: Need %v; Got %v", pre, _pre)
	_assert(!equals(in, _in), "InOrder: Need %v; Got %v", in, _in)
	_assert(!equals(post, _post), "PostOrder: Need %v; Got %v", post, _post)
}
