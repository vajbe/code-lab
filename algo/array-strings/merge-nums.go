// https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
package main

import "fmt"

func MergeSorted() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}
	MergeSortedArray(a, 3, b, 3)
	fmt.Println("Merged array:", a)

}

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
