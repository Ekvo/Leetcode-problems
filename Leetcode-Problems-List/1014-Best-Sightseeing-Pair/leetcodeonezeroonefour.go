// Solution to problem 1014 from leetcode
package leetcodeonezeroonefour

func maxScoreSightseeingPair(values []int) int {
	maxScore := 0
	maxLeftScore := values[0]

	for i := 1; i < len(values); i++ {
		curRightScore := values[i] - i
		maxScore = max(maxScore, maxLeftScore+curRightScore)

		curLeftScore := values[i] + i
		maxLeftScore = max(maxLeftScore, curLeftScore)
	}

	return maxScore
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
