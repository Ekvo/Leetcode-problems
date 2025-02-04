## 3151. Special Array I => [original page](https://leetcode.com/problems/special-array-i/description/ "https://leetcode.com/problems/special-array-i/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> |   _Easy_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
An array is considered **special** if every pair of its adjacent elements contains two numbers with different parity.

You are given an array of integers `nums`. Return `true` if nums is a **special** array, otherwise, return `false`.

---
**Example 1:**

**Input:** `nums = [1]`  
**Output:** `true`  
**Explanation:**  
`There is only one element. So the answer is true.`  

**Example 2:**

**Input:** `nums = [2,1,4]`  
**Output:** `true`  
**Explanation:**  
`There is only two pairs: (2,1) and (1,4), and both of them contain numbers with different parity. So the answer is true.`  

**Example 3:**

**Input:** `nums = [4,3,1,6]`  
**Output:** `false`  
**Explanation:**  
`nums[1] and nums[2] are both odd. So the answer is false.`  

---
**Constraints:**

   * $1$ `<= nums.length <=` $100$
   * $1$ `<= nums[i] <=` $100$
 
---
* ### Base data

```Golang
func isArraySpecial(nums []int) bool {	
}
```

---
### [Solution => 3151. Special Array I](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3151-Special-Array-I/leetcodethreeonefiveone.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3151-Special-Array-I/leetcodethreeonefiveone.go")

---
* ### Metod => Compare two numbers
```Golang
// one loop with comparison of neighboring numbers
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/3151_Special_Array_I.jpg)
