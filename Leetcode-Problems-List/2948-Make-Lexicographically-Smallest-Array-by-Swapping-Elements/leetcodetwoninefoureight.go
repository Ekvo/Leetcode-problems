// Solution to problem 2948 from leetcode
package leetcodetwoninefoureight

import (
	"slices"
	"sort"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)

	numsAsc := slices.Clone(nums)
	sort.Ints(numsAsc)

	// arrays store numbers according to the rule (numsAsc[i] - numsAsc[i-1] <= limit)
	arrays := make([][]int, 1, n)
	arrays[0] = []int{numsAsc[0]}

	indexArr := 0
	numInArr := make(map[int]int, n) // k - number; value - index of Array

	// load arrays
	for i := 1; i < n; i++ {
		cur := numsAsc[i]

		// create new array
		if cur-numsAsc[i-1] > limit {
			arrays = append(arrays, []int{})
			indexArr++
		}

		arrays[indexArr] = append(arrays[indexArr], cur)
		numInArr[cur] = indexArr
	}

	// find number of sorting array and take fist num
	// then decrement the array by one
	for i := 0; i < n; i++ {
		numArr := numInArr[nums[i]]

		nums[i] = arrays[numArr][0]
		arrays[numArr] = arrays[numArr][1:]
	}

	return nums
}
