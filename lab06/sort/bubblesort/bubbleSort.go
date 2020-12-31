package bubblesort

//Sort int numbers in ascenging order
func Sort(list []int) []int {
	if len(list) == 0 || len(list) == 1 {
		return list
	}
	n := len(list) - 1
	for n > 0 {
		for i := 0; i < n; i++ {
			if list[i] > list[i+1] {
				temp := list[i]
				list[i] = list[i+1]
				list[i+1] = temp
			}
		}
		n--
	}
	return list
}
