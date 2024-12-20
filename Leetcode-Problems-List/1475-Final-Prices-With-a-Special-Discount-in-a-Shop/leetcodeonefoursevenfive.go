// You are given an integer array prices where prices[i] is the price of the ith item in a shop.
//
// There is a special discount for items in the shop.
// If you buy the ith item, then you will receive a discount equivalent
// to prices[j] where j is the minimum index such that j > i and prices[j] <= prices[i].
//
//	Otherwise, you will not receive any discount at all.
//
// Return an integer array answer where answer[i] is the final price you
// will pay for the ith item of the shop, considering the special discount.
package leetcodeonefoursevenfive

func finalPrices(prices []int) []int {
	var (
		arrAnswer = make([]int, len(prices))
		ptrS      = &stack{}
	)
	copy(arrAnswer, prices)

	for i := 0; i < len(prices); i++ {

		for !ptrS.Empty() && prices[ptrS.Top()] >= prices[i] {
			arrAnswer[ptrS.Top()] -= prices[i]
			ptrS.Pop()
		}
		ptrS.Push(i)
	}
	return arrAnswer
}

type stack []int

func (s *stack) Empty() bool { return len(*s) == 0 }
func (s *stack) Push(i int)  { *s = append(*s, i) }
func (s *stack) Pop()        { *s = (*s)[:len(*s)-1] }
func (s *stack) Top() int    { return (*s)[len(*s)-1] }
