// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
//
// You can return the answer in any order.
package twoSum

// HashMap metod
func twoSum(nums []int, target int) []int {
	var valueIndex = make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if j, ok := valueIndex[target-nums[i]]; ok {
			return []int{i, j}
		}
		valueIndex[nums[i]] = i
	}

	return nil
}

// Brute Force metod
func twoSum(nums []int, target int) []int {
	var (
		loop   = len(nums)
		result = make([]int, 0, 2)
	)

	for i := 0; i < loop && len(result) < 2; i++ {
		for j := i + 1; j < loop; j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				break
			}
		}
	}

	return result
}
