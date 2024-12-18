// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
package leetcodetwo

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var (
		buffer  = 0
		l3      = &ListNode{0, nil}
		current = l3
	)
	for buffer > 0 || l1 != nil || l2 != nil {
		if l1 != nil {
			current.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			current.Val += l2.Val
			l2 = l2.Next
		}
		current.Val += buffer
		buffer = 0

		if current.Val > 9 {
			buffer = current.Val / 10
			current.Val %= 10
		}
		if current.Next == nil && (buffer > 0 || l1 != nil || l2 != nil) {
			current.Next = &ListNode{0, nil}
		}
		current = current.Next
	}
	return l3
}
