## 1. Two Sum => [original page](https://leetcode.com/problems/two-sum/description/ "https://leetcode.com/problems/two-sum/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Easy_   |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

You can return the answer in any order.

---
**Example 1:**

**Input:** `nums = [2,7,11,15], target = 9`  
**Output:** `[0,1]`  
**Explanation:** `Because nums[0] + nums[1] == 9, we return [0, 1].`

**Example 2:**

**Input:** `nums = [3,2,4], target = 6`  
**Output:** `[1,2]`

**Example 3:**

**Input:** `nums = [3,3], target = 6`  
**Output:** `[0,1]`

---
**Constraints:**

  * $2$ `<= nums.length <=` $10^4$

  * $-10^9$ `<= nums[i] <=1` $10^9$

  * $-10^9$ `<= target <=` $10^9$

  * `Only one valid answer exists.`

---
* ### Base data 

```Golang
func twoSum(nums []int, target int) []int {
}
```
---
### [Solution => 1. Two Sum](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0001-Two-Sum/leetcodeone.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0001-Two-Sum/leetcodeone.go")

---
* ### First metod => HashMap
```Golang
    valueIndex = make(map[int]int,len))
    for i;i<len;i++
```
![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1_Two_Sum_HashMap.jpg)

---
* ### Second metod => Brute Force
```Golang
    for i;i<len;i++{
        for j = i+1; j<len; j++{
		//find result
        }       
    }
``` 
![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1_Two_Sum_BruteForce.jpg)







