// Solution to problem 1752 from leetcode
package leetcodeonesevenfivetwo

func check(nums []int) bool {
	n := len(nums)
	start := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			start = i
		}
	}

	last := nums[n-1]
	previous := 0

	for i := 0; i < start; i++ {
		if nums[i] < last || previous > nums[i] {
			return false
		}
		previous = nums[i]
	}

	return true
}
