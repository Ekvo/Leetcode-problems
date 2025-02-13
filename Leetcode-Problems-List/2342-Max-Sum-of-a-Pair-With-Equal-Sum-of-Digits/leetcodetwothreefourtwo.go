// Solution to problem 2342 from leetcode
package leetcodetwothreefourtwo

func maximumSum(nums []int) int {
	summOfDigits := make([]int, 82)
	ans := -1

	for _, num := range nums {
		curSumDig := 0

		for n := num; n > 0; n /= 10 {
			curSumDig += n % 10
		}

		if summOfDigits[curSumDig] != 0 {
			ans = max(ans, summOfDigits[curSumDig]+num)
		}

		summOfDigits[curSumDig] = max(summOfDigits[curSumDig], num)
	}

	return ans
}
