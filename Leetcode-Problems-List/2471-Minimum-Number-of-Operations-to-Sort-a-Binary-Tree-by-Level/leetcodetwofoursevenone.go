// You are given the root of a binary tree with unique values.
//
// In one operation, you can choose any two nodes at the same level and swap their values.
//
// Return the minimum number of operations needed to make the values at each level sorted in a strictly increasing order.
//
// The level of a node is the number of edges along the path between it and the root node.
package leetcodetwofoursevenone

import "sort"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) int {
	var (
		operations = 0
		queue      []*TreeNode
		lvlArr     []int
	)
	queue = append(queue, root)

	for len(queue) > 0 {
		loop := len(queue)

		for i := 0; i < loop; i++ {
			curNode := queue[0]
			queue = queue[1:]

			lvlArr = append(lvlArr, curNode.Val)
			queue = addTreeNode(queue, curNode)
		}
		operations += getOperations(lvlArr)
		lvlArr = lvlArr[:0]
	}
	return operations
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

func getOperations(lvlArr []int) int {
	localOperations := 0

	valPos := make(map[int]int, len(lvlArr))
	for i, val := range lvlArr {
		valPos[val] = i
	}

	localArr := make([]int, len(lvlArr))
	copy(localArr, lvlArr)
	sort.Ints(localArr)

	for i := 0; i < len(lvlArr); i++ {
		if lvlArr[i] != localArr[i] {
			localOperations++

			curPos := valPos[localArr[i]]
			valPos[lvlArr[i]] = curPos
			lvlArr[curPos], lvlArr[i] = lvlArr[i], lvlArr[curPos]
		}
	}
	return localOperations
}
