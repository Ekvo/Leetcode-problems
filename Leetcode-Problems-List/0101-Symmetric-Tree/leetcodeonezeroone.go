// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
package leetcodeonezeroone

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		//processed nillptr
		return false
	}

	var finder func(left, right *TreeNode) bool

	finder = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		return finder(left.Left, right.Right) && finder(left.Right, right.Left)
	}

	return finder(root.Left, root.Right)
}
