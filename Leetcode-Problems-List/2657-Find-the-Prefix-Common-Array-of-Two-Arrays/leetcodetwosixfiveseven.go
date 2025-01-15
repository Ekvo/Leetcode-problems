// Solution to problem 2657 from leetcode
package leetcodetwosixfiveseven

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	prefixArr := make([]int, n)
	numbers := make([]int, n+1)
	count := 0

	for i := 0; i < n; i++ {

		if numbers[A[i]] < 1 {
			numbers[A[i]]++

		} else {
			count++
		}

		if numbers[B[i]] < 1 {
			numbers[B[i]]++

		} else {
			count++
		}

		prefixArr[i] = count
	}

	return prefixArr
}
