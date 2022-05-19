package dllist

import "fmt"

func ExampleDoublyLinkedListInt_Append() {
	l1 := NewDoublyLinkedListInt()
	l1.Append(42)
	l1.Append(25)
	l1.Append(38)

	fmt.Println(l1)

	// Output:
	// [ 42 25 38 ]
}

func ExampleDoublyLinkedListInt_Insert() {
	// Eine Liste mit ein paar Werten erzeugen.
	l1 := NewDoublyLinkedListInt()
	l1.Append(42)
	l1.Append(25)
	l1.Append(38)

	// Insert an den Stellen 2 und 3 aufrufen.
	l1.Insert(2, 100)
	l1.Insert(3, 1000)
	l1.Insert(0, 55)
	l1.Insert(6, 200)

	// Listenelemente ausgeben:
	fmt.Println(l1)

	// Output:
	// [ 55 42 25 100 1000 38 200 ]
}

func ExampleDoublyLinkedListInt_IsEmpty() {
	l1 := NewDoublyLinkedListInt()
	l2 := NewDoublyLinkedListInt()
	l2.Append(38)

	fmt.Println(l1.IsEmpty())
	fmt.Println(l2.IsEmpty())

	// Output:
	// true
	// false
}

func ExampleDoublyLinkedListInt_GetValue() {
	// Eine Liste mit ein paar Werten erzeugen.
	l1 := NewDoublyLinkedListInt()
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
	// 0 list position out of bounds
	// 0 list position out of bounds
}

func ExampleDoublyLinkedListInt_Length() {
	l1 := NewDoublyLinkedListInt()
	fmt.Println(l1.Length())
	l1.Append(42)
	fmt.Println(l1.Length())
	l1.Append(38)
	fmt.Println(l1.Length())

	// Output:
	// 0
	// 1
	// 2
}

func ExampleDoublyLinkedListElementInt_Swap() {

	l1 := NewDoublyLinkedListInt()

	l1.Append(42)
	l1.Append(25)
	l1.Append(77)
	l1.Append(50)
	l1.Append(103)
	l1.Append(38)
	fmt.Println(l1)

	l1.Swap(2, 4)
	fmt.Println(l1)

	l1.Swap(0, 4)
	fmt.Println(l1)

	err1 := l1.Swap(-1, 2)
	fmt.Println(err1)
	err2 := l1.Swap(0, 10)
	fmt.Println(err2)
	err3 := l1.Swap(-1, 10)
	fmt.Println(err3)

	l1.Swap(2, 3)
	fmt.Println(l1)

	// Output:
	// [ 42 25 77 50 103 38 ]
	// [ 42 25 103 50 77 38 ]
	// [ 77 25 103 50 42 38 ]
	// list position out of bounds
	// list position out of bounds
	// list position out of bounds
	// [ 77 25 50 103 42 38 ]
}