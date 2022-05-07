package main

import (
	"fmt"

	"github.com/tel21a-inf2/linkedlists/sllist"
)

func main() {
	l1 := sllist.NewSinglyLinkedListInt()

	l1.Append(42)
	l1.Append(25)
	l1.Append(77)
	l1.Append(38)

	fmt.Println(l1)
}
