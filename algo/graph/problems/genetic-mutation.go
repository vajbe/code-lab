/*
	Minimum Genetic Mutation

Problem Description

You are given two gene strings start and end (each of length 8) and a list of valid gene strings called bank.
A gene string can only consist of the characters 'A', 'C', 'G', 'T'.

A mutation is defined as changing a single character in the gene string. Each intermediate mutation must exist in the bank.

Task:
Return the minimum number of mutations needed to mutate from start to end. If no valid mutation path exists, return -1.

# Approach

We can model this as a shortest path problem in an unweighted graph:

Each gene string is a node.

An edge exists between two nodes if their strings differ by exactly one character.

We want the shortest path from start → end.

The natural solution is Breadth-First Search (BFS) because:

BFS explores level by level.

The first time we encounter end, we are guaranteed it is the shortest number of mutations.

Steps:

Store all strings from bank in a hash set for O(1) lookup.

Use BFS starting from start.

At each step, try mutating one character at a time to 'A', 'C', 'G', 'T'.

If the new string is in the bank and not visited, push it into the queue with steps+1.

If we reach end, return the current step count.

If BFS finishes without finding end, return -1.

Pseudocode
function minMutation(start, end, bank):

	dict = set(bank)
	if end not in dict:
	    return -1

	queue = [(start, 0)]
	visited = {start}
	choices = ['A', 'C', 'G', 'T']

	while queue not empty:
	    gene, steps = queue.pop_front()
	    if gene == end:
	        return steps

	    for i in 0..7:
	        for ch in choices:
	            if ch == gene[i]:
	                continue
	            new_gene = gene with gene[i] replaced by ch
	            if new_gene in dict and new_gene not visited:
	                visited.add(new_gene)
	                queue.push((new_gene, steps+1))

	return -1

# Time Complexity

Each gene string has length L = 8.

For each character, we can try 3 new mutations (4-1).

So, at most O(L * 3) mutations per gene.

Each valid gene from the bank is visited once.

Time: O(B * L * 3) → where B = len(bank)
Space: O(B) for visited + dictionary.

# Similar Problems

Word Ladder (LeetCode 127) → Almost identical, words instead of genes.

Word Ladder II (LeetCode 126) → Find all shortest mutation sequences.

Rotting Oranges (LeetCode 994) → BFS spreading step by step.

Shortest Path in Binary Matrix (LeetCode 1091) → BFS on a grid.
*/
package problems

func minMutation(startGene string, endGene string, bank []string) int {
	bankMap := make(map[string]bool)

	for _, value := range bank {
		bankMap[value] = true
	}

	if !bankMap[endGene] {
		return -1
	}

	type item struct {
		gene  string
		steps int
	}

	visited := map[string]bool{startGene: true}
	queue := []item{{startGene, 0}}
	choices := []byte{'A', 'C', 'G', 'T'}

	for len(queue) > 0 {

		el := queue[0]
		queue = queue[1:]

		if el.gene == endGene {
			return el.steps
		}

		bs := []byte(el.gene)
		for i := 0; i < len(bs); i++ {
			original := bs[i]

			for _, val := range choices {
				if val == original {
					continue
				}
				bs[i] = val
				nxt := string(bs)
				if !visited[nxt] && bankMap[nxt] {
					visited[nxt] = true
					queue = append(queue, item{nxt, el.steps + 1})
				}
			}
			bs[i] = original
		}

	}

	return -1
}

func MinimumMut() {

}
