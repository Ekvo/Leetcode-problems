// Solution to problem 1930 from leetcode
package leetcodeoneninethreezero

func countPalindromicSubsequence(s string) int {
	n := 'z' - 'a' + 1
	first := make([]int, n)
	last := make([]int, n)

	sourceVal := -1
	//mark all indexes as less than zero
	arrWithSourceVal(first, sourceVal)

	for i := 0; i < len(s); i++ {
		cur := int(s[i] - 'a')

		if first[cur] == sourceVal {
			first[cur] = i
		}
		last[cur] = i
	}

	numberOfPalindromes := 0

	for i := 0; i < len(first); i++ {

		if begin := first[i]; begin != sourceVal {
			visitLetter := make([]bool, n)
			end := last[i]

			for begin++; begin < end; begin++ {
				cur := int(s[begin] - 'a')

				if !visitLetter[cur] {
					visitLetter[cur] = true
					numberOfPalindromes++
				}
			}
		}
	}

	return numberOfPalindromes
}

func arrWithSourceVal(arr []int, sourceVal int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] = sourceVal
	}

	return arr
}
