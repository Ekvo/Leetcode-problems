// Solution to problem 75 from leetcode
package leetcodessevenfive

func sortColors(nums []int) {
	colorCount := make([]int, 3)

	for _, num := range nums {
		colorCount[num]++
	}

	index := 0

	for color, count := range colorCount {

		for ; count > 0; count-- {
			nums[index] = color
			index++
		}
	}
}
