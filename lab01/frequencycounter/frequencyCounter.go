package frequencycounter

import (
	"math"
)

//Same check frequency of elements
func Same(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	m1 := map[int]int{}
	m2 := map[int]int{}
	for _, v := range arr1 {
		if _, ok := m1[v]; !ok {
			m1[v] = 1
		} else {
			m1[v] = m1[v] + 1
		}
	}
	for _, v := range arr2 {
		if _, ok := m2[v]; !ok {
			m2[v] = 1
		} else {
			m2[v] = m2[v] + 1
		}
	}

	// fmt.Printf("a1:%#v", m1)
	// fmt.Printf("a2:%#v", m2)

	for elem, elemFreq := range m1 {
		// sqrElemt, okt := m2[int(math.Pow(float64(elem), 2.0))]
		// fmt.Println(sqrElemt, okt)
		// fmt.Println(m2[int(math.Pow(float64(elem), 2.0))])
		if sqrElemFreq, ok := m2[int(math.Pow(float64(elem), 2.0))]; ok {
			// fmt.Println("\n", sqrElem, ok)
			// fmt.Println("\n", elem, elemFreq, sqrElem, m2[sqrElem], int(math.Pow(float64(elem), 2.0)))
			if sqrElemFreq != elemFreq {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
