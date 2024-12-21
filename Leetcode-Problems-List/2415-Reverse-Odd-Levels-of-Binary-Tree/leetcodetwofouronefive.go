//Given the root of a perfect binary tree, reverse the node values at each odd level of the tree.
//
//For example, suppose the node values at level 3 are [2,1,3,4,7,11,29,18], then it should become [18,29,11,7,4,3,1,2].
//
//Return the root of the reversed tree.
//
//A binary tree is perfect if all parent nodes have two children and all leaves are on the same level.
//
//The level of a node is the number of edges along the path between it and the root node.

package leetcodetwofouronefive

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func reverseOddLevels(root *TreeNode) *TreeNode {
	var (
		lvlNodes  = make(map[int][]*TreeNode)
		lvlFinder func(node *TreeNode, lvl int)
	)
	lvlFinder = func(node *TreeNode, lvl int) {
		if node == nil {
			return
		}
		if lvl%2 != 0 {
			lvlNodes[lvl] = append(lvlNodes[lvl], node)
		}
		lvlFinder(node.Left, lvl+1)
		lvlFinder(node.Right, lvl+1)
	}
	lvlFinder(root, 0)

	for _, nodes := range lvlNodes {
		reverseNodesDFS(nodes)
	}
	return root
}

func reverseNodesDFS(level []*TreeNode) {
	h, t := 0, len(level)-1
	for h < t {
		level[h].Val, level[t].Val = level[t].Val, level[h].Val
		h++
		t--
	}
}

// BFS
func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var (
		queue                = []*treeNodeLvl{}
		reverse              = []*treeNodeLvl{}
		curNode *treeNodeLvl = nil
	)
	queue = append(queue, &treeNodeLvl{root, 0})

	for len(queue) > 0 {
		curNode = queue[0]
		queue = addNodes(queue, curNode)
		if curNode.lvl%2 != 0 {
			reverse = append(reverse, curNode)
		} else if len(reverse) > 0 {
			reverse = reverseNodesBFS(reverse)
		}
		queue = queue[1:]
	}
	reverseNodesBFS(reverse)

	return root
}

type treeNodeLvl struct {
	node *TreeNode
	lvl  int
}

func addNodes(queue []*treeNodeLvl, root *treeNodeLvl) []*treeNodeLvl {
	if root.node.Left != nil {
		queue = append(queue, &treeNodeLvl{root.node.Left, root.lvl + 1})
	}
	if root.node.Right != nil {
		queue = append(queue, &treeNodeLvl{root.node.Right, root.lvl + 1})
	}
	return queue
}

func reverseNodesBFS(reverse []*treeNodeLvl) []*treeNodeLvl {
	h, t := 0, len(reverse)-1
	for h < t {
		reverse[h].node.Val, reverse[t].node.Val = reverse[t].node.Val, reverse[h].node.Val
		h++
		t--
	}
	reverse = reverse[:0]
	return reverse
}
