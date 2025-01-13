// Solution to problem 32 from leetcode
package leetcodetheetwo

func longestValidParentheses(s string) int {
	n := len(s)

	if n < 2 {
		return 0
	}

	open := 0
	count := 0
	validParentheses := 0

	for i := 0; i < n; i++ {

		if s[i] == '(' {
			open++
			continue
		}

		if open > 0 {
			count += 2
			open--

			if open == 0 {
				validParentheses = max(validParentheses, count)
			}
			continue
		}

		count = 0
	}

	open = 0
	count = 0

	for i := n - 1; i > -1; i-- {

		if s[i] == ')' {
			open++
			continue
		}

		if open > 0 {
			count += 2
			open--

			if open == 0 {
				validParentheses = max(validParentheses, count)
			}
			continue
		}

		count = 0
	}

	return validParentheses
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
