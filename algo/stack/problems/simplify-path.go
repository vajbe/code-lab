/*  simplifyPath converts a given Unix-style file path into its canonical form.

Problem Description:
--------------------
Given a string path representing a Unix-style absolute path, simplify it.
The path may contain ".", "..", and multiple consecutive slashes.
The canonical path must:
  1. Start with a single '/'.
  2. Have no trailing '/' (except when it is just root "/").
  3. Ignore '.' (current directory).
  4. Correctly resolve '..' (move one directory up if possible).
  5. Collapse multiple slashes '//' into a single '/'.

Example 1:
  Input:  "/home/user/Documents/../Pictures"
  Output: "/home/user/Pictures"

Example 2:
  Input:  "/a/./b/../../c/"
  Output: "/c"

Example 3:
  Input:  "/../"
  Output: "/"

Approach:
---------
Use a stack to process directory tokens:
  1. Split path into tokens separated by '/' while iterating.
  2. If token is ".." -> pop from stack if not empty.
  3. If token is "." or "" -> ignore.
  4. Otherwise, push token into stack.
  5. Join the stack contents with "/" and prepend "/".

Pseudo-code:
------------
function simplifyPath(path):
    stack = []
    curr = ""
    for char in path + "/":
        if char == "/":
            if curr == "..":
                if stack not empty: pop stack
            else if curr != "" and curr != ".":
                push curr into stack
            curr = ""    //reset token
        else:
            curr += char
    return "/" + join(stack, "/")

Complexity Analysis:
--------------------
Time Complexity:  O(N), where N is length of the path string.
Space Complexity: O(N), for storing tokens in the stack.

Similar Problems:
-----------------
- LeetCode 71: Simplify Path
- Evaluate Reverse Polish Notation (stack usage)
- Valid Parentheses (nested structure with stack)
- Remove K Digits (string simplification with stack)
- Minimum Remove to Make Parentheses Valid */

package problems

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {

	stack := []string{}
	curr := ""

	for _, char := range path + "/" {

		str := string(char)

		if str == "/" {
			if curr == ".." {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else if curr != "" && curr != "." {
				stack = append(stack, curr)

			}
			curr = ""

		} else {
			curr += str
		}

	}

	return "/" + strings.Join(stack, "/")
}

func SimplePath() {
	fmt.Println(simplifyPath("/home/user/Documents/../Pictures"))
}
