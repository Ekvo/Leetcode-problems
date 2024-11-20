// Given the head of a linked list, rotate the list to the right by k places.
package leetcodesixone

type ListNode struct {
	Val  int
	Next *ListNode
}

// Cycle List metod
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		lenght  = 1
		current = head
	)

	for {
		if current.Next == nil { //tail -> head
			current.Next = head
			break
		}
		lenght++
		current = current.Next
	}

	var permutations = 0

	if lenght < k {
		//required for one cycle List (head -> tail)
		permutations = lenght - (k - (k/lenght)*lenght)
	} else if lenght > k {
		permutations = lenght - k
	} else {
		current.Next = nil
		return head
	}

	//One cycle List
	for ; permutations > 0; permutations-- {
		current = head
		head = head.Next
	}
	current.Next = nil

	return head
}
