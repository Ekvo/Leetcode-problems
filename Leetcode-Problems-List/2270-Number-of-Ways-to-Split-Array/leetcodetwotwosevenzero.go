// Solution to problem 2270 from leetcode
package leetcodetowtwosevenzero

func waysToSplitArray(nums []int) int {
	n := len(nums) - 1
	index := n
	leftSum := 0
	rightSum := 0

	for index > 0 {
		rightSum += nums[index]
		index--
	}

	validSplits := 0

	for index < n {
		leftSum += nums[index]
		index++

		if leftSum >= rightSum {
			validSplits++
		}
		rightSum -= nums[index]
	}

	return validSplits
}
