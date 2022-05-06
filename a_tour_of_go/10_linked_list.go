package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type LinkedList[T any] struct {
	next *LinkedList[T]
	val  T
}

func (ll *LinkedList[T]) Len() int {
	length := 0
	for node := ll; node != nil; node = node.next {
		length++
	}
	return length
}

func (ll *LinkedList[T]) Append(value T) *LinkedList[T] {
	return ll.Insert(ll.Len(), value)
}

func (ll *LinkedList[T]) Insert(pos int, value T) *LinkedList[T] {
	if ll == nil || pos < 0 {
		return &LinkedList[T]{
			val:  value,
			next: ll,
		}
	} else if pos > ll.Len() {
		ll.Append(value)
	} else {
		length := 0
		node:=ll
		for length <= pos - 1; node = node.next {
			length++
				if length == pos {
					
				}
			return
		}


	}
	fmt.Printf("value: %v position: %v\n", ll.val, pos)
	ll.next = ll.next.Insert(pos-1, value)
	
	return ll
}

func (ll *LinkedList[T]) String() string {
	if ll == nil {
		return "end"
	}
	return fmt.Sprintf("%v -> %v", ll.val, ll.next.String())
}

func main() {
	var head *LinkedList[string]
	head = head.Append("Hello")
	fmt.Println(head.Len())
	head = head.Append("This")
	head = head.Append("Is")
	head.Insert(1, "Insert")
	head = head.Append("Go")
	head = head.Append("Practice")
	fmt.Println(head)
	head.String()
	var intLinkedList *LinkedList[int]
	intLinkedList = intLinkedList.Append(1)
	fmt.Println(intLinkedList.Len())
	intLinkedList = intLinkedList.Append(2)
	intLinkedList = intLinkedList.Append(3)
	intLinkedList.String()
	fmt.Println(intLinkedList.Len())

}
