## 2658. Maximum Number of Fish in a Grid => [original page](https://leetcode.com/problems/maximum-number-of-fish-in-a-grid/description/ "https://leetcode.com/problems/maximum-number-of-fish-in-a-grid/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> | _Medium_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
You are given a **0-indexed** 2D matrix `grid` of size `m x n`, where `(r, c)` represents:

   * A **land** cell if `grid[r][c] = 0`, or
   * A **water** cell containing `grid[r][c]` fish, if `grid[r][c] > 0`.

A fisher can start at any **water** cell `(r, c)` and can do the following operations any number of times:

   * Catch all the fish at cell `(r, c)`, or
   * Move to any adjacent **water** cell.

Return _the **maximum** number of fish the fisher can catch if he chooses his starting cell optimally, or `0` if no water cell exists_.

An **adjacent** cell of the cell `(r, c)`, is one of the cells `(r, c + 1)`, `(r, c - 1)`, `(r + 1, c)` or `(r - 1, c)` if it exists.

---
**Example 1:**

**Input:** `grid = [[0,2,1,0],[4,0,0,3],[1,0,0,4],[0,3,2,0]]`  
**Output:** `7`  
**Explanation:** `The fisher can start at cell (1,3) and collect 3 fish, then move to cell (2,3) and collect 4 fish.`  

**Example 2:**

**Input:** `1grid = [[1,0,0,0],[0,0,0,0],[0,0,0,0],[0,0,0,1]]`  
**Output:** `1`  
**Explanation:** `The fisher can start at cells (0,0) or (3,3) and collect a single fish.`  

---
**Constraints:**

   * `m == grid.length`
   * `n == grid[i].length`
   * $1$ `<= m, n <=` $10$
   * $0$ `<= grid[i][j] <=` $10$
 
---
* ### Base data

```Golang
func findMaxFish(grid [][]int) int {
}
```

---
### [Solution => 2658. Maximum Number of Fish in a Grid](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2658-Maximum-Number-of-Fish-in-a-Grid/leetcodetwosixfiveeight.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2658-Maximum-Number-of-Fish-in-a-Grid/leetcodetwosixfiveeight.go")

---
* ### Metod => Queue & BFS 
```Golang
// getFishs - find count of fish in one pond
var getFishs func(row, col int) int

for r := 0; r < rows; r++ {
	
   for c := 0; c < cols; c++ {
      // if find fish 
      // let's start counting the number
   }	
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     | $O(M * N)$ |
| Space    | $O(M * N)$ |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2658_Maximum_Number_of_Fish_in_a_Grid.jpg)
