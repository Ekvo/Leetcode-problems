## 124. Binary Tree Maximum Path Sum => [original page](https://leetcode.com/problems/binary-tree-maximum-path-sum/description/ "https://leetcode.com/problems/binary-tree-maximum-path-sum/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Hard_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
A **path** in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence **at most once**. Note that the path does not need to pass through the root.

The **path sum** of a path is the sum of the node's values in the path.

Given the `root` of a binary tree, return the maximum **path sum** of any **non-empty** path.

---
**Example 1:**

**Input:** `root = [1,2,3]`  
**Output:** `6`  
**Explanation:**  `The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.`

**Example 2:**

**Input:** `root = [-10,9,20,null,null,15,7]`  
**Output:** `42`  
**Explanation:** `The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.`

**Constraints:**
   * `The number of nodes in the tree is in the range [1, 3 * 104].`
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
func maxPathSum(root *TreeNode) int { 
```

---
### [Solution => 124. Binary Tree Maximum Path Sum](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0124-Binary-Tree-Maximum-Path-Sum/binaryTreeMaximumPathSum.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0124-Binary-Tree-Maximum-Path-Sum/binaryTreeMaximumPathSum.go")

---
* ### Metod => Depth First Search and Recursion
```Golang
func sumFromBottom(root *TreeNode) int //go down and start counting
func maxSum(values ...int) int 
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/124_Binary_Tree_Maximum_Path_Sum.jpg)
 
