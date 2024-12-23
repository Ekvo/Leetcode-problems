// Given the root of a binary tree, return the bottom-up level order traversal of its nodes values.
// (i.e., from left to right, level by level from leaf to root).
package leetcodeonezeroseven

import "sort"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var (
		lvlNode   = make(map[int][]int)
		lvlFinder func(root *TreeNode, lvl int)
	)
	lvlFinder = func(root *TreeNode, lvl int) {
		if root == nil {
			return
		}
		lvlNode[lvl] = append(lvlNode[lvl], root.Val)
		lvlFinder(root.Left, lvl+1)
		lvlFinder(root.Right, lvl+1)
	}
	lvlFinder(root, 0)
	var (
		result       = make([][]int, 0, len(lvlNode))
		reverseOrder = make([]int, 0, len(lvlNode))
	)
	for lvl, _ := range lvlNode {
		reverseOrder = append(reverseOrder, lvl)
	}
	sort.Slice(reverseOrder, func(i, j int) bool { return reverseOrder[i] > reverseOrder[j] })
	for _, lvl := range reverseOrder {
		result = append(result, lvlNode[lvl])
	}
	return result
}
