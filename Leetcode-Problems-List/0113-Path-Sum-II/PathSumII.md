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

**Input:** `root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22`  
**Output:** `[[5,4,11,2],[5,8,4,5]]`  
**Explanation:** `There are two paths whose sum equals targetSum:`  
`5 + 4 + 11 + 2 = 22`  
 `5 + 8 + 4 + 5 = 22`

**Example 2:**

**Input:** `root = [1,2,3], targetSum = 5`  
**Output:** `[]`  

**Example 3:**

**Input:** `root = [1,2], targetSum = 0`  
**Output:** `[]`

---
**Constraints:**
   * `The number of nodes in the tree is in the range [0, 5000].`
   * $-1000$ `<= Node.val <=` $1000$
   * $-1000$ `<= targetSum <=` $1000$

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
### [Solution => 113. Path Sum II](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0113-Path-Sum-II/pathSumII.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0113-Path-Sum-II/pathSumII.go")

---
* ### Metod => Depth First Search and Recursion
```Golang              
func finder(*TreeNode,int,int)//root,currentSum,index of array
```
![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/113_Path_Sum_II.jpg)
