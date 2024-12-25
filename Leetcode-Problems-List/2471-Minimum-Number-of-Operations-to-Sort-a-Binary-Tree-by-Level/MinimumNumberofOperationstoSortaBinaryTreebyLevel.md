## 2471. Minimum Number of Operations to Sort a Binary Tree by Level => [original page](https://leetcode.com/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/description/ "https://leetcode.com/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You are given the `root` of a binary tree with **unique values**.

In one operation, you can choose any two nodes **at the same level** and swap their values.

Return _the minimum number of operations needed to make the values at each level sorted in a **strictly increasing order**._

The **level** of a node is the number of edges along the path between it and the root node.

---
**Example 1:**

**Input:** `root = [1,4,3,7,6,8,5,null,null,null,null,9,null,10]`  
**Output:** `3`  
**Explanation:**  
`- Swap 4 and 3. The 2nd level becomes [3,4].`  
`- Swap 7 and 5. The 3rd level becomes [5,6,8,7].`  
`- Swap 8 and 7. The 3rd level becomes [5,6,7,8].`  
`We used 3 operations so return 3.`  
`It can be proven that 3 is the minimum number of operations needed.`  

**Example 2:**

**Input:** `root = [1,3,2,7,6,5,4]`  
**Output:** `3`  
**Explanation:**  
`- Swap 3 and 2. The 2nd level becomes [2,3].`  
`- Swap 7 and 4. The 3rd level becomes [4,6,5,7].`  
`- Swap 6 and 5. The 3rd level becomes [4,5,6,7].`  
`We used 3 operations so return 3.`  
`It can be proven that 3 is the minimum number of operations needed.`  

**Example 3:**

**Input:** `root = [1,2,3,4,5,6]`  
**Output:** `0`  
**Explanation:** `Each level is already sorted in increasing order so return 0.`  

---
**Constraints:**

   * `The number of nodes in the tree is in the range` $[1, 10^5].$
   * $1$ `<= Node.val <=` $10^5$
   * `All the values of the tree are unique.`
 
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
func minimumOperations(root *TreeNode) int {	
}
```

---
### [Solution => 2471. Minimum Number of Operations to Sort a Binary Tree by Level](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2471-Minimum-Number-of-Operations-to-Sort-a-Binary-Tree-by-Level/leetcodetwofoursevenone.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2471-Minimum-Number-of-Operations-to-Sort-a-Binary-Tree-by-Level/leetcodetwofoursevenone.go")

---
* ### Metod => Queue & HashMap  
```Golang
queue  []*TreeNode
func addTreeNode(q []*TreeNode, root *TreeNode) []*TreeNode
func getOperations(lvlArr []int) int
   valPos map[int]int // key - integer in lvlArr, value - index of integer in lvlArr
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2471_Minimum_Number_of_Operations_to_Sort_a_Binary_Tree_by_Level.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2471_Minimum_Number_of_Operations_to_Sort_a_Binary_Tree_by_Level_Time.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2471_Minimum_Number_of_Operations_to_Sort_a_Binary_Tree_by_Level_Space.jpg)