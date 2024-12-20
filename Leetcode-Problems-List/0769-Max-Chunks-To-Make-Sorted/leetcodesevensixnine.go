//You are given an integer array arr of length n that represents a permutation
//of the integers in the range [0, n - 1].
//
//We split arr into some number of chunks (i.e., partitions), and individually sort each chunk.
//After concatenating them, the result should equal the sorted array.
//
//Return the largest number of chunks we can make to sort the array.

package leetcodesevensixnine

// Prefix Sums
func maxChunksToSorted(arr []int) int {
	var (
		indexSum       = 0
		integerSum     = 0
		numberOfChunks = 0
	)
	for i := 0; i < len(arr); i++ {
		indexSum += i
		integerSum += arr[i]
		if indexSum == integerSum {
			numberOfChunks++
		}
	}
	return numberOfChunks
}

// Stack
func maxChunksToSorted(arr []int) int {
	s := &stack{}
	for i := 0; i < len(arr); i++ {

		if s.Empty() || s.Top() < arr[i] {
			s.Push(arr[i])
			continue
		}
		integerMax := s.Top()

		for !s.Empty() && arr[i] < s.Top() {
			s.Pop()
		}
		s.Push(integerMax)
	}
	return len(*s)
}

type stack []int

func (s *stack) Empty() bool { return len(*s) == 0 }
func (s *stack) Push(i int)  { *s = append(*s, i) }
func (s *stack) Pop()        { *s = (*s)[:len(*s)-1] }
func (s *stack) Top() int    { return (*s)[len(*s)-1] }
