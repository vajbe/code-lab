/*
Problem:
Given an array of string parts, each part representing a consecutive fragment of the same original message,
rebuild the complete message based on the following rules:
- The message always starts with 'A' and ends with 'Z'.
- Two parts can be matched when the last character of the current part equals the first character of the next part.
- When combining parts, do not duplicate the overlapping character (e.g., "Ab" + "bcZ" = "AbcZ").
- The first characters among all parts are unique.
- It is guaranteed that a valid and unique solution exists.

Approach:
1. Store each part in a map keyed by its first character for O(1) lookup.
2. Start from the part beginning with 'A'.
3. Repeatedly look up the next part by matching the last character of the current message with the first character of the next part.
4. Append the next part without its first (overlap) character.
5. Continue until the last character of the rebuilt message is 'Z'.

Time Complexity:
O(N * L)
- N = number of parts
- L = average length of each part
We process each part exactly once, and appending strings takes proportional time to their length.

Space Complexity:
O(N * L) for storing parts in the map and building the result.

Category:
String manipulation, Hash map, Greedy construction.

Similar Questions:
- LeetCode 3217: Restore the Array from Adjacent Pairs (graph traversal)
- LeetCode 332: Reconstruct Itinerary (path reconstruction)
- String reconstruction problems where overlap matching is required
*/

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
