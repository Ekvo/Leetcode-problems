// You have a bomb to defuse, and your time is running out! Your informer will provide you with a circular array code of length of n and a key k.
//
// To decrypt the code, you must replace every number. All the numbers are replaced simultaneously.
//
// If k > 0, replace the ith number with the sum of the next k numbers.
// If k < 0, replace the ith number with the sum of the previous k numbers.
// If k == 0, replace the ith number with 0.
//
// As code is circular, the next element of code[n-1] is code[0], and the previous element of code[0] is code[n-1].
//
// Given the circular array code and an integer key k, return the decrypted code to defuse the bomb!
package leetcodeonesixfivetwo

func decrypt(code []int, k int) []int {
	var result = make([]int, len(code))

	if k == 0 {
		return result
	}

	var positiveK = abs(k)

	for i := 0; i < len(code); i++ {

		for j, step := 0, i; j < positiveK; j++ {

			if k > 0 {
				step++
				if step == len(code) {
					step = 0
				}
			} else {
				step--
				if step == -1 {
					step = len(code) - 1
				}
			}
			result[i] += code[step]
		}
	}
	return result
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
