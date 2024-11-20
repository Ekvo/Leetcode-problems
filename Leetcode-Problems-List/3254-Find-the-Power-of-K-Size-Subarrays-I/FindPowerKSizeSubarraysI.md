## 3254. Find the Power of K-Size Subarrays I => [original page](https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/description/ "https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You are given an array of integers `nums` of length n and a positive integer `k`.

The **power** of an array is defined as:

  *  Its maximum element if all of its elements are consecutive and sorted in ascending order.
  *  -1 otherwise.

You need to find the **power** of all
subarrays
of nums of size k.

Return an integer array `results` of size `n - k + 1`, where `results[i]` is the power of `nums[i..(i + k - 1)]`.

---
**Example 1:**

**Input:** `nums = [1,2,3,4,3,2,5], k = 3`  
**Output:** `[3,4,-1,-1,-1]`  
**Explanation:** `There are 5 subarrays of nums of size 3:`  
`[1, 2, 3] with the maximum element 3.`  
`[2, 3, 4] with the maximum element 4.`  
`[3, 4, 3] whose elements are not consecutive.`  
`[4, 3, 2] whose elements are not sorted.`
`[3, 2, 5] whose elements are not consecutive.`

**Example 2:**

**Input:** `nums = [2,2,2,2,2], k = 4`  
**Output:** `[-1,-1]`     

**Example 3:**

**Input:** `nums = [3,2,3,2,3,2], k = 2`  
**Output:** `[-1,3,-1,3,-1]` 

---
**Constraints:**  
   * $1$ `<= n == nums.length <=` $500$
   * $1$ `<= nums[i] <=` $10^5$
   * $1$ `<= k <=` $n$

---
* ### Base data

```Golang
func resultsArray(nums []int, k int) []int {
}
```

---
### [Solution => 3254. Find the Power of K-Size Subarrays I](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3254-Find-the-Power-of-K-Size-Subarrays-I/leetcodethreetwofivefour.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3254-Find-the-Power-of-K-Size-Subarrays-I/leetcodethreetwofivefour.go")

---
* ### Metod => Small Big Steps on the Array
```Golang
func resultsArray(nums []int, k int) []int
type subArr struct
func NewSubArr(data []int) *subArr
func (sA *subArr) nextSubArr(data []int)
func (sA *subArr) addToIndex() int 
func (sA *subArr) powerValue() int //append to result
func (sA *subArr) compute()        //the heart of the program
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/3254_Find_the_Power_of_K-Size_Subarrays_I.jpg)
 
