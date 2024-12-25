## 965. Univalued Binary Tree => [original page](https://leetcode.com/problems/univalued-binary-tree/description/ "https://leetcode.com/problems/univalued-binary-tree/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Easy_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
A binary tree is **uni-valued** if every node in the tree has the same value.

Given the `root` of a binary `tree`, return `true` _if the given tree is **uni-valued**, or `false` otherwise._

---
**Example 1:**

**Input:** `root = [1,1,1,1,1,null,1]`  
**Output:** `true`  

**Example 2:**

**Input:** `root = [2,2,2,5,2]`  
**Output:** `false`  

---
**Constraints:**

   * `The number of nodes in the tree is in the range` $[1, 100]$.
   * $0$ `<= Node.val <` $100$
 
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
func isUnivalTree(root *TreeNode) bool {
}
```

---
### [Solution => 965. Univalued Binary Tree](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0965-Univalued-Binary-Tree/leetcodeninesixfive.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0965-Univalued-Binary-Tree/leetcodeninesixfive.go")

---
* ### Metod => DFS & Recursion
```Golang
 // isUnivalTree calls isUnivalTree
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/965_Univalued_Binary_Tree.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/965_Univalued_Binary_Tree_Time.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/965_Univalued_Binary_Tree_Space.jpg)
