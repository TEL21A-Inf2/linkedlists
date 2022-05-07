package sllist

import "errors"

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
func (element *SinglyLinkedListElementInt) GetEnd() *SinglyLinkedListElementInt {
	// Das Ende ist das erste Element, das keinen Nachfolger hat.
	if element.IsEmpty() {
		return element
	}
	return element.next.GetEnd()
}

// Hängt ein Element mit der gegeben Zahl als key ans Ende an.
func (element *SinglyLinkedListElementInt) Append(key int) {
	last := element.GetEnd()
	last.key = key
	last.SetNext(NewSinglyLinkedListElementInt())
}

// Liefert den Wert des pos-ten Elements.
func (element *SinglyLinkedListElementInt) GetValue(pos int) (int, error) {
	if element.IsEmpty() {
		return 0, errors.New("ungültige Position für Listenzugriff")
	}
	if pos == 0 {
		return element.key, nil
	}
	return element.next.GetValue(pos - 1)
}