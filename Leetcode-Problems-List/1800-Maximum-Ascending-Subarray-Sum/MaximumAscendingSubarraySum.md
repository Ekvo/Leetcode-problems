## 1800. Maximum Ascending Subarray Sum => [original page](https://leetcode.com/problems/maximum-ascending-subarray-sum/description/ "https://leetcode.com/problems/maximum-ascending-subarray-sum/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> |   _Easy_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
Given an array of positive integers `nums`, return the _maximum possible sum of an **ascending** subarray in `nums`_.

A subarray is defined as a contiguous sequence of numbers in an array.

A subarray `[numsl, numsl+1, ..., numsr-1, numsr]` is **ascending** if for all `i` where `l <= i < r, numsi  < numsi+1`. Note that a subarray of size `1` is **ascending**.

---
**Example 1:**

**Input:** `nums = [10,20,30,5,10,50]`  
**Output:** `65`  
**Explanation:** `[5,10,50] is the ascending subarray with the maximum sum of 65.`  

**Example 2:**

**Input:** `nums = [10,20,30,40,50]`  
**Output:** `150`  
**Explanation:** `[10,20,30,40,50] is the ascending subarray with the maximum sum of 150.`  

**Example 3:**

**Input:** `nums = [12,17,15,13,10,11,12]`  
**Output:** `33`  
**Explanation:** [10,11,12] is the ascending subarray with the maximum sum of 33.`  

---
**Constraints:**

   * $1$ `<= nums.length <=` $100$
   * $1$ `<= nums[i] <=` $100$
 
---
* ### Base data

```Golang
func maxAscendingSum(nums []int) int {
}
```

---
### [Solution => 1800. Maximum Ascending Subarray Sum](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1800-Maximum-Ascending-Subarray-Sum/leetcodeomeeightzerozero.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1800-Maximum-Ascending-Subarray-Sum/leetcodeomeeightzerozero.go")

---
* ### Metod => One cycle
```Golang
for i := 1; i < len(nums); i++ {
   // if previos num less current num -> increase current summ 
   // else current summ -> current num
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1800_Maximum_Ascending_Subarray_Sum.jpg)
