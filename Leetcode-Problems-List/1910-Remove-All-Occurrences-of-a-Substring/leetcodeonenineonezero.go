// Solution to problem 1910 from leetcode
package leetcodeonenineonezero

func removeOccurrences(s string, part string) string {
	n := len(part)
	line := []byte(s)

	for i, first := 0, -1; i+n <= len(line); {

		if line[i] == part[0] {
			// mark first 'part[0]' on 'line'
			// for comeback
			if first == -1 {
				first = i
			}

			k, j := i+1, 1
			for j < n && line[k] == part[j] {
				k++
				j++
			}

			//if find 'part' on 'line'
			if j == n {
				line = append(line[:i], line[k:]...)

				if i == first {
					first = -1
					i = max(i-1, 0)
				} else {
					i = first
				}
				continue
			}
		}
		i++
	}

	return string(line)
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
