package selectionsort

//Sort SelectionSort
func Sort(list []int) []int {
	if len(list) == 0 {
		return list
	}

	for i := 0; i < len(list); i++ {
		min := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[min] {
				min = j
			}
		}
		temp := list[min]
		list[min] = list[i]
		list[i] = temp
	}
	return list
}
