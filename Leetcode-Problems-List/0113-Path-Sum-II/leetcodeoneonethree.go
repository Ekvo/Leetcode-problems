// Given the root of a binary tree and an integer targetSum,
// return all root-to-leaf paths where the sum of the node values in
// the path equals targetSum. Each path should be returned as a list of the node values, not node references.
//
// A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.
package leetcodeoneonethree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	var (
		result  [][]int
		arrStep []int
		finder  func(root *TreeNode, currentSum, index int)
	)

	finder = func(root *TreeNode, currentSum, index int) {
		if root == nil {
			return
		}

		if len(arrStep) <= index {
			arrStep = append(arrStep, root.Val)
		} else {
			arrStep[index] = root.Val
		}

		currentSum += root.Val

		if root.Left == nil && root.Right == nil && currentSum == targetSum {
			var arrSave = make([]int, index+1)
			copy(arrSave, arrStep[:index+1])
			result = append(result, arrSave)
			return
		}

		finder(root.Left, currentSum, index+1)
		finder(root.Right, currentSum, index+1)
	}

	finder(root, 0, 0)

	return result
}
