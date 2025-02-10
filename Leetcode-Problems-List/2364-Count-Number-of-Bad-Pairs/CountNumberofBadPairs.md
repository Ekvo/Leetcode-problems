## 2364. Count Number of Bad Pairs => [original page](https://leetcode.com/problems/count-number-of-bad-pairs/description/ "https://leetcode.com/problems/count-number-of-bad-pairs/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> | _Medium_ |
| status                  | ---> | _Solved_ |
    
---
**Task:**  
You are given a **0-indexed** integer array `nums`. A pair of indices `(i, j)` is a **bad pair** if `i < j` and `j - i != nums[j] - nums[i]`.

Return _the total number of **bad pairs** in `nums`_.
 
---
**Example 1:**

**Input:** `nums = [4,1,3,3]`  
**Output:** `5`  
**Explanation:** `The pair (0, 1) is a bad pair since 1 - 0 != 1 - 4.`  
`The pair (0, 2) is a bad pair since 2 - 0 != 3 - 4, 2 != -1.`  
`The pair (0, 3) is a bad pair since 3 - 0 != 3 - 4, 3 != -1.`  
`The pair (1, 2) is a bad pair since 2 - 1 != 3 - 1, 1 != 2.`  
`The pair (2, 3) is a bad pair since 3 - 2 != 3 - 3, 1 != 0.`  
`There are a total of 5 bad pairs, so we return 5.`  

**Example 2:**

**Input:** `nums = [1,2,3,4,5]`  
**Output:** `0`  
**Explanation:** `There are no bad pairs.`  

---
**Constraints:**

   * $1$ `<= nums.length <=` $10^5$
   * $1$ `<= nums[i] <=` $10^9$
 
---
* ### Base data

```Golang
func countBadPairs(nums []int) int64 {
}
```

---
### [Solution => 2364. Count Number of Bad Pairs](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2364-Count-Number-of-Bad-Pairs/leetcodetwothreesixfour.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2364-Count-Number-of-Bad-Pairs/leetcodetwothreesixfour.go")

---
* ### Metod => HashMap & Triangle numbers
```Golang
// Number of Bad Pairs is difference between Good Pairs and Total Pairs;
// Total Pairs can be calculated as 'n∗(n−1)/2' where n is length of nums;
// Good Pair is 'nums[i] - i' and has a sense to store it in Hash Table as '[nums[i] - i, count]';
// So sum of Good Pairs is sum of counts those pairs, which 'count > 1'.
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(N)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2364_Count_Number_of_Bad_Pairs.jpg)
