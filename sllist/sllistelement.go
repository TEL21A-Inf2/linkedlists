package sllist

// Datenstruktur für Elemente einer einfach verketteten Liste.
// Die Elemente dieser Liste enthalten ganze Zahlen als Werte.
type SinglyLinkedListElementInt struct {
	key  int
	next *SinglyLinkedListElementInt
}

// Konstruktor für eine Element einer einfach verkettete Int-Liste.
func NewSinglyLinkedListElementInt() *SinglyLinkedListElementInt {
	return &SinglyLinkedListElementInt{0, nil}
}

// Liefert true, falls dieses Listenelement leer ist.
// Ein Element gilt als leer, falls es keinen Nachfolger hat.
func (element SinglyLinkedListElementInt) IsEmpty() bool {
	return element.next == nil
}

// Setzt den Nachfolger des Listenelements.
func (element *SinglyLinkedListElementInt) SetNext(next *SinglyLinkedListElementInt) {
	element.next = next
}

// Liefert das Ende der Liste.
func (element SinglyLinkedListElementInt) GetEnd() *SinglyLinkedListElementInt {
	// Das Ende ist das erste Element, das keinen Nachfolger hat.
	if element.IsEmpty() {
		return &element
	}
	return element.next.GetEnd()
}
