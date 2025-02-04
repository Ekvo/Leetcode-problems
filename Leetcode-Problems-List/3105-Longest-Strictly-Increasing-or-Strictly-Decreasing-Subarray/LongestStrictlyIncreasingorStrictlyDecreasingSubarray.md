## 3105. Longest Strictly Increasing or Strictly Decreasing Subarray => [original page](https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/description/ "https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> |   _Easy_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
You are given an array of integers `nums`. Return the **length** of the longest **subarray** of `nums` which is either **strictly increasing** or **strictly decreasing**.

---
**Example 1:**

**Input:** `nums = [1,4,3,3,2]`  
**Output:** `2`  
**Explanation:**  
`The strictly increasing subarrays of nums are [1], [2], [3], [3], [4], and [1,4].`  
`The strictly decreasing subarrays of nums are [1], [2], [3], [3], [4], [3,2], and [4,3].`  
`Hence, we return 2.`  

**Example 2:**

**Input:** `nums = [3,3,3,3]`  
**Output:** `1`  
**Explanation:**  
`The strictly increasing subarrays of nums are [3], [3], [3], and [3].`  
`The strictly decreasing subarrays of nums are [3], [3], [3], and [3].`  
`Hence, we return 1.`  

**Example 3:**

**Input:** `nums = [3,2,1]`  
**Output:** `3`  
**Explanation:**  
`The strictly increasing subarrays of nums are [3], [2], and [1].`  
`The strictly decreasing subarrays of nums are [3], [2], [1], [3,2], [2,1], and [3,2,1].`  
`Hence, we return 3.` 

---
**Constraints:**

   * $1$ `<= nums.length <=` $50$
   * $1$ `<= nums[i] <`= $50$
 
---
* ### Base data

```Golang
func longestMonotonicSubarray(nums []int) int {	
}
```

---
### [Solution => 3105. Longest Strictly Increasing or Strictly Decreasing Subarray](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3105-Longest-Strictly-Increasing-or-Strictly-Decreasing-Subarray/leetcodethreeonezerofive.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3105-Longest-Strictly-Increasing-or-Strictly-Decreasing-Subarray/leetcodethreeonezerofive.go")

---
* ### Metod => Compare nums[i-1] with nums[i]
```Golang
// one loop with find max on each iteration
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/3105_Longest_Strictly_Increasing_or_Strictly_Decreasing_Subarray.jpg)
