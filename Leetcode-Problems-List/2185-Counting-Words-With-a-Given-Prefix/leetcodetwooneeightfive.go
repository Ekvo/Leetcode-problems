// Solution to problem 2185 from leetcode
package leetcodetwooneeightfive

func prefixCount(words []string, pref string) int {
	numberOfStrings := 0

	for _, word := range words {
		if len(word) < len(pref) {
			continue
		}

		if word[0:len(pref)] == pref {
			numberOfStrings++
		}
	}

	return numberOfStrings
}
