## 2458. Height of Binary Tree After Subtree Removal Queries => [original page](https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/description/ "https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Hard_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You are given the `root` of a **binary tree** with `n` nodes. Each node is assigned a unique value from `1` to `n`. You are also given an array `queries` of size `m`.

You have to perform `m` **independent** queries on the tree where in the $i^{th}$ query you do the following:

   * **Remove** the subtree rooted at the node with the value `queries[i]` from the tree. It is **guaranteed** that `queries[i]` will **not** be equal to the value of the root.

Return _an array `answer` of size `m` where `answer[i]` is the height of the tree after performing the $i^{th}$ query_.

**Note:**
   * The queries are independent, so the tree returns to its **initial** state after each query.
   * The height of a tree is the **number of edges in the longest simple path** from the root to some node in the tree.

---
**Example 1:**

**Input:** `root = [1,3,4,2,null,6,5,null,null,null,null,null,7], queries = [4]`  
**Output:** `[2]`  
**Explanation:** `The diagram above shows the tree after removing the subtree rooted at node with value 4.`  
`The height of the tree is 2 (The path 1 -> 3 -> 2).`  

**Example 2:**

**Input:** `root = [5,8,9,2,1,3,7,4,6], queries = [3,2,4,8]`  
**Output:** `[3,2,3,2]`  
**Explanation:** `We have the following queries:`  
`- Removing the subtree rooted at node with value 3. The height of the tree becomes 3 (The path 5 -> 8 -> 2 -> 4).`  
`- Removing the subtree rooted at node with value 2. The height of the tree becomes 2 (The path 5 -> 8 -> 1).`  
`- Removing the subtree rooted at node with value 4. The height of the tree becomes 3 (The path 5 -> 8 -> 2 -> 6).`  
`- Removing the subtree rooted at node with value 8. The height of the tree becomes 2 (The path 5 -> 9 -> 3).`  

---
**Constraints:**

   * `The number of nodes in the tree is n.`
   * $2$ `<= n <=` $10^5$
   * $1$ `<= Node.val <=` $n$
   * `All the values in the tree are unique.`
   * `m == queries.length`
   * $1$ `<= m <=` $min(n, 10^4)$
   * $1$ `<= queries[i] <=` $n$
   * `queries[i] != root.val`
 
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
func treeQueries(root *TreeNode, queries []int) []int {	
}
```

---
### [Solution => 2458. Height of Binary Tree After Subtree Removal Queries](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2458-Height-of-Binary-Tree-After-Subtree-Removal-Queries/leetcodetwofourfiveeight.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2458-Height-of-Binary-Tree-After-Subtree-Removal-Queries/leetcodetwofourfiveeight.go")

---
* ### Metod => DFS, HashMap, Struct & Recursion
```Golang
type number int

type level int

type deapLevelOfNode struct {
	current level
	deapLvl level
}

numberNodes  make(map[number]*depLevelNode)
levelNodes   make(map[level][]level)
lvlFinder   func(node *TreeNode, i level) level

func maxLvl(i, j level) level
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2458_Height_of_Binary_Tree_After_Subtree_Removal_Queries.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2458_Height_of_Binary_Tree_After_Subtree_Removal_Queries_Time.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2458_Height_of_Binary_Tree_After_Subtree_Removal_Queries_Space.jpg)
