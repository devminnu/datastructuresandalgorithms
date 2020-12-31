package multipointer

//SumZeroPair get the pair of whose sum is zero in a given sorted array
func SumZeroPair(list []int) [][]int {
	sumZeroPairList := [][]int{}
	i, j := 0, len(list)-1

	for i < j {
		if list[i]+list[j] == 0 {
			sumZeroPairList = append(sumZeroPairList, []int{list[i], list[j]})
			i++
		}
		if list[i]+list[j] > 0 {
			j--
			continue
		}
		i++
	}
	return sumZeroPairList
}
