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

// String-Ausgabe für verkettete Listen.
func (list SinglyLinkedListInt) String() string {
	result := "[ "
	for current := list.head; !current.IsEmpty(); current = current.next {
		result += current.String() + " "
	}
	result += "]"
	return result
}

// Hängt ein Element mit dem gegebenen Wert an.
func (list *SinglyLinkedListInt) Append(value int) {
	list.head.Append(value)
}

// Fügt ein Element ein.
func (list *SinglyLinkedListInt) Insert(pos, value int) {
	list.head = list.head.Insert(pos, value)
}

// Prüft, ob die Liste leer ist.
func (list SinglyLinkedListInt) IsEmpty() bool {
	return list.head.IsEmpty()
}

// Liefert den Wert an der Stelle pos.
func (list SinglyLinkedListInt) GetValue(pos int) (int, error) {
	return list.head.GetValue(pos)
}
