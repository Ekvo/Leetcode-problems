// Solution to problem 2364 from leetcode
package leetcodethreesixfour

func countBadPairs(nums []int) int64 {
	n := len(nums)
	difNumIndex := make(map[int]int, n)

	for i := 0; i < n; i++ {
		difNumIndex[nums[i]-i]++
	}

	badPairs := n * (n - 1) / 2

	for _, count := range difNumIndex {
		if count > 1 {
			badPairs -= count * (count - 1) / 2
		}
	}

	return int64(badPairs)
}
