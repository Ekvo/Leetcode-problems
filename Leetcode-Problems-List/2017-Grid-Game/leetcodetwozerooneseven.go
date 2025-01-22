// Solution to problem 2017 from leetcode
package leetcodetwozerooneseven

func gridGame(grid [][]int) int64 {
	n := len(grid[0])

	if n < 2 {
		return 0
	}

	// Aliases for comfortable
	arr1 := grid[0]
	arr2 := grid[1]

	sumArr1 := 0

	for i := 0; i < n; i++ {
		sumArr1 += arr1[i]
	}

	result := 0

	// [2,5,4] -> sumArr1 = 11 - 2 -> 9 - 5 -> 4 - 4
	// [1,5,1] ->  result = 0  + 1 -> 1 + 5 ->
	// them put down and compare
	for i := 0; i < n; i++ {
		sumArr1 -= arr1[i]

		if sumArr1 < result {
			result = min(result, sumArr1+arr1[i])
			break
		}

		result += arr2[i]
	}

	return int64(result)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
