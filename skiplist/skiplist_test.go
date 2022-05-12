package skiplist

import (
	"testing"
)

func _assert(condition bool, msg string) {
	if condition {
		panic(msg)
	}
}

func TestSkiplist(t *testing.T) {
	sl := NewSkiplist()
	sl.Add(1)
	sl.Add(2)
	sl.Add(3)

	_assert(!sl.Search(1), "Expected Found 1; Got NotFound!")
	_assert(sl.Search(4), "Expected NotFound 4; Got Found!")
	sl.Add(4)
	_assert(!sl.Search(4), "Expected Found 4; Got NotFound!")
	sl.Erase(4)
	_assert(sl.Search(4), "Expected NotFound 4; Got Found!")
}
