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
