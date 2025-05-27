package main

import "fmt"

func FirstDuplicate() {
	data := []int{1, 2, 3, 5, 2, 3, 4, 5}
	hashMap := make(map[int]bool)
	first := -1

	for _, val := range data {
		if !hashMap[val] {
			hashMap[val] = true
		} else {
			first = val
			break
		}
	}

	fmt.Print("First duplicate is ", first)

}
