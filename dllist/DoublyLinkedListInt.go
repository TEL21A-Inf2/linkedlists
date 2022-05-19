package dllist

import "errors"

type DoublyLinkedListInt struct {
	dummy  *DoublyLinkedListElementInt
	length int
}

// Konstruktor, erzeugt eine neue Liste.
func NewDoublyLinkedListInt() DoublyLinkedListInt {
	return DoublyLinkedListInt{NewDoublyLinkedListElementInt(), 0}
}

// Fügt einen neuen Wert an Stelle pos ein.
// Liefert einen Fehler, falls pos keine gültige Stelle ist.
// Sonderfall: pos darf == list.Length() sein.
func (list *DoublyLinkedListInt) Insert(pos, value int) error {
	if pos < 0 || pos > list.Length() {
		return errors.New("list position out of bounds")
	}
	if pos == list.length {
		list.Append(value)
		return nil
	}

	element := list.getElement(pos)
	newElement := NewDoublyLinkedListElementInt()
	newElement.key = value

	// Die Sonderfälle sind bereits geprüft, wir fügen mitten in der Liste ein.
	newElement.previous = element.previous
	element.previous = newElement
	newElement.next = element
	newElement.previous.next = newElement
	list.length++
	return nil
}

// Hängt einen neuen Wert ans Ende der Liste an.
func (list *DoublyLinkedListInt) Append(value int) {
	newElement := NewDoublyLinkedListElementInt()
	newElement.key = value

	if list.IsEmpty() {
		list.dummy.next = newElement
		list.dummy.previous = newElement
		newElement.next = list.dummy
		newElement.previous = list.dummy
	} else {
		last := list.getLastElement()
		last.next = newElement
		newElement.previous = last
		newElement.next = list.dummy
		list.dummy.previous = newElement
	}
	list.length++
}

// Liefert die Länge der Liste.
func (list DoublyLinkedListInt) Length() int {
	return list.length
}

// Prüft, ob die Liste leer ist.
func (list DoublyLinkedListInt) IsEmpty() bool {
	return list.Length() == 0
}

// Liefert den Wert an der Stelle pos.
func (list DoublyLinkedListInt) GetValue(pos int) (int, error) {
	element := list.getElement(pos)
	if element == nil {
		return 0, errors.New("list position out of bounds")
	}
	return element.key, nil
}

// String-Ausgabe für verkettete Listen.
func (list DoublyLinkedListInt) String() string {
	result := "[ "
	for current := list.dummy.next; current != list.dummy; current = current.next {
		result += current.String() + " "
	}
	result += "]"
	return result
}
