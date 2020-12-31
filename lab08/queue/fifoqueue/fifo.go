package fifoqueue

import "fmt"

//Node node of queue
type Node struct {
	value interface{}
	next  *Node
}

//FifoQueue fifo queue
type FifoQueue struct {
	head   *Node
	tail   *Node
	length int
}

//New return fifo queue
func New() *FifoQueue {
	return &FifoQueue{}
}

//Enqueue Add Element at the End of Queue
func (fifo *FifoQueue) Enqueue(value interface{}) interface{} {
	node := &Node{
		value: value,
		next:  nil,
	}
	if fifo.head == nil {
		fifo.head, fifo.tail, fifo.length = node, node, 1
		return value
	}
	fifo.tail.next = node
	fifo.tail = node
	fifo.length++
	return value
}

//Dequeue remove element from beginning of the queue
func (fifo *FifoQueue) Dequeue() interface{} {
	head := fifo.head
	if head == nil {
		return nil
	}

	if head == fifo.tail {
		fifo.head, fifo.tail, fifo.length = nil, nil, 0
		return head.value
	}
	fifo.head = fifo.head.next
	fifo.length--
	return head.value
}

//Print all nodes in a queue
func (node *Node) Print() {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	node = node.next
	node.Print()
}

func (fifo *FifoQueue) Head() *Node {
	return fifo.head
}
