## 827. Making A Large Island => [original page](https://leetcode.com/problems/making-a-large-island/description/ "https://leetcode.com/problems/making-a-large-island/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> |   _Hard_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
You are given an `n x n` binary matrix `grid`. You are allowed to change **at most one** `0` to be `1`.

Return _the size of the largest **island** in `grid` after applying this operation_.

An **island** is a 4-directionally connected group of `1`s.

---
**Example 1:**

**Input:** `grid = [[1,0],[0,1]]`  
**Output:** `3`  
**Explanation:** `Change one 0 to 1 and connect two 1s, then we get an island with area = 3.`  

**Example 2:**

**Input:** `grid = [[1,1],[1,0]]`  
**Output:** `4`  
**Explanation:** `Change the 0 to 1 and make the island bigger, only one island with area = 4.`  

**Example 3:**

**Input:** `grid = [[1,1],[1,1]]`  
**Output:** `4`  
**Explanation:** `Can't change any 0 to 1, only one island with area = 4.` 

---
**Constraints:**

   * `n == grid.length`
   * `n == grid[i].length`
   * $1$ `<= n <=` $500$
   * `grid[i][j] is either 0 or 1.`
 
---
* ### Base data

```Golang
func largestIsland(grid [][]int) int {
}
```

---
### [Solution => 827. Making A Large Island](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0827-Making-A-Large-Island/leetcodeeighttwoseven.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0827-Making-A-Large-Island/leetcodeeighttwoseven.go")

---
* ### Metod => BFS label each island
```Golang
islandSize map[int]int // key - islandID ;value size of island

createIsland func(r, c int)
newIslandWithSize func(r, c int) int

for r := 0; r < rows; r++ {
	
   for c := 0; c < cols; c++ {
      // if find '1' - label island 
   }
}

for r := 0; r < rows; r++ {
	
   for c := 0; c < cols; c++ {
      // if find '0' - find size of new iland
   }
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |  $O(N^2)$  |
| Space    |  $O(N^2)$  |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/827_Making_A_Large_Island.jpg)
