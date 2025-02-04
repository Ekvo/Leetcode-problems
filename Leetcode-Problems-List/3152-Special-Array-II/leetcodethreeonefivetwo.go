// Solution to problem 3152 from leetcode
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
