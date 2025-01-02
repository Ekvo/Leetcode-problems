// Solution to problem 1422 from leetcode
package leetcodeonefourtwotwo

func maxScore(s string) int {
	zero := 0
	one := 0
	n := len(s) - 1
	index := n

	// count '1' from right: (0 , len(s)-1]
	for index > 0 {
		if s[index] == '1' {
			one++
		}
		index--
	}

	result := 0

	for index < n {
		if s[index] == '0' {
			zero++
		}
		result = max(result, zero+one)
		index++

		if s[index] == '1' {
			one--
		}
	}

	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
