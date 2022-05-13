package sllist

/*
Aufgabe 1: Implementieren Sie die Methode Swap() für den Datentyp SinglyLinkedListInt.

Die Methode Swap() erwartet zwei Parameter vom Typ int. Dies sind die Positionen
zweier Elemente, die vertauscht werden sollen.

Hinweis: Anders als in der ersten Variante aus der Vorlesung ist diese Methode
für den Container-Datentyp SinglyLinkedListInt definiert und nicht für dessen Elemente.

Hinweis: Beachten Sie in erster Näherung keine Fehlerfälle, d.h. nur gültige Werte
für die Positionen. Die Tests aus der Datei `aufgabeswap_test.go` müssen funktionieren.
Erweitern Sie anschließend die Methode um eine Fehlerbehandlung, d.h. prüfen Sie die
Positionen und liefern Sie einen Wert vom Typ error (vgl. die GetValue()-Methoden).
Passen Sie auch die Tests an bzw. erweitern Sie diese.
*/

func (list *SinglyLinkedListInt) Swap(pos1, pos2 int) {
	if pos1 > pos2 {
		list.Swap(pos2, pos1)
		return
	}

	var A, B, C, D, E, F *SinglyLinkedListElementInt

	pos := 0
	for current := list.head; !current.IsEmpty(); current = current.next {
		if pos == pos1 {
			B = current
		}
		if pos == pos2 {
			E = current
		}
		if pos == pos1-1 {
			A = current
		}
		if pos == pos2-1 {
			D = current
		}
		if pos == pos1+1 {
			C = current
		}
		if pos == pos2+1 {
			F = current
		}
		pos++
	}

	// A := list.head.GetElement(pos1 - 1) // == nil, falls pos1 == 0
	// B := list.head.GetElement(pos1)
	// C := list.head.GetElement(pos1 + 1)
	// D := list.head.GetElement(pos2 - 1)
	// E := list.head.GetElement(pos2)
	// F := list.head.GetElement(pos2 + 1)

	B.SetNext(F)
	E.SetNext(C)
	D.SetNext(B)

	if A != nil {
		A.next = E
	} else {
		list.head = E
	}
}
