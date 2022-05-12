# Verkettete Listen

Dieses Repo enthält Beispiele und Aufgaben zu verketteten Listen.

Bisher ist eine einfach verkettete Liste implementiert, inkl. Tests.
Für die Listenelemente gibt es einen Datentyp `SinglyLinkedListElementInt`.
Zusätzlich gibt es einen Container-Datentyp `SinglyLinkedListInt`, der einen Pointer
auf den Kopf der Liste speichert.

Eine der Methoden, nämlich `Swap()` ist noch offen.
Diese Funktion ist in einer eigenen Datei `aufgabeswap.go` angelegt, aber noch nicht
implementiert.

## Aufgaben

1. Implementieren Sie die Methode `Swap()` der einfach verketteten Liste.

2. Erweitern Sie die Liste zu einer doppelt verketteten Liste.
    - Definieren Sie neue Datentypen `DoublyLinkedListElementInt` und `DoublyLinkedListInt`.
      Der Element-Datentyp sollte dem der einfach verketteten Liste ähnlich sein,
      zusätzlich aber einen Pointer `previous` besitzen.
    - Duplizieren Sie dann nach und nach die Funktionen aus der einfach verketteten Liste.
      Passen Sie sie dabei jeweils an, so dass sie für eine doppelt verkettete Liste funktionieren.
    - Schreiben Sie entweder Tests oder prüfen Sie zumindest mittels der `main()`-Funktion,
      ob die neuen Datentypen funktionieren.
    - Definieren Sie die neuen Datentypen idealerweise in einem Package `dllist` in einem
      entsprechenden Ordner und binden Sie die in der `main()` nach dem Vorbild der
      einfach verketteten Liste ein.

## Lösungen

Der Branch `loesungen` in diesem Repository enthält bereits eine Lösung für die
`Swap()`-Aufgabe. Wer nicht weiter weiß oder die eigene Lösung abgleichen will,
kann dies über den Branch `loesungen` tun.

Die Lösung in diesem Branch ist allerdings noch nicht optimal, sie iteriert insgesamt
6 Mal über die Liste, um die Aufgabe zu erfüllen. Dies geht auch wesentlich schneller
mit nur einem Durchlauf. Sie können also auch probieren, ob Sie diese Lösung
verbessern können.

Die Lösungen für die weiteren Aufgaben folgen dann später auch in diesem Branch.
