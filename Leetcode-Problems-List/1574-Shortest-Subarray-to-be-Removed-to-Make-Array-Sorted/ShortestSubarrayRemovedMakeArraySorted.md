## 1574. Shortest Subarray to be Removed to Make Array Sorted => [original page](https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted/description/ "https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given an integer array `arr`, remove a subarray (can be empty) from `arr` such that the remaining elements in `arr` **are non-decreasing**.

Return the length of the shortest subarray to remove.

A **subarray** is a contiguous subsequence of the array.


---
**Example 1:**

**Input:** `arr = [1,2,3,10,4,2,3,5]`  
**Output:** `3`  
**Explanation:** `The shortest subarray we can remove is [10,4,2] of length 3. The remaining elements after that will be [1,2,3,3,5] which are sorted.`  
`Another correct solution is to remove the subarray [3,10,4].`

**Example 2:**

**Input:** `arr = [5,4,3,2,1]`  
**Output:** `4`  
**Explanation:** `Since the array is strictly decreasing, we can only keep a single element. Therefore we need to remove a subarray of length 4, either [5,4,3,2] or [4,3,2,1].`  

**Example 3:**

**Input:** `arr = [1,2,3]`  
**Output:** `0`  
**Explanation:** `The array is already non-decreasing. We do not need to remove any elements.`

---
**Constraints:**
   * $1$ `<= arr.length <=` $10^5$  
   * $0$ `<= arr[i] <=` $10^9$  

---
* ### Base data

```Golang
func findLengthOfShortestSubarray(arr []int) int {
}
```

---
### [Solution => 1574. Shortest Subarray to be Removed to Make Array Sorted](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1574-Shortest-Subarray-to-be-Removed-to-Make-Array-Sorted/shortestSubarrayRemovedMakeArraySorted.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1574-Shortest-Subarray-to-be-Removed-to-Make-Array-Sorted/shortestSubarrayRemovedMakeArraySorted.go")

---
* ### Metod => Two Pointers
```Golang
var (
    head      = 1
    tail      = len(arr) - 1
    minSubArr = 0
)
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1574_Shortest_Subarray_to_be_Removed_to_Make_Array_Sorted.jpg)
 
