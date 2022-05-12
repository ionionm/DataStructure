package skiplist

import (
	"math/rand"
	"time"
)

const (
	maxLevel int = 31
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Skiplist struct {
	header *SkiplistNode
	length int
	level  int
}

type SkiplistNode struct {
	value int
	next  []*SkiplistNode
	prev  []*SkiplistNode
}

func NewSkiplist() *Skiplist {
	return &Skiplist{
		header: &SkiplistNode{
			next: make([]*SkiplistNode, maxLevel+1),
		},
	}
}

func search(sl *Skiplist, target int) *SkiplistNode {
	cur := sl.header
	for level := sl.level; level >= 0; level-- {
		for cur.next[level] != nil && cur.next[level].value <= target {
			cur = cur.next[level]
		}
		if cur != sl.header && cur.value == target {
			return cur
		}
	}
	return nil
}

//返回target是否存在于跳表中
func (sl *Skiplist) Search(target int) bool {
	return search(sl, target) != nil
}

func randomLevel() int {
	level := 0
	for rand.Intn(2) == 1 {
		level++
	}
	if level > maxLevel {
		level = maxLevel
	}
	return level
}

//插入一个元素到跳表，可以存在重复元素
func (sl *Skiplist) Add(num int) {
	level := randomLevel()
	node := &SkiplistNode{
		value: num,
		next:  make([]*SkiplistNode, level+1),
		prev:  make([]*SkiplistNode, level+1),
	}
	add(sl, node)
}

func add(sl *Skiplist, node *SkiplistNode) {
	updateN := sl.header
	level := len(node.next) - 1
	for i := sl.level; i >= 0; i-- {
		//这里每一层都从上一层找到的节点开始寻找，省去了不少查找过程
		updateN = findClosest(updateN, i, node.value)
		if i <= level {
			node.prev[i] = updateN
			node.next[i] = updateN.next[i]
			updateN.next[i] = node
			if node.next[i] != nil {
				node.next[i].prev[i] = node
			}
		}
	}
	if level > sl.level {
		for i := sl.level + 1; i <= level; i++ {
			sl.header.next[i] = node
			node.prev[i] = sl.header
		}
		sl.level = level
	}
	sl.length++
}

func findClosest(header *SkiplistNode, level, target int) *SkiplistNode {
	for header.next[level] != nil && header.next[level].value < target {
		header = header.next[level]
	}
	return header
}

//在跳表中删除一个值，如果num不存在，返回false，
//如果存在多个num，删除其中任意一个即可。
func (sl *Skiplist) Erase(num int) bool {
	node := search(sl, num)
	if node == nil {
		return false
	}
	for i := len(node.next) - 1; i >= 0; i-- {
		prev, next := node.prev[i], node.next[i]
		prev.next[i] = next
		if next != nil {
			next.prev[i] = prev
		}
	}
	return true
}

func Metadata(sl *Skiplist) []int {
	if sl.length == 0 {
		return nil
	}
	ret := make([]int, 0)
	for p := sl.header.next[0]; p != nil; p = p.next[0] {
		ret = append(ret, p.value)
	}
	return ret
}
