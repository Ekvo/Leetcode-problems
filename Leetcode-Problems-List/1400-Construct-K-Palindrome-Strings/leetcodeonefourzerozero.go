// Solution to problem 1400 from leetcode
package leetcodeonefourzerozero

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}
	if len(s) == k {
		return true
	}

	n := 'z' - 'a' + 1
	letters := make([]int, n)

	for _, ch := range s {
		letters[ch-'a']++
	}

	oddLetters := 0

	for _, count := range letters {

		if count%2 != 0 {
			oddLetters++
		}
	}

	return oddLetters <= k
}
