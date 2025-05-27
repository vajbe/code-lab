// https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
package main

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	i := 0
	j := 0
	result := []int{}
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	// a := []int{1, 2, 3, 0, 0, 0}
	// b := []int{2, 5, 6}
	/* 	result = append(result, nums1[i:]...)
	   	result = append(result, nums2[j:]...) */
	return result
}
