// You are given an integer array nums where the ith bag contains nums[i] balls.
// You are also given an integer maxOperations.
//
// You can perform the following operation at most maxOperations times:
//
// Take any bag of balls and divide it into two new bags with a positive number of balls.
// For example, a bag of 5 balls can become two new bags of 1 and 4 balls, or two new bags of 2 and 3 balls.
//
// Your penalty is the maximum number of balls in a bag. You want to minimize your penalty after the operations.
//
// Return the minimum possible penalty after performing the operations.
package leetcodeonesevensixzero

import "math"

func minimumSize(nums []int, maxOperations int) int {
	left := 1
	right := 0
	midle := (left + right) / 2

	for _, num := range nums {
		right = max(right, num)
	}

	isPosible := func(midle int) bool {
		totalOperations := 0
		for _, num := range nums {
			operations := int(math.Ceil(float64(num)/float64(midle)) - 1.0)
			totalOperations += operations

			if totalOperations > maxOperations {
				return false
			}
		}
		return true
	}

	for left < right {
		midle = (left + right) / 2

		if isPosible(midle) {
			right = midle
		} else {
			left = midle + 1
		}
	}
	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
