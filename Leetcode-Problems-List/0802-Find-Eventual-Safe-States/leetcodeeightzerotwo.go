// Solution to problem 802 from leetcode
package leetcodeeightzerotwo

func eventualSafeNodes(graph [][]int) []int {
	rows := len(graph)

	safeNodes := make([]bool, rows)
	visited := make([]bool, rows)

	var findSafeNode func(r int) bool

	findSafeNode = func(r int) bool {
		if safeNodes[r] {
			return true
		}

		if visited[r] {
			return false
		}
		visited[r] = true

		cols := len(graph[r])

		for c := 0; c < cols; c++ {
			if !findSafeNode(graph[r][c]) {
				return false
			}
		}

		safeNodes[r] = true

		return true
	}

	result := make([]int, 0, rows)

	for r := 0; r < rows; r++ {

		if safeNodes[r] || findSafeNode(r) {
			result = append(result, r)
		}
	}

	return result
}
