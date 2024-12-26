## 515. Find Largest Value in Each Tree Row => [original page](https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/ "https://leetcode.com/problems/find-largest-value-in-each-tree-row/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given the `root` of a binary tree, return _an array of the largest value in each row_ of the tree (**0-indexed**).

---
**Example 1:**

**Input:** `root = [1,3,2,5,3,null,9]`  
**Output:** `[1,3,9]`  

**Example 2:**

**Input:** `root = [1,2,3]`  
**Output:** `[1,3]`  

---
**Constraints:**

   * `The number of nodes in the tree will be in the range` $[0, 10^4]$.
   * $-2^{31}$ `<= Node.val <=` $2^{31} - 1$
 
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
func largestValues(root *TreeNode) []int {	
}
```

---
### [Solution => 515. Find Largest Value in Each Tree Row](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0515-Find-Largest-Value-in-Each-Tree-Row/leetcodefiveonefive.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0515-Find-Largest-Value-in-Each-Tree-Row/leetcodefiveonefive.go")

---
* ### Metod => BFS & Queue
```Golang
queue      []*TreeNode
func addTreeNode(q []*TreeNode, root *TreeNode) []*TreeNode
func max(a,b int) int
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/515_Find_Largest_Value_in_Each_Tree_Row_BFS.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/515_Find_Largest_Value_in_Each_Tree_Row_BFS_Time.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/515_Find_Largest_Value_in_Each_Tree_Row__BFS_Space.jpg)

---
* ### Metod => DFS & Recursion
```Golang
lvlFinder  func(root *TreeNode, lvl int)
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/515_Find_Largest_Value_in_Each_Tree_Row_DFS.jpg)

