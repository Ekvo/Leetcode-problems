// Solution to problem 2116 from leetcode
package leetcodetwooneonesix

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}

	n := len(s)
	openCount := 0

	for i := 0; i < n; i++ {

		if locked[i] == '0' || s[i] == '(' {
			openCount++
		} else {
			openCount--
		}

		if openCount < 0 {
			return false
		}
	}

	openCount = 0

	for i := n - 1; i > -1; i-- {

		if locked[i] == '0' || s[i] == ')' {
			openCount++
		} else {
			openCount--
		}

		if openCount < 0 {
			return false
		}
	}

	return true
}
