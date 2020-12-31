package main

import (
	"datastructuresandalgorithms/lab07/singlylinkedlist"
	"fmt"
)

func main() {
	sll := singlylinkedlist.New()
	// sll.Push(10)
	// sll.Push(20)
	// sll.Push(30)
	// sll.Push(40)
	// fmt.Println("Length:", sll.Length())
	// sll.Print()
	sll.Pop()
	sll.Pop()
	sll.Shift()
	// fmt.Println("Length:", sll.Length())
	// sll.Print()
	sll.Unshift(50)
	sll.Unshift(60)
	// fmt.Println("Length:", sll.Length())
	sll.Print()
	fmt.Println(sll.Get(1))

}
