package sllist

import "fmt"

func ExampleSinglyLinkedListElementInt_Swap() {

	l1 := NewSinglyLinkedListInt()

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
	// error: swap positions out of range
	// error: swap positions out of range
	// error: swap positions out of range
	// [ 77 25 50 103 42 38 ]
}
