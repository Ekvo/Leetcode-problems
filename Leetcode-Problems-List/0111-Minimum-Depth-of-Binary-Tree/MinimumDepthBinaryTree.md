## 111. Minimum Depth of Binary Tree => [original page](https://leetcode.com/problems/minimum-depth-of-binary-tree/description/ "https://leetcode.com/problems/minimum-depth-of-binary-tree/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Easy_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

**Note:** A leaf is a node with no children.

---
**Example 1:**

**Input:** `root = [3,9,20,null,null,15,7]`  
**Output:** `2`

**Example 2:**

**Input:** `root = [2,null,3,null,4,null,5,null,6]`  
**Output:** `5`

---
**Constraints:**
   * `The number of nodes in the tree is in the range [`$0, 10^5$`].`
   * $-1000$ `<= Node.val <=` $1000$

---
* ### Base data

```Golang
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
}
```

---
### [Solution => 111. Minimum Depth of Binary Tree](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0111-Minimum-Depth-of-Binary-Tree/minimumDepthBinaryTree.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0111-Minimum-Depth-of-Binary-Tree/minimumDepthBinaryTree.go")

---
* ### Metod => BFS and Queue
```Golang

type nodeLvl struct {
    node *TreeNode
    lvl  int
}
func pushBack(q []*nodeLvl, data *nodeLvl) []*nodeLvl
func nodeNoChildren(data nodeLvl) (int, bool)
```
![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/111_Minimum_Depth_of_Binary_Tree_BFS.jpg)

---
* ### Metod => DFS and Recursion
```Golang
func depth(root *TreeNode, lvl int) int
func minlvl(l,r int) int
```
![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/111_Minimum_Depth_of_Binary_Tree_DFS.jpg)



 
