## 2017. Grid Game => [original page](https://leetcode.com/problems/grid-game/description/ "https://leetcode.com/problems/grid-game/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> | _Medium_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
You are given a **0-indexed** 2D array `grid` of size `2 x n`, where `grid[r][c]` represents the number of points at position `(r, c)` on the matrix. Two robots are playing a game on this matrix.

Both robots initially start at `(0, 0)` and want to reach `(1, n-1)`. Each robot may only move to the **right** (`(r, c)` to `(r, c + 1)`) or **down** (`(r, c)` to `(r + 1, c)`).

At the start of the game, the **first** robot moves from `(0, 0)` to `(1, n-1)`, collecting all the points from the cells on its path. For all cells `(r, c)` traversed on the path, `grid[r][c]` is set to `0`. Then, the **second** robot moves from `(0, 0)` to `(1, n-1)`, collecting the points on its path. Note that their paths may intersect with one another.

The **first** robot wants to **minimize** the number of points collected by the **second** robot. In contrast, the **second** robot wants to **maximize** the number of points it collects. If both robots play **optimally**, return _the **number of points** collected by the **second** robot_.

---
**Example 1:**

**Input:** `grid = [[2,5,4],[1,5,1]]`  
**Output:** `4`  
**Explanation:** `The optimal path taken by the first robot is shown in red, and the optimal path taken by the second robot is shown in blue.`  
`The cells visited by the first robot are set to 0.`  
`The second robot will collect 0 + 0 + 4 + 0 = 4 points.`  

**Example 2:**

**Input:** `grid = [[3,3,1],[8,5,2]]`  
**Output:** `4`  
**Explanation:** `The optimal path taken by the first robot is shown in red, and the optimal path taken by the second robot is shown in blue.`  
`The cells visited by the first robot are set to 0.`  
`The second robot will collect 0 + 3 + 1 + 0 = 4 points.`  

**Example 3:**

**Input:** `grid = [[1,3,1,15],[1,3,3,1]]`  
**Output:** `7`  
**Explanation:** `The optimal path taken by the first robot is shown in red, and the optimal path taken by the second robot is shown in blue.`  
`The cells visited by the first robot are set to 0.`  
`The second robot will collect 0 + 1 + 3 + 3 + 0 = 7 points.`  

---
**Constraints:**

   * `grid.length == 2`
   * `n == grid[r].length`
   * $1$ `<= n <=` $5 * 10^4$
   * $1$ `<= grid[r][c] <=` $10^5$
 
---
* ### Base data

```Golang
func gridGame(grid [][]int) int64 {	
}
```

---
### [Solution => 2017. Grid Game](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2017-Grid-Game/leetcodetwozerooneseven.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2017-Grid-Game/leetcodetwozerooneseven.go")

---
* ### Metod => Prefix and Suffix Sum
```Golang
for i := 0; i < len(grid[0]); i++ {
    //find sum grid[0] - Suffix Sum
}

for i := 0; i < len(grid[0]); i++ {
    //decrement Suffix Sum
	
    //compare Prefix and Suffix Sum
	
    //increment Prefix Sum	
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2017_Grid_Game.jpg)
