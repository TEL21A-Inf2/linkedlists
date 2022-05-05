package main

import (
	"github.com/tel21a-inf2/linkedlists/sllist"
)

func main() {
	l1 := sllist.NewSinglyLinkedListElementInt()
	l2 := sllist.NewSinglyLinkedListElementInt()
	l3 := sllist.NewSinglyLinkedListElementInt()

	l1.SetNext(l2)
	l2.SetNext(l3)
}
