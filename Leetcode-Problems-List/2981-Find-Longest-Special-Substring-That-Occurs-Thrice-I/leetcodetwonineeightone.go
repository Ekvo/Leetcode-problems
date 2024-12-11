// You are given a string s that consists of lowercase English letters.
//
// A string is called special if it is made up of only a single character.
// For example, the string "abc" is not special, whereas the strings "ddd", "zz", and "f" are special.
//
// Return the length of the longest special substring of s which occurs at least thrice,
// or -1 if no special substring occurs at least thrice.
//
// A substring is a contiguous non-empty sequence of characters within a string.
package leetcodetwonineeightone

func maximumLength(s string) int {
	var (
		start     = 0
		n         = len(s) - 1
		wordCount = make(map[string]int)

		mapFill = func(end int) {
			count := 1
			for start <= end {
				wordCount[s[start:end+1]] += count
				start++
				count++
			}
		}
	)
	for i := 0; i < n; i++ {
		if s[i] != s[i+1] {
			mapFill(i)
		}
	}
	mapFill(n)
	var lenght = -1

	for w, c := range wordCount {
		if c > 2 && lenght < len(w) {
			lenght = len(w)
		}
	}
	return lenght
}
