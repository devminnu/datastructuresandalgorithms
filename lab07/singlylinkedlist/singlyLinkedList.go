package singlylinkedlist

import (
	"errors"
	"fmt"
)

//Node for sll node
type Node struct {
	data interface{}
	next *Node
}

//SinglyLinkedList implementation
type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

//New Creates an empty singly linked list
func New() SinglyLinkedList {
	return SinglyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

//Push add node at the end of SinglyLinkedList
func (sll *SinglyLinkedList) Push(data interface{}) (interface{}, error) {
	node := Node{
		data: data,
		next: nil,
	}

	if sll.head == nil {
		sll.head = &node
		sll.tail = &node
	} else {
		sll.tail.next = &node
		sll.tail = &node
	}
	sll.length++
	return data, nil
}

//Pop element from the end of list
func (sll *SinglyLinkedList) Pop() *Node {
	if sll.head == nil {
		return nil
	}
	node := sll.head
	if sll.head == sll.tail {
		sll.head, sll.tail, sll.length = nil, nil, 0
		return node
	}
	for node != nil {
		if node.next == sll.tail {
			sll.tail = node
			node.next = nil
			sll.length--
			return node
		}
		node = node.next
	}
	return nil
}

//Shift remove lement from beginning of SLL
func (sll *SinglyLinkedList) Shift() *Node {
	if sll.head == nil {
		return nil
	}
	node := sll.head
	sll.head = sll.head.next
	sll.length--
	return node
}

//Unshift add element at the beginning
func (sll *SinglyLinkedList) Unshift(data interface{}) *Node {
	node := &Node{
		data: data,
		next: nil,
	}
	if sll.head == nil {
		sll.head, sll.tail, sll.length = node, node, 1
		return node
	}
	node.next = sll.head
	sll.head = node
	sll.length++
	return node
}

//Print print SinglyLinkedList elements
func (sll SinglyLinkedList) Print() {
	if sll.head == nil {
		return
	}
	node := sll.head
	for node != nil {
		fmt.Println("Node:", node.data)
		node = node.next
	}
}

//Get the item at index
func (sll *SinglyLinkedList) Get(index int) (interface{}, error) {
	//Check if Empty
	if sll.head == nil {
		return nil, errors.New("Empty LinkedList")
	}
	//Check valid index: index starts at 0
	if index >= sll.length || index < 0 {
		return nil, errors.New("Invalid Index")
	}

	for i, node := 0, sll.head; node != nil; i++ {
		if i == index {
			return node.data, nil
		}
		node = node.next
	}
	return nil, nil
}

func (sll *SinglyLinkedList) Length() int {
	return sll.length
}
