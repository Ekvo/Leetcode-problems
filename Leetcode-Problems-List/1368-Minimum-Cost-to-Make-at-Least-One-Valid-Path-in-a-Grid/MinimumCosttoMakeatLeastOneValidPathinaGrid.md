## 1368. Minimum Cost to Make at Least One Valid Path in a Grid => [original page](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/description/ "https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> |   _Hard_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
Given an `m x n` grid. Each cell of the grid has a sign pointing to the next cell you should visit if you are currently in this cell. The sign of `grid[i][j]` can be:

   * 1 which means go to the cell to the right. (i.e go from `grid[i][j]` to `grid[i][j + 1]`)
   * 2 which means go to the cell to the left. (i.e go from `grid[i][j]` to `grid[i][j - 1]`)
   * 3 which means go to the lower cell. (i.e go from `grid[i][j]` to `grid[i + 1][j]`)
   * 4 which means go to the upper cell. (i.e go from `grid[i][j]` to `grid[i - 1][j]`)

Notice that there could be some signs on the cells of the grid that point outside the grid.

You will initially start at the upper left cell `(0, 0)`. A valid path in the grid is a path that starts from the upper left cell `(0, 0)` and ends at the bottom-right cell `(m - 1, n - 1)` following the signs on the grid. The valid path does not have to be the shortest.

You can modify the sign on a cell with `cost = 1`. You can modify the sign on a cell one time only.

Return _the minimum cost to make the grid have at least **one valid path**_.

---
**Example 1:**

**Input:** `grid = [[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]]`  
**Output:** `3`  
**Explanation:** `You will start at point (0, 0).`    
`The path to (3, 3) is as follows. (0, 0) --> (0, 1) --> (0, 2) --> (0, 3) change the arrow to down with cost = 1 --> (1, 3) --> (1, 2) --> (1, 1) --> (1, 0) change the arrow to down with cost = 1 --> (2, 0) --> (2, 1) --> (2, 2) --> (2, 3) change the arrow to down with cost = 1 --> (3, 3)`  
`The total cost = 3.`  

**Example 2:**

**Input:** `grid = [[1,1,3],[3,2,2],[1,1,4]]`  
**Output:** `0`  
**Explanation:** `You can follow the path from (0, 0) to (2, 2).`  

**Example 3:**

**Input:** `grid = [[1,2],[4,3]]`  
**Output:** `1`  

---
**Constraints:**

   * `m == grid.length`
   * `n == grid[i].length`
   * $1$ `<= m, n <=` $100$
   * $1$ `<= grid[i][j] <=` $4$
 
---
* ### Base data

```Golang
func minCost(grid [][]int) int {
}
```

---
### [Solution => 1368. Minimum Cost to Make at Least One Valid Path in a Grid](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1368-Minimum-Cost-to-Make-at-Least-One-Valid-Path-in-a-Grid/leetcodeonethreesixeight.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1368-Minimum-Cost-to-Make-at-Least-One-Valid-Path-in-a-Grid/leetcodeonethreesixeight.go")

---
* ### Metod => Dijkstra's Algorithm & Heap
```Golang
type Costs []PlaceCost

type PlaceCost struct {
	cost int
	row  int
	col  int
}

func (c Costs) Len() int
func (c Costs) Swap(i, j int)
func (c Costs) Less(i, j int) bool
func (c *Costs) Push(obj any)
func (c *Costs) Pop() any

func sliceWithStartValue(slice []int, startVal int) []int
```
| Property |    Complexity     |              
|:---------|:-----------------:|
| Time     | $O(N^2 * Log(N))$ |
| Space    |     $O(N^2)$      |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1368_Minimum_Cost_to_Make_at_Least_One_Valid_Path_in_a_Grid.jpg)
