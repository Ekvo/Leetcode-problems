// There is an undirected tree with n nodes labeled from 0 to n - 1.
// You are given the integer n and a 2D integer array edges of length n - 1,
// where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.
//
// You are also given a 0-indexed integer array values of length n, where values[i]
// is the value associated with the ith node, and an integer k.
//
// A valid split of the tree is obtained by removing any set of edges, possibly empty,
// from the tree such that the resulting components all have values that are divisible by k,
// where the value of a connected component is the sum of the values of its nodes.
//
// Return the maximum number of components in any valid split.
package leetcodetwoeightseventwo

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	var (
		result             = 0
		nodePath           = make([][]int, n)
		componentSumFinder func(cur, source int) int
	)
	componentSumFinder = func(cur, source int) int {
		sum := 0
		for _, path := range nodePath[cur] {
			if path != source {
				sum += componentSumFinder(path, cur)
				sum %= k // Ensure the sum stays within bounds
			}
		}
		sum += values[cur]
		sum %= k
		if sum == 0 {
			result++
		}
		return sum
	}
	for _, node := range edges {
		nodePath[node[0]] = append(nodePath[node[0]], node[1])
		nodePath[node[1]] = append(nodePath[node[1]], node[0])
	}
	componentSumFinder(0, -1)
	return result
}
