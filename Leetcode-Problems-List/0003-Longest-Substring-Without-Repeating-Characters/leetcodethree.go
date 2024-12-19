// Given a string s, find the length of the longest substring without repeating characters.
package leetcodethree

import "bytes"

// Two pointers
func lengthOfLongestSubstring(s string) int {
	var (
		simbols = make([]int, 1<<8)
		start   = 0
		lenght  = 0
	)
	for i := 0; i < len(s); i++ {
		simbols[s[i]]++

		for simbols[s[i]] > 1 {
			simbols[s[start]]--
			start++
		}
		lenght = max(lenght, i-start+1)
	}
	return lenght
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Brute force
func lengthOfLongestSubstring(s string) int {
	var (
		buff   bytes.Buffer
		lenght = 0
		count  = 0
	)
	buff.Grow(len(s))

	for i := 0; i < len(s); i++ {
		count = 0

		for j := i; j < len(s); j++ {

			if bytes.IndexByte(buff.Bytes(), s[j]) != -1 {
				buff.Reset()
				break
			} else {
				buff.WriteByte(s[j])
				count++
			}
		}
		if lenght < count {
			lenght = count
		}
	}
	return lenght
}
