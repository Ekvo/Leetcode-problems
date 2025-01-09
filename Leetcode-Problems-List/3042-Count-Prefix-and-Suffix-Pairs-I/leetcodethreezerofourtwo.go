// Solution to problem 3042 from leetcode
package leetcodethreezerofourtwo

import "strings"

func countPrefixSuffixPairs(words []string) int {
	numberOfIndexs := 0

	for i := 0; i < len(words); i++ {

		for j := i + 1; j < len(words); j++ {

			if isPrefixAndSuffix(words[i], words[j]) {
				numberOfIndexs++
			}
		}
	}

	return numberOfIndexs
}

// With package strings
func isPrefixAndSuffix(prefSuf, word string) bool {
	return strings.HasPrefix(word, prefSuf) && strings.HasSuffix(word, prefSuf)
}

// With comparison of string objects
func isPrefixAndSuffix(prefSuf, word string) bool {
	if len(prefSuf) > len(word) {
		return false
	}

	return prefSuf == word[:len(prefSuf)] && word[len(word)-len(prefSuf):] == prefSuf
}

// With character comparison
func isPrefixAndSuffix(prefSuf, word string) bool {
	if len(prefSuf) > len(word) {
		return false
	}

	left := 0
	right := len(word) - 1

	// head & tail
	for h, t := 0, len(prefSuf)-1; h < len(prefSuf); h, t = h+1, t-1 {

		if word[left] != prefSuf[h] ||
			word[right] != prefSuf[t] {
			return false
		}
		left++
		right--
	}

	return true
}
