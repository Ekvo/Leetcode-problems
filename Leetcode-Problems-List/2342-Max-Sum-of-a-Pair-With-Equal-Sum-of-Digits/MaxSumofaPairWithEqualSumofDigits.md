## 2342. Max Sum of a Pair With Equal Sum of Digits => [original page](https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/description/ "https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> | _Medium_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
You are given a **0-indexed** array nums consisting of **positive** integers. You can choose two indices `i` and `j`, such that `i != j`, and the sum of digits of the number `nums[i]` is equal to that of `nums[j]`.

Return _the **maximum** value of `nums[i] + nums[j]` that you can obtain over all possible indices `i` and `j` that satisfy the conditions_.

---
**Example 1:**

**Input:** `nums = [18,43,36,13,7]`  
**Output:** `54`  
**Explanation:** `The pairs (i, j) that satisfy the conditions are:`  
`- (0, 2), both numbers have a sum of digits equal to 9, and their sum is 18 + 36 = 54.`  
`- (1, 4), both numbers have a sum of digits equal to 7, and their sum is 43 + 7 = 50.`  
`  So the maximum sum that we can obtain is 54.`  

**Example 2:**

**Input:** `nums = [10,12,19,14]`  
**Output:** `-1`  
**Explanation:** `There are no two numbers that satisfy the conditions, so we return -1.`  

---
**Constraints:**

   * $1$ `<= nums.length <=` $10^5$
   * $1$ `<= nums[i] <=` $10^9$
 
---
* ### Base data

```Golang
func maximumSum(nums []int) int {
}
```

---
### [Solution => 2342. Max Sum of a Pair With Equal Sum of Digits](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2342-Max-Sum-of-a-Pair-With-Equal-Sum-of-Digits/leetcodetwothreefourtwo.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2342-Max-Sum-of-a-Pair-With-Equal-Sum-of-Digits/leetcodetwothreefourtwo.go")

---
* ### Metod =>  Array lenght of '82' 
```Golang
// 1 <= nums[i] <= 10^9 ~> 9 * 9 = 81

for _, num := range nums {

   for n := num; n > 0; n /= 10 {
      // find sum of digits
   }
	
   // save sum of digits if it is max
   // find answer	
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2342_Max_Sum_of_a_Pair_With_Equal_Sum_of_Digits.jpg)
