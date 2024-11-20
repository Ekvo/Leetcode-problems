## 101. Symmetric Tree => [original page](https://leetcode.com/problems/symmetric-tree/description/ "https://leetcode.com/problems/symmetric-tree/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Easy_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given the `root` of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

---
**Example 1:**

**Input:** `root = [1,2,2,3,4,4,3]`  
**Output:** `true`  

**Example 2:**

**Input:** `root = [1,2,2,null,3,null,3]`  
**Output:** `false`

---
**Constraints:**
   * `The number of nodes in the tree is in the range [1, 1000].`
   * $-100$ `<= Node.val <=` $100$

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
func isSymmetric(root *TreeNode) bool {
}
```

---
### [Solution => 101. Symmetric Tree](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0101-Symmetric-Tree/leetcodeonezeroone.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0101-Symmetric-Tree/leetcodeonezeroone.go")

---
* ### Metod => Depth First Search and Recursion
```Golang
func finder(root.Left, root.Right)
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/101_Symmetric_Tree.jpg)
    