package arraylist

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25) //shrink when size is 0.25 of capacity.
)

//List's elements shouln't export, because the main purpose of
//List is to provide some universal interface. If you want access
//elements of List, why not use slice directly?
type List struct {
	elements []interface{}
	size     int
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	return list.elements[index], true
}

func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()
}

func (list *List) growBy(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n > currentCapacity {
		newCapacity := int(float32((currentCapacity + n)) * growthFactor)
		list.resize(newCapacity)
	}
}

func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}
