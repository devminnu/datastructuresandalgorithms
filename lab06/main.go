package main

import (
	"fmt"
	"lab06/sort/selectionsort"
)

func main() {
	list := []int{
		8, 18, 3, 5, 1, 12, 9, 21, 23, 7,
	}
	res := selectionsort.Sort(list)
	fmt.Println("List:", res)
}
