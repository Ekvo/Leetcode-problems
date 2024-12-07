// You are given an integer array banned and two integers n and maxSum.
// You are choosing some number of integers following the below rules:
//
// The chosen integers have to be in the range [1, n].
// Each integer can be chosen at most once.
// The chosen integers should not be in the array banned.
// The sum of the chosen integers should not exceed maxSum.
//
// Return the maximum number of integers you can choose following the mentioned rules.
package leetcodetwofivefivefour

import (
	"math"
	"sort"
)

func maxCount(banned []int, n int, maxSum int) int {
	sort.Ints(banned)

	nForMaxSum := lengthBySum(maxSum)
	newN := false

	if n < nForMaxSum {
		maxSum = arithmeticProgressionSum(1, n)
	} else {
		n = nForMaxSum
		newN = true
	}
	currentSum := arithmeticProgressionSum(1, n)
	count := lengthByMembers(1, n)
	badNum := make(map[int]bool, count)

	for i := 0; i < len(banned); i++ {
		if banned[i] > n {
			break
		}
		if !badNum[banned[i]] {
			badNum[banned[i]] = true
			currentSum -= banned[i]
			// there is no way to determine the exact (integer) number of elements in an arithmetic progression
			// example: sum = 20 -> 1 + 2 + 3 + 4 + 5 = 15 but 15 - 2 + 6 = 19
			// need compare with maxSum
			if newN && currentSum+n+1 <= maxSum {
				n++
				currentSum += n
			} else {
				count--
			}
		}
	}
	return count
}

func arithmeticProgressionSum(a1, an int) int {
	return int(float64(a1+an) / 2.0 * float64(lengthByMembers(a1, an)))
}

func lengthByMembers(a1, an int) int {
	return an - a1 + 1
}

// Sn = 2*a1 + d(n - 1)/2*n
// d = 1; a1 = 1
// n^2 + n - 2*Sn = 0
// n:=(-1(+/-)sqrt(1^2 - 4 * Sn * 2))/2
func lengthBySum(arProgSum int) int {
	n := (-1 + math.Sqrt(float64(1+4*arProgSum*2))) / 2.0
	return int(n)
}
