// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
package validPalindrome

import "unicode"

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	var (
		head   = 0
		tail   = len(s) - 1
		chHead rune
		chTail rune
	)
	for head < tail {

		chHead = rune(s[head])
		chTail = rune(s[tail])

		if !unicode.IsDigit(chHead) && !unicode.IsLetter(chHead) {
			head++
			continue
		}
		if !unicode.IsDigit(chTail) && !unicode.IsLetter(chTail) {
			tail--
			continue
		}

		if 65 <= chHead && chHead <= 90 {
			chHead += 32
		}
		if 65 <= chTail && chTail <= 90 {
			chTail += 32
		}

		if chHead != chTail {
			return false
		}
		head++
		tail--
	}
	return true
}
