// Solution to problem 3223 from leetcode
package leetcodethreetwotwothree

func minimumLength(s string) int {
	n := 'z' - 'a' + 1
	letters := make([]int, n)

	for _, ch := range s {
		letters[ch-'a']++
	}

	minimumLength := len(s)

	for _, count := range letters {

		if count < 3 {
			continue
		}

		if count%2 == 0 {
			minimumLength -= (count - 2)
		} else {
			minimumLength -= (count - 1)
		}

	}

	return minimumLength
}
