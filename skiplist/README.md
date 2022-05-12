![skiplist](./asserts/skiplist.png)

---



# Structure

每一个跳表节点都有`n`个指向下一个元素和前一个元素的指针，`n`为该节点的层高

```go
type Skiplist struct {
	header *SkiplistNode
	length int						//跳表总元素长度
	level  int						//跳表的最高层数（小bug：删元素不变）
}

type SkiplistNode struct {
	value int						
  //指向下一个元素的指针，因为跳表元素有层数，所以指针也是数组的形式
	next  []*SkiplistNode	
	prev  []*SkiplistNode
}
```

---

# Initialize

跳表最大层数为`31`，头节点不存放元素，头节点指针数组长度为**`最大层高+1（0～maxLevel）`**

```go
const (
	maxLevel int = 31
)

func NewSkiplist() *Skiplist {
	return &Skiplist{
		header: &SkiplistNode{
			next: make([]*SkiplistNode, maxLevel+1),
		},
	}
}
```

---

# Insert

## 获取元素的层高

> 元素层高计算的原则：以更小的概率获得更高的层数

```go
func init() {
	rand.Seed(time.Now().UnixNano())
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
```

初始层数为0，连续获得1的概率越来越低，相应的层数越来越高

---

## Insert

```go
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
    //对于待插入节点层高以内的元素，需要修改至多4处指针引用
		if i <= level {
			node.prev[i] = updateN
			node.next[i] = updateN.next[i]
			updateN.next[i] = node
			if node.next[i] != nil {
				node.next[i].prev[i] = node
			}
		}
	}
  //对于待插入节点层高为新高时，直接修改头节点与待插入节点的引用关系
	if level > sl.level {
		for i := sl.level + 1; i <= level; i++ {
			sl.header.next[i] = node
			node.prev[i] = sl.header
		}
		sl.level = level
	}
  sl.length++
}
```

---

# Search

```go
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
```

---

# Erase

```go
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
```

