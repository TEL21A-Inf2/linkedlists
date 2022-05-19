package sllist

import "errors"

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

func (list *SinglyLinkedListInt) Swap(pos1, pos2 int) error {

	// "Organisatorisches": Sonderfälle und umgedrehte Listenpositionen abfangen.
	if pos1 == pos2 {
		return nil
	}
	length := list.Length()
	if pos1 < 0 || pos1 >= length || pos2 < 0 || pos2 >= length {
		return errors.New("error: swap positions out of range")
	}
	if pos1 > pos2 {
		return list.Swap(pos2, pos1)
	}

	// Was übrig bleibt, sind zwei Fälle:
	// 1. pos1 und pos2 liegen direkt nebeneinander.
	// 2. pos1 und pos2 liegen nicht direkt nebeneinander.
	// Für diese beiden Fälle gibt es jeweils eine Spezialfunktion,
	// die wir hier aufrufen.
	if pos1 == pos2-1 {
		list.swapNext(pos1)
		return nil
	}
	list.swapNonAdjacent(pos1, pos2)
	return nil
}

// Intern benutzte Hilfsmethode, um zwei beliebige nicht benachbarte Elemente
// zu vertauschen. Es gibt keine Sicherheitsprüfungen und keine Fehlerbehandlung.
// Die Methode darf nur aufgerufen werden, wenn sicher ist, dass sie auch funktioniert.
func (list *SinglyLinkedListInt) swapNonAdjacent(pos1, pos2 int) {

	// Wir holen uns die Elemente an den Stellen pos1 und pos2
	// und nennen sie B und E:
	B := list.head.GetElement(pos1)
	E := list.head.GetElement(pos2)

	// Nun holen wir uns die Vorgänger und nennen sie A und D:
	A := list.head.GetElement(pos1 - 1)
	D := list.head.GetElement(pos2 - 1)

	// Die Nachfolger nennen wir C und F:
	C := list.head.GetElement(pos1 + 1)
	F := list.head.GetElement(pos2 + 1)

	/* Wir haben jetzt folgende Listenstruktur:

	... -> A -> B -> C -> ... -> D -> E -> F -> ...

	B und E sollen getauscht werden. Die nötigen Pointer-Verbiegungen
	können wir jetzt ganz leicht durchführen:
	*/
	B.next = F
	E.next = C
	D.next = B

	// Sonderfall: Falls pos1 == 0, dann ist A == nil.
	// Dies müssen wir prüfen und in diesem Fall stattdessen list.head anpassen:
	if A != nil {
		A.next = E
	} else {
		list.head = E
	}
}

// Intern benutzte Hilfsmethode, um ein Element mit seinem Nachfolger
// zu vertauschen. Es gibt keine Sicherheitsprüfungen und keine Fehlerbehandlung.
// Die Methode darf nur aufgerufen werden, wenn sicher ist, dass sie auch funktioniert.
func (list *SinglyLinkedListInt) swapNext(pos1 int) {

	// Wir holen uns die beiden Elemente und ihre Vorgänger und Nachfolger.
	A := list.head.GetElement(pos1 - 1)
	B := list.head.GetElement(pos1)
	C := list.head.GetElement(pos1 + 1)
	D := list.head.GetElement(pos1 + 2)

	/* Wir haben jetzt folgende Listenstruktur:

	... -> A -> B -> C -> D -> ...

	B und C sollen getauscht werden.
	Wir verbiegen die Pointer.
	*/
	B.next = D
	C.next = B

	// Für A gibt es wieder den Sonderfall, falls B der Kopf der Liste ist.
	if A != nil {
		A.next = C
	} else {
		list.head = C
	}
}
