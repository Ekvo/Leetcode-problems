// Solution to problem 1408 from leetcode
package leetcodeonefourzeroeight

import (
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	substrings := make([]string, 0, len(words))

	for i := 0; i < len(words); i++ {

		for j := i + 1; j < len(words); j++ {
			if len(words[i]) == len(words[j]) {
				continue
			}

			if strings.Contains(words[j], words[i]) {
				substrings = append(substrings, words[i])
				break
			}
		}
	}

	return substrings
}
