package main

import "fmt"

func main() {
	str := "helloworld"
	_ = str
	count := naiveSearch(str, "lo")
	fmt.Println("count:", count)
}

func naiveSearch(str, substr string) int {
	count := 0

	for k := range str {
		match := true
		for ks := range substr {
			fmt.Println(
				"str[k+ks]:", string(str[k+ks]),
				"substr[ks]:", string(substr[ks]),
				substr[ks] != str[k+ks], k, ks,
				"count:", count,
			)
			if substr[ks] != str[k+ks] {
				match = false
				break
			}
		}
		if match {
			count += 1
		}
	}
	return count
}
