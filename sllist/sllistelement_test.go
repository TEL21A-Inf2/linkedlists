package sllist

import "fmt"

func ExampleSinglyLinkedListElementInt_IsEmpty() {

	// Eine neue leere Liste erzeugen.
	l1 := NewSinglyLinkedListElementInt()

	// l1.IsEmpty() sollte true sein.
	fmt.Println(l1.IsEmpty())

	// Ein weiteres Element erzeugen, auch dieses ist leer.
	l2 := NewSinglyLinkedListElementInt()
	fmt.Println(l2.IsEmpty())

	// l2 an l1 anhängen.
	l1.SetNext(l2)

	// l1 ist nun nicht mehr leer, l2 ist es immer noch.
	fmt.Println(l1.IsEmpty())
	fmt.Println(l2.IsEmpty())

	// Output:
	// true
	// true
	// false
	// true
}

func ExampleSinglyLinkedListElementInt_GetLast() {

	// Eine Liste erzeugen.
	l1 := NewSinglyLinkedListElementInt()
	l2 := NewSinglyLinkedListElementInt()
	l3 := NewSinglyLinkedListElementInt()
	l1.SetNext(l2)
	l2.SetNext(l3)

	// Letztes Element aller drei Elemente sollte l3 sein.
	fmt.Println(*l1.GetEnd() == *l3)
	fmt.Println(*l2.GetEnd() == *l3)
	fmt.Println(*l3.GetEnd() == *l3)

	// Output:
	// true
	// true
	// true
}

func ExampleSinglyLinkedListElementInt_Append() {

	// Ein einzelnes Listenelement erzeugen.
	l1 := NewSinglyLinkedListElementInt()

	// Nacheinander drei Zahlen anhängen.
	l1.Append(42)
	l1.Append(25)
	l1.Append(38)

	// Die Elemente sollten nun über l1 erreichbar sein:
	fmt.Println(l1.key)
	fmt.Println(l1.next.key)
	fmt.Println(l1.next.next.key)

	// Output:
	// 42
	// 25
	// 38
}

func ExampleSinglyLinkedListElementInt_GetValue() {
	// Eine Liste mit ein paar Werten erzeugen.
	l1 := NewSinglyLinkedListElementInt()
	l1.Append(42)
	l1.Append(25)
	l1.Append(38)

	// Die folgenden GetValue()-Abfragen sollten die jeweiligen Werte liefern und die
	// Fehler sollten jeweils nil sein.
	fmt.Println(l1.GetValue(0))
	fmt.Println(l1.GetValue(1))
	fmt.Println(l1.GetValue(2))

	// Ungültige Zugriffe sollten 0 liefern und der Fehler sollte nicht nil sein.
	fmt.Println(l1.GetValue(-1))
	fmt.Println(l1.GetValue(3))

	// Output:
	// 42 <nil>
	// 25 <nil>
	// 38 <nil>
	// 0 ungültige Position für Listenzugriff
	// 0 ungültige Position für Listenzugriff
}
