// Given the root of a binary tree and an integer targetSum,
// return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.
//
// A leaf is a node with no children.
package pathSum

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var finder func(root *TreeNode, currentSum int) bool

	finder = func(root *TreeNode, currentSum int) bool {
		if root == nil {
			return false
		}

		currentSum += root.Val

		if root.Left == nil && root.Right == nil && currentSum == targetSum {
			return true
		}

		return finder(root.Left, currentSum) || finder(root.Right, currentSum)
	}

	return finder(root, 0)
}
