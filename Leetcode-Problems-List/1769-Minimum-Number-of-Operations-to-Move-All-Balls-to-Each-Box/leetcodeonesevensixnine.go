// Solution to problem 1769 from leetcode
package leetcodeonesevensixnine

// Sum of Left and Right Moves
/*
In each iteration, we update the answer array by adding the moves calculated from both the left and right sides.
The value for each box in answer[h] (for the left pass) and answer[t] (for the right pass)
represents the total moves required for balls to reach that box.
*/
func minOperations(boxes string) []int {
	n := len(boxes)
	answer := make([]int, n)

	ballsLeft, ballsRight := 0, 0
	moveLeft, moveRight := 0, 0

	for h, t := 0, n-1; h < n; h, t = h+1, t-1 {

		answer[h] += moveLeft
		ballsLeft += int(boxes[h] - '0')
		moveLeft += ballsLeft

		answer[t] += moveRight
		ballsRight += int(boxes[t] - '0')
		moveRight += ballsRight
	}

	return answer
}

// Brute Force
func minOperations(boxes string) []int {
	oneIndexs := make([]int, 0, len(boxes))

	for i, ch := range boxes {
		if ch == '1' {
			oneIndexs = append(oneIndexs, i)
		}
	}
	answer := make([]int, len(boxes))

	for i, _ := range boxes {

		for _, index := range oneIndexs {
			answer[i] += abs(i - index)
		}
	}

	return answer
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
