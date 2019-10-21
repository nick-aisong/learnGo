/*
1. Make use of the package container/list to create a (doubly) linked list.
Push the values 1, 2 and 4 to the list and then print it.
2. Create your own linked list implementation.
And perform the same actions as in question 1
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(4)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}
}
