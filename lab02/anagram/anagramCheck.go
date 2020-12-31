package anagram

// IsAnagram checks is string a and b is anagram.
func IsAnagram(strA, strB string) bool {
	if len(strA) != len(strB) {
		return false
	}
	strAElemMap := map[string]int{}
	strBElemMap := map[string]int{}

	for _, v := range strA {
		if _, ok := strAElemMap[string(v)]; ok {
			strAElemMap[string(v)] = strAElemMap[string(v)] + 1
		} else {
			strAElemMap[string(v)] = 1
		}
	}
	for _, v := range strB {
		if _, ok := strBElemMap[string(v)]; ok {
			strBElemMap[string(v)] = strBElemMap[string(v)] + 1
		} else {
			strBElemMap[string(v)] = 1
		}
	}

	for strAElem, strAElemFreq := range strAElemMap {
		if strAElemFreq != strBElemMap[strAElem] {
			return false
		}
	}
	return true
}
