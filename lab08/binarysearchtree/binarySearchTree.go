package binarysearchtree

import (
	"fmt"
	"lab08/queue/fifoqueue"
)

//Node Tree Node
type Node struct {
	value int
	left  *Node
	right *Node
}

//BinarySearchTree bst
type BinarySearchTree struct {
	root *Node
}

//New returns new binary search tree
func New() BinarySearchTree {
	return BinarySearchTree{
		root: nil,
	}
}

//Insert insert node
func (bst *BinarySearchTree) Insert(value int) {
	node := &Node{
		value: value,
		left:  nil,
		right: nil,
	}
	if bst.root == nil {
		bst.root = node
		return
	}
	root := bst.root
	for {
		if value == root.value {
			return
		}
		if value < root.value {
			if root.left == nil {
				root.left = node
				break
			}
			root = root.left
			continue
		} else {
			if root.right == nil {
				root.right = node
				break
			}
			root = root.right
		}
	}
}

//Print print bst
func (root *Node) Print() interface{} {
	if root == nil {
		return nil
	}
	fmt.Println(root.value)
	root = root.left
	return root.Print()
}

//Get Find the node in a tree
func (root *Node) Get(value int) bool {
	if root == nil {
		return false
	}
	if root.value == value {
		return true
	} else if value < root.value {
		root = root.left
		return root.Get(value)
	} else {
		root = root.right
		return root.Get(value)
	}
	// return false
}

//Root Get tree root
func (bst *BinarySearchTree) Root() *Node {
	return bst.root
}

//BFS Breadth First Search
func (bst *BinarySearchTree) BFS() (visitedNodeList []interface{}) {
	fifo := fifoqueue.New()
	if bst.root == nil {
		return visitedNodeList
	}

	fifo.Enqueue(bst.root)

	for node := fifo.Dequeue().(*Node); node != nil; {
		visitedNodeList = append(visitedNodeList, node.value)
		if node.left != nil {
			fifo.Enqueue(node.left)
		}
		if node.right != nil {
			fifo.Enqueue(node.right)
		}
		v := fifo.Dequeue()
		if v != nil {
			node = v.(*Node)
		} else {
			node = nil
		}
	}
	return visitedNodeList
}

//DFSPreOrder DFSPreOrder
func (bst *BinarySearchTree) DFSPreOrder() (visitedNodeList []interface{}) {
	if bst.root == nil {
		return
	}
	root := bst.root
	var traverse func(root *Node)
	traverse = func(root *Node) {
		visitedNodeList = append(visitedNodeList, root.value)
		if root.left != nil {
			traverse(root.left)
		}
		if root.right != nil {
			traverse(root.right)
		}
	}
	traverse(root)
	return
}

//DFSPostOrder DFSPostOrder
func (bst *BinarySearchTree) DFSPostOrder() (visitedNodeList []interface{}) {
	if bst.root == nil {
		return
	}
	root := bst.root
	var traverse func(root *Node)
	traverse = func(root *Node) {
		if root.left != nil {
			traverse(root.left)
		}
		if root.right != nil {
			traverse(root.right)
		}
		visitedNodeList = append(visitedNodeList, root.value)
	}
	traverse(root)
	return
}

//DFSInOrder DFSInOrder
func (bst *BinarySearchTree) DFSInOrder() (visitedNodeList []interface{}) {
	if bst.root == nil {
		return
	}
	root := bst.root
	var traverse func(root *Node)
	traverse = func(root *Node) {
		if root.left != nil {
			traverse(root.left)
		}
		visitedNodeList = append(visitedNodeList, root.value)
		if root.right != nil {
			traverse(root.right)
		}
	}
	traverse(root)
	return
}
