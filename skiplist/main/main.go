package main

import (
	"fmt"
	. "skiplist"
)

func main() {
	sl := NewSkiplist()
	sl.Add(1)
	fmt.Println(sl.Search(1))
}
