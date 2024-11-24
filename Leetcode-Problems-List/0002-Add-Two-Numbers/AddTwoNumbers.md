## 2. Add Two Numbers => [original page](https://leetcode.com/problems/add-two-numbers/description/ "https://leetcode.com/problems/add-two-numbers/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

---
**Example 1:**

**Input:** `l1 = [2,4,3], l2 = [5,6,4]`  
**Output:** `[7,0,8]`  
**Explanation:** `342 + 465 = 807.`  

**Example 2:**

**Input:** `l1 = [0], l2 = [0]`  
**Output:** `[0]`  

**Example 3:**

**Input:** `l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]`  
**Output:** `[8,9,9,9,0,0,0,1]`  

---
**Constraints:**

   * `The number of nodes in each linked list is in the range [1, 100].`
   * $0$ `<= Node.val <=` $9$
   * `It is guaranteed that the list represents a number that does not have leading zeros.`

---
* ### Base data

```Golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
}
```

---
### [Solution => 2. Add Two Numbers](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0002-Add-Two-Numbers/leetcodetwo.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0002-Add-Two-Numbers/leetcodetwo.go")

---
* ### Metod => Linear filling
```Golang
for l1!=nil||l2!=nil|| buffer>0 {
	//add to the list
}
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2_Add_Two_Numbers.jpg)
 
