## 107. Binary Tree Level Order Traversal II => [original page](https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/ "https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given the `root` of a binary tree, return _the bottom-up level order traversal of its nodes' values._ (i.e., from left to right, level by level from leaf to root).

---
**Example 1:**

**Input:** `root = [3,9,20,null,null,15,7]`  
**Output:** `[[15,7],[9,20],[3]]`  

**Example 2:**

**Input:** `root = [1]`  
**Output:** `[[1]]`  

**Example 3:**

**Input:** `root = []`  
**Output:** `[]`  

---
**Constraints:**

   * `The number of nodes in the tree is in the range [0, 2000].`
   * $-1000$ `<= Node.val <=` $1000$
 
---
* ### Base data

```Golang

// Definition for a binary tree node.
type TreeNode struct {
   Val   int
   Left  *TreeNode
   Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
}
```

---
### [Solution => 107. Binary Tree Level Order Traversal II](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0107-Binary-Tree-Level-Order-Traversal-II/leetcodeonezeroseven.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0107-Binary-Tree-Level-Order-Traversal-II/leetcodeonezeroseven.go")

---
* ### Metod => DFS, HashMap, sort.Slice & Recursion 
```Golang
lvlNode   map[int][]int // key - level, value - values at level
lvlFinder func(root *TreeNode, lvl int)
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/107_Binary_Tree_Level_Order_Traversal_II.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/107_Binary_Tree_Level_Order_Traversal_II_Time.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/107_Binary_Tree_Level_Order_Traversal_II_Space.jpg)
