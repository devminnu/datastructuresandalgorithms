package multipointer

import (
	"fmt"
	"sort"
)

//CountUniqueValues get the values from an array
func CountUniqueValues(list []int) ([]int, int) {
	if len(list) == 1 {
		return list, 1
	}
	sort.Ints(list)
	fmt.Println(list)
	i, j := 0, 1
	for j < len(list) {
		if list[i] == list[j] {
			j++
			continue
		}
		i++
		list[i] = list[j]
	}
	return list[:i+1], i
}
