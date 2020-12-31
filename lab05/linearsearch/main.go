package main

import "fmt"

func main() {
	list := []int{18, 2, 8, 12, 3, 4, 20, 10, 7, 11}
	index := linearSearch(list, 20)
	fmt.Println("index:", index)
}

func linearSearch(list []int, elem int) int {
	for k, v := range list {
		if v == elem {
			return k
		}
	}
	return -1
}
