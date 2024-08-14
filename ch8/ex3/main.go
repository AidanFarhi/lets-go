package main

import "fmt"

/*
Write a generic singly linked list data type. Each element can hold a
comparable value and has a pointer to the next element in the list.
The methods to implement are as follows:

// adds a new element to the end of the linked list
Add(T)
// adds an element at the specified position in the linked list
Insert(T, int)
// returns the position of the supplied value, -1 if it's not present
Index (T) int
*/

type LinkedList[T comparable] struct {
	head *Node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil}
}

func (ll *LinkedList[T]) Add(v T) {
	newNode := NewNode(v)
	if ll.head == nil {
		ll.head = newNode
		return
	}
	n := ll.head
	for n.next != nil {
		n = n.next
	}
	n.next = newNode
}

func (ll *LinkedList[T]) Insert(v T, index int) {
	newNode := NewNode(v)
	if ll.head == nil {
		ll.head = newNode
		return
	}
	if index == 0 {
		newNode.next = ll.head
		ll.head = newNode
		return
	}
	currIndex := 0
	n := ll.head
	for n.next != nil {
		if currIndex == index-1 {
			break
		}
		currIndex++
		n = n.next
	}
	newNode.next = n.next
	n.next = newNode
}

func (ll *LinkedList[T]) Index(v T) int {
	index := -1
	n := ll.head
	currIndex := 0
	for n != nil {
		if n.val == v {
			index = currIndex
			break
		}
		n = n.next
		currIndex++
	}
	return index
}

func (ll *LinkedList[T]) NodeList() []T {
	var nodes []T
	n := ll.head
	for n != nil {
		nodes = append(nodes, n.val)
		n = n.next
	}
	return nodes
}

type Node[T comparable] struct {
	val  T
	next *Node[T]
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{val: val, next: nil}
}

func main() {
	intList := NewLinkedList[int]()
	for _, n := range []int{1, 2, 3, 4, 5} {
		intList.Add(n)
	}
	fmt.Println(intList.NodeList())

	stringList := NewLinkedList[string]()
	for _, n := range []string{"1", "2", "3", "4", "5"} {
		stringList.Add(n)
	}
	fmt.Println(stringList.NodeList())

	floatList := NewLinkedList[float64]()
	floatList.Insert(10.23, 5)
	fmt.Println(floatList.NodeList())
	floatList.Insert(53.22, 0)
	fmt.Println(floatList.NodeList())
	for _, n := range []float64{52.33, 75.23, 88.23, 462.32} {
		floatList.Add(n)
	}
	fmt.Println(floatList.NodeList())
	floatList.Insert(99.99, 5)
	fmt.Println(floatList.NodeList())
	floatList.Insert(77777.2323, 99)
	fmt.Println(floatList.NodeList())

	fmt.Println(stringList.Index("3"))
	fmt.Println(stringList.Index("5"))
	fmt.Println(stringList.Index("1"))
	fmt.Println(stringList.Index("99"))
}
