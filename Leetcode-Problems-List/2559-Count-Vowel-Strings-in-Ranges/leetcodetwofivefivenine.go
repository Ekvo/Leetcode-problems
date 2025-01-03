// Solution to problem 2559 from leetcode
package leetcodetwofivefivenine

func vowelStrings(words []string, queries [][]int) []int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	countVowelWords := make([]int, len(words))
	stepSum := 0

	for i, word := range words {

		if vowels[word[0]] && vowels[word[len(word)-1]] {
			stepSum++
		}
		countVowelWords[i] = stepSum
	}

	result := make([]int, len(queries))

	for i, querie := range queries {
		count := 0

		if querie[0] != 0 {
			count = countVowelWords[querie[1]] - countVowelWords[querie[0]-1]
		} else {
			count = countVowelWords[querie[1]]
		}
		result[i] = count
	}

	return result
}
