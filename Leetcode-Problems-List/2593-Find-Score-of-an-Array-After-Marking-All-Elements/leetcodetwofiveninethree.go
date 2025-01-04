// You are given an array nums consisting of positive integers.
//
// Starting with score = 0, apply the following algorithm:
//
// Choose the smallest integer of the array that is not marked.
// If there is a tie, choose the one with the smallest index.
// Add the value of the chosen integer to score.
// Mark the chosen element and its two adjacent elements if they exist.
// Repeat until all the array elements are marked.
//
// Return the score you get after applying the above algorithm.
package leetcodetwofiveninethree

import "sort"

func findScore(nums []int) int64 {
	var (
		store = make([]*indexValue, 0, len(nums))

		//+2 ~> nums[-1], nums[len(nums)]
		visited = make([]bool, len(nums)+2)
	)
	for i := 0; i < len(nums); i++ {
		store = append(store, &indexValue{i + 1, nums[i]})
	}
	sort.Slice(store, func(i, j int) bool {
		return store[i].value < store[j].value ||
			(store[i].value == store[j].value &&
				store[i].index < store[j].index)
	})
	var sum = 0

	for i := 0; i < len(store); i++ {
		if !visited[store[i].index] {
			sum += store[i].value
			visited[store[i].index-1] = true
			visited[store[i].index] = true
			visited[store[i].index+1] = true
		}
	}
	return int64(sum)
}

type indexValue struct {
	index int
	value int
}
