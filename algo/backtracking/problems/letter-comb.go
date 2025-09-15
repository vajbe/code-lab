/*
Problem:
---------
Given a string containing digits from 2 to 9 inclusive, return all possible letter
combinations that the number could represent. Return the answer in any order.

This is based on the mapping of digits to letters on a phone keypad:
2 -> "abc"
3 -> "def"
4 -> "ghi"
5 -> "jkl"
6 -> "mno"
7 -> "pqrs"
8 -> "tuv"
9 -> "wxyz"

Examples:
----------
Input: "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Input: "2"
Output: ["a","b","c"]

Input: ""
Output: []

---

Approach:
---------
We can solve this problem using two standard traversal strategies:

1. DFS (Depth-First Search) with Backtracking:
   - Start with an empty prefix.
   - At each step, take the current digit and expand all possible letters.
   - Recursively build combinations until all digits are processed.
   - Backtrack after each recursive call to explore the next possibility.
   - This is a top-down recursive approach.

2. BFS (Breadth-First Search) / Iterative Expansion:
   - Use a queue (list of prefixes).
   - For each digit, expand all current prefixes by appending the possible letters.
   - Replace the queue with the newly generated set of prefixes.
   - This is a bottom-up iterative approach.

Both approaches generate all possible combinations.

---

Pseudo code (BFS version):
--------------------------
function letterCombinations(digits):
    if digits is empty:
        return []

    map = {2: ["a","b","c"], 3:["d","e","f"], ..., 9:["w","x","y","z"]}
    result = [""]
    for each digit in digits:
        next = []
        for prefix in result:
            for letter in map[digit]:
                next.append(prefix + letter)
        result = next
    return result

---

Complexity Analysis:
--------------------
Let n = length of digits
Let k = max number of letters for a digit (k = 4 for digits 7 and 9)

- Time Complexity: O(k^n * n)
  Each of the k^n combinations is of length n, and we generate all of them.

- Space Complexity: O(k^n * n)
  To store all the generated combinations in the result.
  Recursion depth for DFS is O(n).

---

Notes:
-------
- DFS is more natural for backtracking solutions.
- BFS is often more intuitive and avoids recursion overhead.
- Both are valid and optimal approaches since generating all combinations is required.
*/

package problems

import "fmt"

func dfs(char string, m map[string][]string, digits string) []string {
	// dfs with backtracking

	if len(digits) == 0 {
		return []string{char}
	}
	res := []string{}
	letters := m[string(digits[0])]
	for _, c := range letters {
		res = append(res, dfs(char+string(c), m, digits[1:])...)
	}
	return res
}

func bfs(m map[string][]string, digits string) []string {
	// bfs
	res := []string{""}

	for i := 0; i < len(digits); i++ {
		next := []string{}
		for _, prefix := range res { //queue
			for _, letter := range m[string(digits[i])] {
				next = append(next, prefix+letter)
			}
		}
		res = next
	}
	return res
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := make(map[string][]string)
	m["2"] = []string{"a", "b", "c"}
	m["3"] = []string{"d", "e", "f"}
	m["4"] = []string{"g", "h", "i"}
	m["5"] = []string{"j", "k", "l"}
	m["6"] = []string{"m", "n", "o"}
	m["7"] = []string{"p", "q", "r", "s"}
	m["8"] = []string{"t", "u", "v"}
	m["9"] = []string{"w", "x", "y", "z"}
	// res := dfs("", m, digits)
	res := bfs(m, digits)
	return res
}

func LetterComb() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("234"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}
