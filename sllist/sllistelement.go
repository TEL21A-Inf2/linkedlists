package sllist

import (
	"errors"
	"fmt"
)

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

// Hilfsfunktion: Liefert einen Pointer auf das Ende der Liste.
// Ist nicht für den Gebrauch durch den Client-Programmierer bestimmt,
// aber für den internen Gebrauch nützlich.
func (element *SinglyLinkedListElementInt) GetEnd() *SinglyLinkedListElementInt {
	// Das Ende ist das erste Element, das keinen Nachfolger hat.
	if element.IsEmpty() {
		return element
	}
	return element.next.GetEnd()
}

// Hilfsfunktion: Liefert einen Pointer auf das n-te Element der Liste.
// Ist nicht für den Gebrauch durch den Client-Programmierer bestimmt,
// aber für den internen Gebrauch nützlich.
func (element *SinglyLinkedListElementInt) GetElement(n int) *SinglyLinkedListElementInt {
	if n < 0 || element.IsEmpty() {
		return nil
	}
	if n == 0 {
		return element
	}
	return element.next.GetElement(n - 1)
}

// Hängt ein Element mit der gegeben Zahl als key ans Ende an.
func (element *SinglyLinkedListElementInt) Append(key int) {
	last := element.GetEnd()
	last.key = key
	last.SetNext(NewSinglyLinkedListElementInt())
}

// Liefert den Wert des pos-ten Elements.
func (element *SinglyLinkedListElementInt) GetValue(pos int) (int, error) {
	el := element.GetElement(pos)
	if el == nil {
		return 0, errors.New("error: position out of range")
	}
	return el.key, nil
}

// Fügt ein Element mit dem gegebenen Wert an Stelle pos ein.
// Genauer: Zwischen Stelle pos-1 und pos.
func (element *SinglyLinkedListElementInt) Insert(pos, value int) *SinglyLinkedListElementInt {
	// Wenn pos == 0, dann erzeugen wir hier intern ein Element,
	// tragen die Daten und hängen dort element an.
	// Dieses Element ist der neue Kopf der Liste und muss zurückgegeben werden.
	if pos == 0 {
		d := NewSinglyLinkedListElementInt()
		d.key = value
		d.next = element
		return d
	}

	// Wenn pos == 1, dann wollen wir zwischen element und element.next
	// ein neues Element einfügen.
	if pos == 1 {
		n := NewSinglyLinkedListElementInt()
		n.key = value
		n.next = element.next
		element.next = n
		return element
	}

	// Wenn pos > 1, dann können wir das Problem durch rekursives Einfügen in
	// element.next lösen.
	element.next.Insert(pos-1, value)
	return element
}

// Liefert die Länge der Liste.
func (element SinglyLinkedListElementInt) Length() int {
	if element.IsEmpty() {
		return 0
	}
	return 1 + element.next.Length()
}

// String-Ausgabe von Elementen.
func (element SinglyLinkedListElementInt) String() string {
	return fmt.Sprintf("%v", element.key)
}
