## 1752. Check if Array Is Sorted and Rotated => [original page](https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/description/ "https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> |   _Easy_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
Given an array `nums`, return `true` _if the array was originally sorted in non-decreasing order, then rotated **some** number of positions (including zero). Otherwise, return `false`_.

There may be **duplicates** in the original array.

**Note:** An array `A` rotated by `x` positions results in an array `B` of the same length such that `A[i] == B[(i+x) % A.length]`, where `%` is the modulo operation.

---
**Example 1:**

**Input:** `nums = [3,4,5,1,2]`  
**Output:** `true`  
**Explanation:** `[1,2,3,4,5] is the original sorted array.`  
`You can rotate the array by x = 3 positions to begin on the the element of value 3: [3,4,5,1,2].`  

**Example 2:**

**Input:** `nums = [2,1,3,4]`  
**Output:** `false`  
**Explanation:** `There is no sorted array once rotated that can make nums.`  

**Example 3:**

**Input:** `nums = [1,2,3]`  
**Output:** `true`  
**Explanation:** `[1,2,3] is the original sorted array.`  
`You can rotate the array by x = 0 positions (i.e. no rotation) to make nums.`  

---
**Constraints:**

   * $1$ `<= nums.length <=` $100$
   * $1$ `<= nums[i] <=` $100$
 
---
* ### Base data

```Golang
func check(nums []int) bool {
}
```

---
### [Solution => 1752. Check if Array Is Sorted and Rotated](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1752-Check-if-Array-Is-Sorted-and-Rotated/leetcodeonesevenfivetwo.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1752-Check-if-Array-Is-Sorted-and-Rotated/leetcodeonesevenfivetwo.go")

---
* ### Metod => Find index for less num - 'start'
```Golang
for i:=1;i<len(nums);i++{
   // find 'start' index
}

last:=nums[n-1]
previous:=0

for i:=0;i<start;i++{
   // compare nums[i] with previous and last
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1752_Check_if_Array_Is_Sorted_and_Rotated.jpg)
