package main

import "fmt"

func FindDuplicate() {
	hashMap := make(map[int]bool)
	data := []int{1, 2, 3, 5, 1, 2, 3, 4, 5}

	for _, val := range data {
		if !hashMap[val] {
			hashMap[val] = true
		}
	}

	for k := range hashMap {
		fmt.Print(k, " ")
	}
}
