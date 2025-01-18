// Solution to problem 2683 from leetcode
package leetcodetwosixeightthree

func doesValidArrayExist(derived []int) bool {
	xor := 0 // a ^ 0 = a

	// original [a,b,c] -> derived [(a^b),(b^c),(c^a)] -> (a^b)^(b^c)^(c^a) = 0 -> a ^ b ^ c = 0
	for _, d := range derived {
		xor ^= d
	}

	return xor == 0
}
