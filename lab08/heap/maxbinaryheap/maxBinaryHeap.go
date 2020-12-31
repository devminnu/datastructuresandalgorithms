package maxbinaryheap

import "fmt"

/*
Max Heap: All Nodes are smaller than parent node
Min Heap: All Nodes are greater than parent node
*/

//MaxBinaryHeap MaxBinaryHeap
type MaxBinaryHeap struct {
	values []int
}

//childIndex=2*ParentIndex+1 (leftChild)
//childIndex=2*ParentIndex+2 (rightChild)

//parentIndex=(childIndex-1)/2

// New create new MaxBinaryHeap
func New() *MaxBinaryHeap {
	return &MaxBinaryHeap{}
}

//Insert add val
func (mbh *MaxBinaryHeap) Insert(val int) {
	mbh.values = append(mbh.values, val)
	valindex := len(mbh.values) - 1
	parentIndex := (valindex - 1) / 2
	for mbh.values[parentIndex] < mbh.values[valindex] {
		temp := mbh.values[parentIndex]
		mbh.values[parentIndex] = mbh.values[valindex]
		mbh.values[valindex] = temp
		valindex = parentIndex
		parentIndex = (valindex - 1) / 2
	}
}

// PrintHeap PrintHeap
func (mbh *MaxBinaryHeap) PrintHeap() {
	fmt.Println(mbh.values)
}

/*
					    100
			   50				      90
		  40	  30		   80 	 20
  	35  10			   25

*/
