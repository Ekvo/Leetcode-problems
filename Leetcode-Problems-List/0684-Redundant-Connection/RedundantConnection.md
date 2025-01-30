## 684. Redundant Connection => [original page](https://leetcode.com/problems/redundant-connection/description/ "https://leetcode.com/problems/redundant-connection/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> | _Medium_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
In this problem, a tree is an **undirected graph** that is connected and has no cycles.

You are given a graph that started as a tree with `n` nodes labeled from `1` to `n`, with one additional edge added. The added edge has two **different** vertices chosen from `1` to `n`, and was not an edge that already existed. The graph is represented as an array `edges` of length n where `edges[i] = [ai, bi]` indicates that there is an edge between nodes $a_i$ and $b_i$ in the graph.

Return _an edge that can be removed so that the resulting graph is a tree of `n` nodes_. If there are multiple answers, return the answer that occurs last in the input.

---
**Example 1:**

**Input:** `edges = [[1,2],[1,3],[2,3]]`  
**Output:** `[2,3]`  

**Example 2:**

**Input:** `edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]`  
**Output:** `[1,4]`  

---
**Constraints:**

   * `n == edges.length`
   * $3$ `<= n <=` $1000$
   * `edges[i].length == 2`
   * $1$ `<= ai < bi <=` $edges.length$
   * $a_i$ != $b_i$
   * `There are no repeated edges.`
   * `The given graph is connected.`


 
---
* ### Base data

```Golang
func findRedundantConnection(edges [][]int) []int {
}
```

---
### [Solution => 684. Redundant Connection](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0684-Redundant-Connection/leetcodesixeightfour.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0684-Redundant-Connection/leetcodesixeightfour.go")

---
* ### Metod => DFS
```Golang
var getParent func(b int) int
// get and compare parents 'a' and 'b'
var equalParents func(a, b int) bool

for _, edge := range edges {	
	
   if equalParents(edge[0], edge[1]){
      // get result and return		
   }
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(N)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/684_Redundant_Connection.jpg)
