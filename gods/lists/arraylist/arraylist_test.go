package arraylist

import (
	"fmt"
	"testing"
)

func _assert(condition bool, mas string, v ...interface{}) {
	if condition {
		panic(fmt.Sprintf("Assertion failed: "+mas, v...))
	}
}

func TestGet(t *testing.T) {
	list := New()
	list.Add(1, 2, "3", "a", "b")
	_, ok := list.Get(1)
	_assert(!ok, "Get(1): Need true, Got false")
	s, _ := list.Get(3)
	_assert(s != "a", "Get(3): Need a, Got %s", s)
}
