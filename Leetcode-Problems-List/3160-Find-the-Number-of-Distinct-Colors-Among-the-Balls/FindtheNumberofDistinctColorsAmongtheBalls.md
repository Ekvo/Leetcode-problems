## 3160. Find the Number of Distinct Colors Among the Balls => [original page](https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/description/ "https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> | _Medium_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
You are given an integer `limit` and a 2D array queries of size `n x 2`.

There are `limit + 1` balls with **distinct** labels in the range `[0, limit]`. Initially, all balls are uncolored. For every query in `queries` that is of the form `[x, y]`, you mark ball `x` with the color `y`. After each query, you need to find the number of **distinct** colors among the balls.

Return an array `result` of length `n`, where `result[i]` denotes the number of distinct colors _after_ $i^{th}$ query.

**Note** that when answering a query, lack of a color _will not_ be considered as a color.

---
**Example 1:**

**Input:** `limit = 4, queries = [[1,4],[2,5],[1,3],[3,4]]`  
**Output:** `[1,2,2,3]`  
**Explanation:**  
`    After query 0, ball 1 has color 4.`  
`    After query 1, ball 1 has color 4, and ball 2 has color 5.`  
`    After query 2, ball 1 has color 3, and ball 2 has color 5.`  
`    After query 3, ball 1 has color 3, ball 2 has color 5, and ball 3 has color 4.`  

**Example 2:**

**Input:** `limit = 4, queries = [[0,1],[1,2],[2,2],[3,4],[4,5]]`  
**Output:** `[1,2,2,3,4]`  
**Explanation:**  
`    After query 0, ball 0 has color 1.`  
`    After query 1, ball 0 has color 1, and ball 1 has color 2.`  
`    After query 2, ball 0 has color 1, and balls 1 and 2 have color 2.`  
`    After query 3, ball 0 has color 1, balls 1 and 2 have color 2, and ball 3 has color 4.`  
`    After query 4, ball 0 has color 1, balls 1 and 2 have color 2, ball 3 has color 4, and ball 4 has color 5.`  

---
**Constraints:**

   * $1$ `<= limit <=` $10^9$
   * $1$ `<= n == queries.length <=` $10^5$
   * `queries[i].length == 2`
   * $0$ `<= queries[i][0] <=` $limit$
   * $1$ `<= queries[i][1] <=` $10^9$
 
---
* ### Base data

```Golang
func queryResults(limit int, queries [][]int) []int {
}
```

---
### [Solution => 3160. Find the Number of Distinct Colors Among the Balls](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3160-Find-the-Number-of-Distinct-Colors-Among-the-Balls/leetcodethreeonesixzero.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/3160-Find-the-Number-of-Distinct-Colors-Among-the-Balls/leetcodethreeonesixzero.go")

---
* ### Metod =>  Two HashMap 
```Golang
ballColor  map[int]int
colorCount map[int]int

for _,q:=range queries{
   // ballColor - add ball 'q[0]' equal color 'q[1]'
   // colorCount - increment color 'q[1]'
	
   // add len('colorCount') to 'result'	
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(N)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/3160_Find_the_Number_of_Distinct_Colors_Among_the_Balls.jpg)
