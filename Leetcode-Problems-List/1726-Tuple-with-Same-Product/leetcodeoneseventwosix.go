// Solution to problem 1726 from leetcode
package leetcodeoneseventwosix

func tupleSameProduct(nums []int) int {
	n := len(nums)
	multiplyTwoNums := make(map[int]int, n/2)

	for i := n - 1; i > 0; i-- {
		a := nums[i]

		for j := i - 1; j > -1; j-- {

			multiplyTwoNums[a*nums[j]]++
		}
	}

	result := 0

	for _, count := range multiplyTwoNums {
		if count > 1 {
			// Triangular Numbers: n*(n+1)/2
			// if count == 3: 3 - 1 = n -> (3 - 1) * 3 / 2 * 8
			result += (count - 1) * count * 4
		}
	}

	return result
}
