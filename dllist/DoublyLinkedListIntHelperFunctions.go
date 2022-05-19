package dllist

// Diese Datei enthält private Hilfsfunktionen für DoublyLinkedListInt.
// Ziel ist es, die Dateien etwas übersichtlicher zu halten:
// DoublyLinkedListInt.go enthält das öffentliche Interface des Listen-Datentyps,
// In dieser Datei sind nur Funktionen, die ausschließlich intern verwendet werden.

// Erwartet eine Position und liefert das Element an dieser Stelle, falls es existiert.
// Liefert nil, falls die Position nicht existiert.
func (list *DoublyLinkedListInt) getElement(pos int) *DoublyLinkedListElementInt {
	if pos < 0 || pos >= list.Length() {
		return nil
	}

	current := list.dummy.next
	for i := 0; i != pos; i++ {
		current = current.next
	}
	return current
}

// Liefert das letzte Element der Liste, falls es existiert.
// Liefert nil, falls die Liste leer ist.
func (list *DoublyLinkedListInt) getLastElement() *DoublyLinkedListElementInt {
	if list.IsEmpty() {
		return nil
	}
	return list.dummy.previous
}
