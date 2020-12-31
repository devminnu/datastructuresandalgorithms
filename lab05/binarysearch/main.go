package main

import "fmt"

func main() {
	list := []int{2, 3, 6, 8, 11, 14, 19, 25, 26}
	index := binarySearch(list, 25)
	fmt.Println("index:", index)
}

func binarySearch(list []int, elem int) int {
	i := 0
	j := len(list) - 1
	for i < j || i == j {
		mid := (i + j) / 2
		fmt.Println("mid:", mid, list[mid])
		if list[mid] == elem {
			return mid
		}
		if elem < list[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}
