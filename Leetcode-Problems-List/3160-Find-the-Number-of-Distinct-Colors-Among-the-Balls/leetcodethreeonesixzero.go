// Solution to problem 3160 from leetcode
package leetcodethreeonesixzero

func queryResults(limit int, queries [][]int) []int {
	n := len(queries)
	ballColor := make(map[int]int, n)
	colorCount := make(map[int]int)

	result := make([]int, 0, n)

	for _, q := range queries {
		b := q[0]
		c := q[1]

		if color, ex := ballColor[b]; ex {
			if color != c {
				colorCount[color]--
				if colorCount[color] == 0 {
					delete(colorCount, color)
				}

				ballColor[b] = c
				colorCount[c]++
			}
		} else {
			ballColor[b] = c
			colorCount[c]++
		}

		result = append(result, len(colorCount))
	}

	return result
}
