## 113. Path Sum II => [original page](https://leetcode.com/problems/path-sum-ii/description/ "https://leetcode.com/problems/path-sum-ii/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**
Given the root of a binary tree and an integer `targetSum`, return all **root-to-leaf** paths where the sum of the node values in the path equals `targetSum`. Each path should be returned as a list of the node **values**, not node references.

A **root-to-leaf** path is a path starting from the root and ending at any leaf node. A **leaf** is a node with no children.

---
**Example 1:**



---
**Constraints:**

   * `The number of nodes in the tree is in the range [0, 5000].`
   * $-1000 <=$ `Node.val` $<= 1000$
   * $-1000 <=$ `targetSum` $<= 1000$


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
func pathSum(root *TreeNode, targetSum int) [][]int {
}
```

---
### [Solution => 0113. Path Sum II](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0113-Path-Sum-II/pathSumII.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0113-Path-Sum-II/pathSumII.go")

---
* ### Metod => Depth First Search and Recursion
```Golang              
    finder(*TreeNode,int,int)
```