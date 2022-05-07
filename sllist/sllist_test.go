package sllist

import "fmt"

func ExampleSinglyLinkedListInt_Append() {
	l1 := NewSinglyLinkedListInt()
	l1.Append(42)
	l1.Append(25)
	l1.Append(38)

	for current := l1.head; !current.IsEmpty(); current = current.next {
		fmt.Println(current)
	}

	// Output:
	// 42
	// 25
	// 38
}
