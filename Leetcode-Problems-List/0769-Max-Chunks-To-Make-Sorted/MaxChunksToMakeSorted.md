## 769. Max Chunks To Make Sorted => [original page](https://leetcode.com/problems/max-chunks-to-make-sorted/description/ "https://leetcode.com/problems/max-chunks-to-make-sorted/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You are given an integer array `arr` of length `n` that represents a permutation of the integers in the range `[0, n - 1]`.

We split `arr` into some number of **chunks** (i.e., partitions), and individually sort each chunk. After concatenating them, the result should equal the sorted array.

Return _the largest number of chunks we can make to sort the array_.

---
**Example 1:**

**Input:** `arr = [4,3,2,1,0]`  
**Output:** `1`  
**Explanation:**  
`Splitting into two or more chunks will not return the required result.`  
`For example, splitting into [4, 3], [2, 1, 0] will result in [3, 4, 0, 1, 2], which isn't sorted.`  

**Example 2:**

**Input:** `arr = [1,0,2,3,4]`  
**Output:** `4`  
**Explanation:**  
`We can split into two chunks, such as [1, 0], [2, 3, 4].`  
`However, splitting into [1, 0], [2], [3], [4] is the highest number of chunks possible.`  

---
**Constraints:**

   * `n == arr.length`
   * $1$ `<= n <=` $10$
   * $0$ `<= arr[i] <` $n$
   * All the elements of `arr` are **unique**.
 
---
* ### Base data

```Golang
func maxChunksToSorted(arr []int) int {
}
```

---
### [Solution => 769. Max Chunks To Make Sorted](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0769-Max-Chunks-To-Make-Sorted/leetcodesevensixnine.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0769-Max-Chunks-To-Make-Sorted/leetcodesevensixnine.go")

---
* ### Metod => Prefix Sums
```Golang
indexSum       = 0
integerSum     = 0
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/769_Max_Chunks_To_Make_Sorted_Prefix_Sum.jpg)

---
* ### Metod => Stack
```Golang
type stack []int

func (s *stack) Empty() bool
func (s *stack) Push(i int)
func (s *stack) Pop()
func (s *stack) Top() int
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/769_Max_Chunks_To_Make_Sorted_Stack.jpg)
