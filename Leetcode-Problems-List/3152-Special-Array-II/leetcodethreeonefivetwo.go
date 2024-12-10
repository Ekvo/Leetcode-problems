// An array is considered special if every pair of its adjacent elements contains two numbers with different parity.
//
// You are given an array of integer nums and a 2D integer matrix queries, where for queries[i] = [fromi, toi] your task is to check that
// subarray
// nums[fromi..toi] is special or not.
//
// Return an array of booleans answer such that answer[i] is true if nums[fromi..toi] is special.
package leetcodethreeonefivetwo

func isArraySpecial(nums []int, queries [][]int) []bool {
	var (
		start    = 0
		n        = len(nums) - 1
		startEnd = make(map[int]int, len(nums))

		mapFill = func(end int) {
			for start <= end {
				startEnd[start] = end
				start++
			}
		}
	)
	startEnd[n] = n // if len(nums)=1
	for i := 0; i < n; i++ {
		if nums[i]%2 == nums[i+1]%2 {
			mapFill(i)
		} else if i == n-1 {
			mapFill(i + 1)
		}
	}
	var result = make([]bool, len(queries))

	for i := 0; i < len(queries); i++ {
		if num, ex := startEnd[queries[i][0]]; ex && queries[i][1] <= num {
			result[i] = true
		}
	}
	return result
}
