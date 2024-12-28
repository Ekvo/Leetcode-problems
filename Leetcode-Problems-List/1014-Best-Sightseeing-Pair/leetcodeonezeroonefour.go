// You are given an integer array values where values[i] represents the value of the ith sightseeing spot.
// Two sightseeing spots i and j have a distance j - i between them.
//
// The score of a pair (i < j) of sightseeing spots is values[i] + values[j] + i - j:
// the sum of the values of the sightseeing spots, minus the distance between them.
//
// Return the maximum score of a pair of sightseeing spots.
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
