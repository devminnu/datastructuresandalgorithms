package main

import (
	"lab08/heap/maxbinaryheap"
)

func main() {
	// bst := binarysearchtree.New()
	// bst.Insert(10)
	// bst.Insert(8)
	// bst.Insert(9)
	// bst.Insert(7)
	// bst.Insert(15)
	// bst.Insert(12)
	// bst.Insert(16)
	// bst.Insert(14)
	// fmt.Println("BFS:         ", bst.BFS())
	// fmt.Println("DFSPreOrder: ", bst.DFSPreOrder())
	// fmt.Println("DFSPostOrder:", bst.DFSPostOrder())
	// fmt.Println("DFSInOrder:  ", bst.DFSInOrder())

	//					     	10
	//					8							15
	//			7				9			12			16
	//						  					14

	// // fmt.Println(bst.Root().Get(16))
	// bst.Insert(16)
	// // fmt.Println(bst.Root().Get(16))

	// bst.Root().Print()

	// fifo := fifoqueue.New()
	// fifo.Enqueue(10)
	// fifo.Enqueue(20)
	// fifo.Enqueue(30)
	// fifo.Enqueue(50)
	// fifo.Head().Print()
	// fmt.Println("DE", fifo.Dequeue())
	// fmt.Println("DE", fifo.Dequeue())
	// fmt.Println("DE", fifo.Dequeue())
	// fmt.Println("DE", fifo.Dequeue())
	// fmt.Println("DE", fifo.Dequeue())
	// fifo.Head().Print()
	// fifo.Enqueue(50)
	// fifo.Enqueue(60)
	// fifo.Head().Print()

	/*
							  	100
					   50				    90
				  40	  30		  80 	 20
		  	35  10			  25

	*/
	mbh := maxbinaryheap.New()
	mbh.Insert(100)
	mbh.Insert(50)
	mbh.Insert(90)
	mbh.Insert(40)
	mbh.Insert(30)
	mbh.Insert(80)
	mbh.Insert(20)
	mbh.Insert(35)
	mbh.Insert(10)
	mbh.Insert(25)
	mbh.Insert(500)
	mbh.Insert(300)
	mbh.Insert(400)
	mbh.PrintHeap()
}
