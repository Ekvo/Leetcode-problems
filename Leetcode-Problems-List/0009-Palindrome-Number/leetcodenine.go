// Given an integer x, return true if x is a palindrome, and false otherwise.
package leetcodenine

import "strconv"

// Two Pointers
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var strInt = strconv.Itoa(x)

	for h, t := 0, len(strInt)-1; h < t; h, t = h+1, t-1 {
		if strInt[h] != strInt[t] {
			return false
		}
	}
	return true
}

// Digit
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var (
		digit  = 0
		number = 0
		source = x
	)
	for x > 0 {
		digit = x % 10
		x /= 10
		number = number*10 + digit
	}
	return number == source
}
