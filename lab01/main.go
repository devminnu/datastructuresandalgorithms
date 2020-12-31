package main

import (
	"fmt"
	"lab01/frequencycounter"
)

func main() {
	a1 := []int{1, 2, 3, 2}
	a2 := []int{4, 1, 4, 9}

	res := frequencycounter.Same(a1, a2)
	fmt.Println("result:", res)
}
