package dllist

type DoublyLinkedListElementInt struct {
	key      int
	next     *DoublyLinkedListElementInt
	previous *DoublyLinkedListElementInt
}

func NewDoublyLinkedListElementInt() *DoublyLinkedListElementInt {
	result := DoublyLinkedListElementInt{0, nil, nil}
	result.next = &result
	result.previous = &result
	return &result
}
