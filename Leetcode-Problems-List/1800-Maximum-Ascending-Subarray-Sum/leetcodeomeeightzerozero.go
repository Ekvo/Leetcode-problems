// Solution to problem 1800 from leetcode
package leetcodeoneeightzerozero

func maxAscendingSum(nums []int) int {
	ans := nums[0]
	curSum := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] > nums[i-1] {
			curSum += nums[i]
		} else {
			curSum = nums[i]
		}
		ans = max(ans, curSum)
	}

	return ans
}
