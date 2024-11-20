// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes
// in the sequence has an edge connecting them. A node can only appear in the sequence at most once.
// Note that the path does not need to pass through the root.
//
// The path sum of a path is the sum of the node's values in the path.
//
// Given the root of a binary tree, return the maximum path sum of any non-empty path.
package leetcodeonetwofour

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var (
		pathSum       = math.MinInt64
		sumFromBottom func(root *TreeNode) int
	)
	sumFromBottom = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		var (
			leftSum  = sumFromBottom(root.Left)
			rightSum = sumFromBottom(root.Right)
			rootSum  = leftSum + rightSum + root.Val
		)
		pathSum = maxSum(pathSum, root.Val, leftSum+root.Val, rightSum+root.Val, rootSum)

		return maxSum(root.Val, leftSum+root.Val, rightSum+root.Val)
	}

	sumFromBottom(root)

	return pathSum
}

func maxSum(values ...int) int {
	var maxVal = math.MinInt64
	for _, v := range values {
		if maxVal < v {
			maxVal = v
		}
	}
	return maxVal
}
