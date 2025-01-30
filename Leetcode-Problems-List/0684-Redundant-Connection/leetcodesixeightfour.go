// Solution to problem 684 from leetcode
package leetcodesixeightfour

func Leetcode684findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parents := make([]int, n+1)

	for i := 1; i <= n; i++ {
		parents[i] = i
	}

	var getParent func(b int) int

	getParent = func(b int) int {
		if parents[b] != b {
			parents[b] = getParent(parents[b])
		}

		return parents[b]
	}

	equalParents := func(a, b int) bool {
		pA, pB := getParent(a), getParent(b)
		if pA == pB {
			return true
		}

		if pA < pB {
			parents[pB] = pA
		} else {
			parents[pA] = pB
		}

		return false
	}

	var result []int = nil

	for _, edge := range edges {
		a, b := edge[0], edge[1]

		if equalParents(a, b) {
			result = edge
			break
		}
	}

	return result
}
