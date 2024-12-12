// You are given an integer array gifts denoting the number of gifts in various piles.
// Every second, you do the following:
//
// Choose the pile with the maximum number of gifts.
// If there is more than one pile with the maximum number of gifts, choose any.
// Leave behind the floor of the square root of the number of gifts in the pile. Take the rest of the gifts.
//
// Return the number of gifts remaining after k seconds.
package leetcodetwofivefiveeight

import (
	"container/heap"
	"math"
)

func pickGifts(gifts []int, k int) int64 {
	store := Gift(gifts)
	heap.Init(&store)

	for i := 0; i < k; i++ {
		store[0] = int(math.Sqrt(float64(store[0])))
		heap.Fix(&store, 0)
	}
	numberOfGifts := 0
	for i := 0; i < store.Len(); i++ {
		numberOfGifts += store[i]
	}
	return int64(numberOfGifts)
}

type Gift []int

func (obj Gift) Len() int           { return len(obj) }
func (obj Gift) Less(i, j int) bool { return obj[i] > obj[j] }
func (obj Gift) Swap(i, j int)      { obj[i], obj[j] = obj[j], obj[i] }
func (obj *Gift) Push(x any)        { *obj = append(*obj, x.(int)) }
func (obj *Gift) Pop() any {
	old := *obj
	n := len(old)
	result := old[n-1]
	*obj = old[:n-1]
	return result
}
