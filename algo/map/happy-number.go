// Problem: Happy Number
//
// Description:
// A happy number is defined by the following process:
// - Start with any positive integer.
// - Replace the number with the sum of the squares of its digits.
// - Repeat the process until the number equals 1 (where it will stay),
//   or it loops endlessly in a cycle which does not include 1.
// - Numbers that eventually reach 1 are "happy"; others are not.
//
// Approach:
// - Use a helper function getNext(n) to compute the sum of squares of digits.
// - Track visited numbers using a map (or set) to detect cycles.
// - If we encounter 1 → return true (happy number).
// - If we encounter a number already seen → return false (cycle detected).
// - Alternative approach: Floyd’s cycle detection (Tortoise & Hare) can be used
//   to optimize space to O(1).
//
// Time Complexity:
// - Each step computes sum of squares of digits → O(log n) digits.
// - Numbers shrink quickly and converge to <= 243 (since max 9² * digits for int range).
// - Overall: O(log n) per step, and bounded iterations → effectively O(1) in practice.
//
// Space Complexity:
// - O(k) for visited set (where k is number of unique numbers before cycle/1).
// - Optimized version: O(1) using cycle detection.
//
// Pseudocode:
//   function isHappy(n):
//       seen = empty set
//       while n != 1 and n not in seen:
//           add n to seen
//           n = sum of squares of digits of n
//       return n == 1
//
// Similar Problems:
// - Linked List Cycle Detection (Floyd’s Tortoise and Hare)
// - Detect cycle in a functional graph
// - Digital Root / Sum of Digits problems
// - Collatz Conjecture style problems (iterative transformation + cycle detection)

package main

import "fmt"

func isHappy(n int) bool {
	seen := make(map[int]bool)
	getNext := func(num int) int {
		total := 0

		for num > 0 {
			digit := num % 10
			total += digit * digit
			num /= 10
		}
		return total
	}

	for !seen[n] && n != 1 {
		seen[n] = true
		n = getNext(n)
	}

	return n == 1
}

func HappyNum() {
	fmt.Println("Is happy num ", isHappy(19))
}
