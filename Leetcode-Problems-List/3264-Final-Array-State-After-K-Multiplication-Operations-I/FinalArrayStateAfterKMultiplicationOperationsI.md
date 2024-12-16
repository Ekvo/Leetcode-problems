## 3264. Final Array State After K Multiplication Operations I => [original page](https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/description/ "https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Easy_ |
| status                 | ---> | _Splved_ |

---
**Task:**  
You are given an integer array `nums`, an integer `k`, and an integer `multiplier`.

You need to perform `k` operations on `nums`. In each operation:

   * Find the **minimum** value `x` in `nums`. If there are multiple occurrences of the minimum value, select the one that appears **first**.
   * Replace the selected minimum value `x` with `x * multiplier`.

Return an integer array denoting the _final state of `nums` after performing all `k` operations._

---
**Example 1:**  

**Input:** `nums = [2,1,3,5,6], k = 5, multiplier = 2`  
**Output:** `[8,4,6,5,6]`  
**Explanation:**  

| Operation         | Result             | 
|-------------------|--------------------|
| After operation 1 | 	`[2, 2, 3, 5, 6]` | 
| After operation 2 | 	`[4, 2, 3, 5, 6]` | 
| After operation 3 | 	`[4, 4, 3, 5, 6]` | 
| After operation 4 | 	`[4, 4, 6, 5, 6]` | 
| After operation 5 | 	`[8, 4, 6, 5, 6]` | 

** Example 2:**  

**Input:** `nums = [1,2], k = 3, multiplier = 4`   

**Output:** `[16,8]`    
**Explanation:**  

| Operation         | Result      | 
|-------------------|-------------|
| After operation 1 | `[4, 2]`    | 
| After operation 2 | `[4, 8]`    |  
| After operation 3 | `[16, 8]`   | 

---
**Constraints:**

   * $1$ `<= nums.length <=` $100$
   * $1$ `<= nums[i] <=` $100$
   * $1$ `<= k <=` $10$
   * $1$ `<= multiplier <=` $5$
 
---
* ### Base data

```Golang
func getFinalState(nums []int, k int, multiplier int) []int {
}
```

---
### [Solution => 3264. Final Array State After K Multiplication Operations I](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3264-Final-Array-State-After-K-Multiplication-Operations-I/leetcodethreetwosixfour.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3264-Final-Array-State-After-K-Multiplication-Operations-I/leetcodethreetwosixfour.go")

---
* ### Metod => Heap/Sort & Struct
```Golang
type indexNums []*indexNum

type indexNum struct {
	Index int
	Nums  []int
}
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/3264_Final_Array_State_After_K_Multiplication_Operations_I.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/3264_Final_Array_State_After_K_Multiplication_Operations_I_Time.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/3264_Final_Array_State_After_K_Multiplication_Operations_I_Space.jpg)
