## 3097. Shortest Subarray With OR at Least K II => [original page](https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-ii/description/ "https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-ii/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You are given an array `nums` of **non-negative** integers and an integer `k`.

An array is called **special** if the bitwise `OR` of all of its elements is **at least** `k`.

Return the length of the **shortest special non-empty** subarray
of `nums`, or return `-1` if no special subarray exists.

---
**Example 1:**

**Input:** `nums = [1,2,3], k = 2`  
**Output:** `1`  
**Explanation:**   
`The subarray [3] has OR value of 3. Hence, we return 1.`

**Example 2:**

**Input:** `nums = [2,1,8], k = 10`  
**Output:** `3`  
**Explanation:**  
`The subarray [2,1,8] has OR value of 11. Hence, we return 3.`

**Example 3:**

**Input:** `nums = [1,2], k = 0`  
**Output:** `1`  
**Explanation:**  
`The subarray [1] has OR value of 1. Hence, we return 1.`

---
**Constraints:**

   * $1$ `<= nums.length <=` $2 * 10^5$
   * $0$ `<= nums[i] <=` $10^9$
   * $0$ `<= k <=` $10^9$
---
* ### Base data

```Golang
func minimumSubarrayLength(nums []int, k int) int {
}
```

---
### [Solution => 3097. Shortest Subarray With OR at Least K II](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3097-Shortest-Subarray-With-OR-at-Least-K-II/leetcodethreezeronineseven.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3097-Shortest-Subarray-With-OR-at-Least-K-II/leetcodethreezeronineseven.go")

---
* ### Metod => Sliding Window
```Golang
for {//count to Right    
	if find  {
	    for {//count to Left         
        }   
    }   
} 
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/3097_Shortest_Subarray_With_OR_at_Least_K_II.jpg)
 
