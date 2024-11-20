// Given a binary tree, find its minimum depth.
//
// The minimum depth is the number of nodes along the
// shortest path from the root node down to the nearest leaf node.
//
// Note: A leaf is a node with no children.
package leetcodeoneoneone

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Definition for a binary tree node depth
type nodeLvl struct {
	node *TreeNode
	lvl  int
}

// BFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var (
		result      = 0
		ok          = false
		queue       []*nodeLvl
		currentNode *nodeLvl = nil
	)

	queue = append(queue, &nodeLvl{root, 1})

	for len(queue) > 0 {
		currentNode = queue[0]

		result, ok = nodeNoChildren(*currentNode)

		if ok {
			break
		}

		queue = pushBack(queue, currentNode)
		queue = queue[1:]
	}
	return result
}

func pushBack(q []*nodeLvl, data *nodeLvl) []*nodeLvl {
	if data.node.Left != nil {
		q = append(q, &nodeLvl{data.node.Left, data.lvl + 1})
	}
	if data.node.Right != nil {
		q = append(q, &nodeLvl{data.node.Right, data.lvl + 1})
	}
	return q
}

func nodeNoChildren(data nodeLvl) (int, bool) {
	if data.node != nil && data.node.Left == nil && data.node.Right == nil {
		return data.lvl, true
	}
	return -1, false
}

// DFS
func Leetcode111minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var depth func(root *TreeNode, lvl int) int

	depth = func(root *TreeNode, lvl int) int {
		if root == nil {
			return math.MaxInt64
		}
		if root.Left == nil && root.Right == nil {
			return lvl
		}
		return minLvl(depth(root.Left, lvl+1), depth(root.Right, lvl+1))
	}
	return depth(root, 1)
}

func minLvl(l, r int) int {
	if l < r {
		return l
	}
	return r
}
