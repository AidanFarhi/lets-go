package main

import "fmt"

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

func NewNode[T comparable](element T) *Node[T] {
	return &Node[T]{
		Val:  element,
		Next: nil,
	}
}

type LinkedList[T comparable] struct {
	Head *Node[T]
	Size int
}

func (ll *LinkedList[T]) Add(element T) {
	newNode := NewNode(element)
	if ll.Size == 0 {
		ll.Head = newNode
	} else {
		newNode.Next = ll.Head
		ll.Head = newNode
	}
	ll.Size++
}

func (ll *LinkedList[T]) Insert(element T, position int) {
	count := 0
	newNode := NewNode(element)
	current := ll.Head
	if position == 0 {
		ll.Add(element)
	} else {
		for count < ll.Size {
			if count == position-1 {
				newNode.Next = current.Next
				current.Next = newNode
				break
			}
			current = current.Next
			count++
		}
		ll.Size++
	}
}

func (ll *LinkedList[T]) Index(element T) int {
	index := -1
	count := 0
	n := ll.Head
	for n != nil {
		if n.Val == element {
			index = count
			break
		}
		n = n.Next
		count++
	}
	return index
}

func main() {

	ll := LinkedList[int]{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	ll.Insert(99, 1)
	ll.Insert(77, 4)
	ll.Insert(555, 0)

	n := ll.Head
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}

	fmt.Println("Index of 20:", ll.Index(20))
	fmt.Println("Index of 555:", ll.Index(555))
}
