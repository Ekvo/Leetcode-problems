// Given a sentence that consists of some words separated by a single space, and a searchWord,
// check if searchWord is a prefix of any word in sentence.
//
// Return the index of the word in sentence (1-indexed) where searchWord is a prefix of this word.
// If searchWord is a prefix of more than one word, return the index of the first word (minimum index).
// If there is no such word return -1.
//
// A prefix of a string s is any leading contiguous substring of s.
package leetcodeonefourfivefive

import "strings"

// Compare characters
func isPrefixOfWord(sentence string, searchWord string) int {
	var indexWord = 0
link:
	for i := 0; i+len(searchWord) <= len(sentence); i++ {
		if i == 0 || sentence[i-1] == ' ' {
			indexWord++

			for j := 0; j < len(searchWord); j++ {
				if sentence[i] != searchWord[j] {

					for ; i < len(sentence) && sentence[i] != ' '; i++ {
					}
					continue link
				}
				i++
			}
			return indexWord
		}
	}
	return -1
}

// Strings
func isPrefixOfWord(sentence string, searchWord string) int {
	var words = strings.Split(sentence, " ")
	for index, word := range words {
		if strings.HasPrefix(word, searchWord) {
			return index + 1
		}
	}
	return -1
}
