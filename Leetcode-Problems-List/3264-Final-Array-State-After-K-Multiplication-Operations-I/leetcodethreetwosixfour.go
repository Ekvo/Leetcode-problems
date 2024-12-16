// You are given an integer array nums, an integer k, and an integer multiplier.
//
// You need to perform k operations on nums. In each operation:
//
// Find the minimum value x in nums. If there are multiple occurrences of the minimum value, select the one that appears first.
// Replace the selected minimum value x with x * multiplier.
//
// Return an integer array denoting the final state of nums after performing all k operations.
package leetcodethreetwosixfour

import "container/heap"

func getFinalState(nums []int, k int, multiplier int) []int {
	var arrIndexNums = make(indexNums, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		arrIndexNums = append(arrIndexNums, &indexNum{i, nums})
	}
	heap.Init(&arrIndexNums)

	for i := 0; i < k; i++ {
		nums[arrIndexNums[0].Index] *= multiplier
		heap.Fix(&arrIndexNums, 0)
	}
	return nums
}

type indexNums []*indexNum

type indexNum struct {
	Index int
	Nums  []int
}

func (iN indexNums) Len() int      { return len(iN) }
func (iN indexNums) Swap(i, j int) { iN[i], iN[j] = iN[j], iN[i] }
func (iN indexNums) Less(i, j int) bool {
	return iN[i].Nums[iN[i].Index] < iN[j].Nums[iN[j].Index] ||
		(iN[i].Nums[iN[i].Index] == iN[j].Nums[iN[j].Index] &&
			iN[i].Index < iN[j].Index)
}
func (iN *indexNums) Push(obj any) { *iN = append(*iN, obj.(*indexNum)) }
func (iN *indexNums) Pop() any {
	old := *iN
	lenght := len(old)
	elem := old[lenght-1]
	*iN = old[:lenght-1]
	return elem
}
