// Given an integer array arr, remove a subarray (can be empty) from arr such that
// the remaining elements in arr are non-decreasing.
//
// Return the length of the shortest subarray to remove.
//
// A subarray is a contiguous subsequence of the array.
package shortestSubarrayRemovedMakeArraySorted

func findLengthOfShortestSubarray(arr []int) int {
	var (
		head      = 0
		tail      = len(arr) - 1
		minSubArr = 0
	)

	for tail > 0 {
		if arr[tail-1] > arr[tail] {
			break
		}
		tail--
	}

	// array sorted
	if tail < 1 {
		return minSubArr
	}

	minSubArr = tail

	var contenderSubArr = 0

	for ; head < len(arr); head++ {

		//need compare first element with tail --> arr[0]
		if head > 0 && arr[head-1] > arr[head] {
			break
		}

		for tail < len(arr) {
			if arr[head] <= arr[tail] {
				break
			}
			tail++
		}
		contenderSubArr = tail - head - 1
		minSubArr = min(minSubArr, contenderSubArr)
	}
	return minSubArr
}
