// You are given a 0-indexed integer array nums of length n.
//
// You can perform the following operation as many times as you want:
//
// Pick an index i that you haven’t picked before,
// and pick a prime p strictly less than nums[i], then subtract p from nums[i].
//
// Return true if you can make nums a strictly increasing array using the above operation and false otherwise.
//
// A strictly increasing array is an array whose each element is strictly greater than its preceding element.
package leetcodetwosixzeroone

import "math"

// primeSubOperation - is main function in 2601 problem
func primeSubOperation(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	var loop = len(nums)

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] <= nums[i-1] {
			break
		}
		loop--
	}

	var primeNums = primeNumbers(2, 1001)

	nums[0] = toLessValue(primeNums, nums[0], 0)

	for i := 1; i < loop; i++ {
		nums[i] = toLessValue(primeNums, nums[i], nums[i-1])
		if nums[i] <= nums[i-1] {
			return false
		}
	}

	return true

}

// toLessValue -  decreases current value with help pN(primeNums)

// Pick an index i that you haven’t picked before,
// and pick a prime p strictly less than nums[i],
// then subtract p from nums[i].
func toLessValue(pN []int, current, previos int) int {
	var (
		head = 0
		tail = len(pN) - 1
		mid  = (head + tail) / 2
		//lenght binary seach
		lenght = int(math.Log(float64(len(pN))))
	)

	for ; lenght > 0; lenght-- {
		if pN[mid] > current {
			tail = mid - 1
		} else if pN[mid] < current {
			head = mid + 1
		}
		mid = (head + tail) / 2
	}

	for i := tail; i > -1; i-- {
		if pN[i] < current && current-pN[i] > previos {
			current -= pN[i]
			break
		}
	}

	return current
}

// Create an array of prime numbers
func primeNumbers(min, max int) []int {
	if min < 2 {
		min = 2
	}
	if max < min {
		max = min
	}

	var (
		loop = max - min
		//+3 because we have [0],[1] and plus 1
		readNums = make([]bool, loop+3)
		//168 max cap with range [0 - 1001] for prime numbers array
		primeNums = make([]int, 0, 170)
	)

	for i := min; i <= max; i++ {
		if !readNums[i] {
			primeNums = append(primeNums, i)
			for j := i + i; j <= max; j += i {
				readNums[j] = true
			}
		}
	}

	return primeNums
}
