// A binary tree is uni-valued if every node in the tree has the same value.
//
// Given the root of a binary tree, return true if the given tree is uni-valued, or false otherwise.
package leetcodeninesixfive

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	} else if root.Left != nil && root.Left.Val != root.Val ||
		root.Right != nil && root.Right.Val != root.Val {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
