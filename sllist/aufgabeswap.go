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
	length := list.Length()
	if pos1 < 0 || pos1 >= length || pos2 < 0 || pos2 >= length {
		return errors.New("error: swap positions out of range")
	}

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
	return nil
}
