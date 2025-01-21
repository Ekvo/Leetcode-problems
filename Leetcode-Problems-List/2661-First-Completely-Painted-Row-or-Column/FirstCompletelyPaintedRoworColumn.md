## 2661. First Completely Painted Row or Column => [original page](https://leetcode.com/problems/first-completely-painted-row-or-column/description/ "https://leetcode.com/problems/first-completely-painted-row-or-column/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> | _Medium_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
You are given a **0-indexed** integer array `arr`, and an `m x n` integer **matrix** `mat`. `arr` and `mat` both contain **all** the integers in the range `[1, m * n]`.

Go through each index `i` in `arr` starting from index `0` and paint the cell in mat containing the integer `arr[i]`.

Return _the smallest index `i` at which either a row or a column will be completely painted in `mat`_.

---
**Example 1:**

**Input:** `arr = [1,3,4,2], mat = [[1,4],[2,3]]`  
**Output:** `2`  
**Explanation:** `The moves are shown in order, and both the first row and second column of the matrix become fully painted at arr[2].`  

**Example 2:**

**Input:** `arr = [2,8,7,4,1,3,5,6,9], mat = [[3,2,5],[1,4,6],[8,7,9]]`  
**Output:** `3`  
**Explanation:** `The second column becomes fully painted at arr[3].`  

---
**Constraints:**

   * `m == mat.length`
   * `n = mat[i].length`
   * `arr.length == m * n`
   * $1$ `<= m, n <=` $10^5$
   * $1$ `<= m * n <=` $10^5$
   * $1$ `<= arr[i], mat[r][c] <=` $m * n$
   * `All the integers of arr are unique.`
   * `All the integers of mat are unique.`
 
---
* ### Base data

```Golang
func firstCompleteIndex(arr []int, mat [][]int) int {
}
```

---
### [Solution => 2661. First Completely Painted Row or Column](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2661-First-Completely-Painted-Row-or-Column/leetcodetwosixsixone.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2661-First-Completely-Painted-Row-or-Column/leetcodetwosixsixone.go")

---
* ### Metod => Brute Force & HashMap
```Golang
valAndPoint map[int][]int

arrRows []int
arrCols []int
```
| Property | Complexity |              
|:---------|:----------:|
| Time     | $O(N * M)$ |
| Space    | $O(N * M)$ |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2661_First_Completely_Painted_Row_or_Column.jpg)
