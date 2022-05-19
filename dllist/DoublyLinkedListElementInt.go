package dllist

import (
	"fmt"
)

// Datenstruktur für Elemente einer einfach verketteten Liste.
// Die Elemente dieser Liste enthalten ganze Zahlen als Werte.
type DoublyLinkedListElementInt struct {
	key      int
	next     *DoublyLinkedListElementInt
	previous *DoublyLinkedListElementInt
}

// Konstruktor für ein Element einer doppelt verkettete Int-Liste.
func NewDoublyLinkedListElementInt() *DoublyLinkedListElementInt {
	element := &DoublyLinkedListElementInt{0, nil, nil}
	element.next = element
	element.previous = element
	return element
}

// String-Ausgabe von Elementen.
func (element DoublyLinkedListElementInt) String() string {
	return fmt.Sprintf("%v", element.key)
}
