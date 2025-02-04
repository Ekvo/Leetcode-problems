// Solution to problem 3105 from leetcode
package leetcodethreeonezerofive

func longestMonotonicSubarray(nums []int) int {
	ans := 1
	asc := 1
	desc := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			asc++
			desc = 1
		} else if nums[i] < nums[i-1] {
			desc++
			asc = 1
		} else {
			asc = 1
			desc = 1
			continue
		}
		ans = max(ans, max(asc, desc))
	}

	return ans
}
