// Solution to problem 1790 from leetcode
package leetcodeonesevenninezero

func areAlmostEqual(s1 string, s2 string) bool {
	n := len(s1)

	countDif := 0
	ch11 := byte('+')
	ch12 := byte('-')
	ch21 := byte('*')
	ch22 := byte('/')

	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			countDif++

			if countDif > 2 {
				return false
			} else if countDif == 1 {
				ch11 = s1[i]
				ch21 = s2[i]
			} else {
				ch12 = s1[i]
				ch22 = s2[i]
			}
		}
	}

	if countDif != 0 && (ch11 != ch22 || ch21 != ch12) {
		return false
	}

	return true
}
