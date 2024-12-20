## 2601. Prime Subtraction Operation => [original page](https://leetcode.com/problems/prime-subtraction-operation/description/ "https://leetcode.com/problems/rotate-list/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You are given a **0-indexed** integer array nums of length n.

You can perform the following operation as many times as you want:

   * Pick an index `i` that you haven’t picked before, and pick a prime p **strictly less than** `nums[i]`, then subtract `p` from `nums[i]`.

*Return true if you can make **nums** a strictly increasing array using the above operation and false otherwise.*

A **strictly increasing array** is an array whose each element is strictly greater than its preceding element.

---
**Example 1:**

**Input:** `nums = [4,9,6,10]`  
**Output:** `true`  
**Explanation:** `In the first operation:`  
`Pick i = 0 and p = 3, and then subtract 3 from nums[0], so that nums becomes [1,9,6,10].`  
`In the second operation: i = 1, p = 7, subtract 7 from nums[1], so nums becomes equal to [1,2,6,10].`  
`After the second operation, nums is sorted in strictly increasing order, so the answer is true.`  

**Example 2:**

**Input:** `nums = [6,8,11,12]`  
**Output:** `true`  
**Explanation:** `Initially nums is sorted in strictly increasing order, so we don't need to make any operations.`

**Example 3:**

**Input:** `nums = [5,8,3]`  
**Output:** `false`  
**Explanation:** `It can be proven that there is no way to perform operations to make nums sorted in strictly increasing order, so the answer is false.`

---

**Constraints:**
   * $1$ `<= nums.length <=` $1000$

   * $1$ `<= nums[i] <=` $1000$

   * `nums.length == n`

---

### Base data

```Golang
func primeSubOperation(nums []int) bool {	
}
```

---
### [Solution => 2601. Prime Subtraction Operation](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2601-Prime-Subtraction-Operation/leetcodetwosixzeroone.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2601-Prime-Subtraction-Operation/leetcodetwosixzeroone.go")

---
## Metod => Binary Seach

* ### Skeleton

```Golang
func primeSubOperation([]int) bool   //main function
func toLessValue([]int,int,int) int  //decreases nums[i] value
func primeNumbers(int,int) []int     //create and return an array of prime numbers
```
* ### Main idea for binary seach metod for 2601

```Golang
    primeNumbers []int
    for //binary search for nearest tail of larger number for nums[i] in primeNumber
```
![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2601_Prime_Subtraction_Operation.jpg) 
