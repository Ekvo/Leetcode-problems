//You are given the root of a binary tree with n nodes. Each node is assigned a unique value from 1 to n.
//You are also given an array queries of size m.
//
//You have to perform m independent queries on the tree where in the ith query you do the following:
//
//Remove the subtree rooted at the node with the value queries[i] from the tree.
//It is guaranteed that queries[i] will not be equal to the value of the root.
//
//Return an array answer of size m where answer[i] is the height of the tree after performing the ith query.
//
//Note:
// The queries are independent, so the tree returns to its initial state after each query.
// The height of a tree is the number of edges in the longest simple path from the root to some node in the tree.
package leetcodetwofourfiveeight

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type number int

type level int

type deapLevelOfNode struct {
	current level
	deapLvl level
}

func treeQueries(root *TreeNode, queries []int) []int {
	var (
		numberNodes = make(map[number]*deapLevelOfNode)
		levelNodes  = make(map[level][]level)
		lvlFinder   func(node *TreeNode, i level) level
	)

	lvlFinder = func(n *TreeNode, i level) level {
		if n == nil {
			return i - 1
		}

		var dpLN = &deapLevelOfNode{
			current: i,
			deapLvl: i,
		}

		numberNodes[number(n.Val)] = dpLN
		dpLN.deapLvl = maxLvl(lvlFinder(n.Left, i+1), lvlFinder(n.Right, i+1))

		if len(levelNodes[i]) < 2 {
			levelNodes[i] = make([]level, 2)
			levelNodes[i][0] = dpLN.deapLvl
			levelNodes[i][1] = dpLN.current - 1
		} else if levelNodes[i][0] <= dpLN.deapLvl {
			levelNodes[i][1] = levelNodes[i][0]
			levelNodes[i][0] = dpLN.deapLvl
		} else if levelNodes[i][1] < dpLN.deapLvl {
			levelNodes[i][1] = dpLN.deapLvl
		}

		return dpLN.deapLvl
	}

	lvlFinder(root, 0)

	for i := 0; i < len(queries); i++ {

		if obj, ex := numberNodes[number(queries[i])]; ex {

			if levelNodes[obj.current][0] == obj.deapLvl {
				queries[i] = int(levelNodes[obj.current][1])
			} else {
				queries[i] = int(levelNodes[obj.current][0])
			}
		}
	}

	return queries
}

func maxLvl(i, j level) level {
	if i < j {
		return j
	}
	return i
}
