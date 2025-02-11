// Solution to problem 3174 from leetcode
package leetcodethreeonesevenfour

func clearDigits(s string) string {
	line := []byte(s)
	end := 0

	// if need Space O(1)
	// line := unsafe.Slice(unsafe.StringData(s), len(s))

	for i, ch := range line {
		if '0' <= ch && ch <= '9' {
			end--
			if end < 0 {
				end = 0
			}
		} else {
			line[end] = line[i]
			end++
		}
	}

	// if need Space O(1)
	// return unsafe.String(unsafe.SliceData(line), end)

	return string(line[:end])
}
