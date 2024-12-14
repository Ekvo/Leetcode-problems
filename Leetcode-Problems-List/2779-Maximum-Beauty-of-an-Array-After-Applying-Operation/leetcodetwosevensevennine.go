// You are given a 0-indexed array nums and a non-negative integer k.
//
// In one operation, you can do the following:
//
// Choose an index i that hasn't been chosen before from the range [0, nums.length - 1].
// Replace nums[i] with any integer from the range [nums[i] - k, nums[i] + k].
//
// The beauty of the array is the length of the longest subsequence consisting of equal elements.
//
// Return the maximum possible beauty of the array nums after applying the operation any number of times.
//
// Note that you can apply the operation to each index only once.
//
// A subsequence of an array is a new array generated from the original array by deleting
// some elements (possibly none) without changing the order of the remaining elements.
package leetcodetwosevensevennine

import "sort"

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)
	var (
		start   = 0
		lenght  = 1
		doubleK = k + k
	)
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[start] > doubleK {
			lenght = max(lenght, i-start)
			start++
		}
	}
	lenght = max(lenght, len(nums)-start)

	return lenght
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
