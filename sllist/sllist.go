package sllist

// Datentyp für einfach verkettete Listen.
// Ein Struct, das einen Pointer auf den Kopf der Liste enthält.
type SinglyLinkedListInt struct {
	head *SinglyLinkedListElementInt
}

// Konstruktor für eine einfach verkettete Liste.
func NewSinglyLinkedListInt() SinglyLinkedListInt {
	return SinglyLinkedListInt{NewSinglyLinkedListElementInt()}
}

// Hängt ein Element mit dem gegebenen Wert an.
func (list *SinglyLinkedListInt) Append(value int) {
	list.head.Append(value)
}
