// Given the root of a binary tree, return an array of the largest value in each row of the tree (0-indexed).
package leetcodefiveonefive

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		queue      []*TreeNode
		largeValue []int
	)
	queue = append(queue, root)

	for len(queue) > 0 {
		loop := len(queue)
		maxVal := math.MinInt64

		for i := 0; i < loop; i++ {
			curNode := queue[0]

			maxVal = max(maxVal, curNode.Val)

			queue = addTreeNode(queue, curNode)
			queue = queue[1:]
		}
		largeValue = append(largeValue, maxVal)
	}
	return largeValue
}

func addTreeNode(q []*TreeNode, root *TreeNode) []*TreeNode {
	if root.Left != nil {
		q = append(q, root.Left)
	}
	if root.Right != nil {
		q = append(q, root.Right)
	}
	return q
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// DFS
func largestValues(root *TreeNode) []int {
	var (
		largeValue []int
		lvlFinder  func(root *TreeNode, lvl int)
	)
	lvlFinder = func(root *TreeNode, lvl int) {
		if root == nil {
			return
		}
		if len(largeValue) <= lvl {
			largeValue = append(largeValue, root.Val)
		} else if largeValue[lvl] < root.Val {
			largeValue[lvl] = root.Val
		}
		lvlFinder(root.Left, lvl+1)
		lvlFinder(root.Right, lvl+1)
	}
	lvlFinder(root, 0)

	return largeValue
}
