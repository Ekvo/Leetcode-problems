// You are given an array nums of non-negative integers and an integer k.
//
// An array is called special if the bitwise OR of all of its elements is at least k.
//
// Return the length of the shortest special non-empty
// subarray
// of nums, or return -1 if no special subarray exists.
package leetcodethreezeronineseven

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	var (
		result       = math.MaxInt64
		current      = 0
		countToRight = 0
	)
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			countToRight++
			current |= nums[i]

			if k <= current {

				if result > countToRight {
					result = countToRight
				}
				var (
					j           = i
					countToLeft = 1
				)
				current = nums[i]

				for j--; j > -1 && countToLeft < countToRight-1; j-- {

					countToLeft++
					current |= nums[j]

					if k <= current {

						if result > countToLeft {
							result = countToLeft
						}
						break
					}
				}
				j++
				i = j
				current = nums[i]
				countToRight = 1
			}
		} else {
			//find subarray of length 1
			result = 1
			break
		}
	}
	if result != math.MaxInt64 {
		return result
	}
	return -1
}
