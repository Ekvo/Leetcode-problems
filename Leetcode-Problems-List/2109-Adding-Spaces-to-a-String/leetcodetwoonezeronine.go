// You are given a 0-indexed string s and a 0-indexed integer array spaces that describes the indices in the original string
// where spaces will be added. Each space should be inserted before the character at the given index.
//
// For example, given s = "EnjoyYourCoffee" and spaces = [5, 9],
// we place spaces before 'Y' and 'C', which are at indices 5 and 9 respectively.
// Thus, we obtain "Enjoy Your Coffee".
//
// Return the modified string after the spaces have been added.
package leetcodetwoonezeronine

import "bytes"

func addSpaces(s string, spaces []int) string {
	var (
		buff  bytes.Buffer
		start = 0
	)
	buff.Grow(len(s) + len(spaces))

	for i := 0; i < len(spaces); i++ {
		buff.WriteString(s[start:spaces[i]])
		buff.WriteByte(' ')
		start = spaces[i]
	}
	buff.WriteString(s[start:len(s)])

	return buff.String()
}
