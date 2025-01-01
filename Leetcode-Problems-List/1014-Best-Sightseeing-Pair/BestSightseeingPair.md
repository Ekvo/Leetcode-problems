## 1014. Best Sightseeing Pair => [original page](https://leetcode.com/problems/best-sightseeing-pair/description/ "https://leetcode.com/problems/best-sightseeing-pair/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You are given an integer array `values` where `values[i]` represents the value of the $i^{th}$ sightseeing spot. Two sightseeing spots `i` and `j` have a distance `j - i` between them.

The score of a pair (`i < j`) of sightseeing spots is `values[i] + values[j] + i - j`: the sum of the values of the sightseeing spots, minus the distance between them.

Return _the maximum score of a pair of sightseeing spots._

---
**Example 1:**

**Input:** `values = [8,1,5,2,6]`  
**Output:** `11`  
**Explanation:** `i = 0, j = 2, values[i] + values[j] + i - j = 8 + 5 + 0 - 2 = 11`  

**Example 2:**

**Input:** `values = [1,2]`  
**Output:** `2`  

---
**Constraints:**

   * $2$ `<= values.length <=` $5 * 10^4$
   * $1$ `<= values[i] <=` $1000$ 

---
* ### Base data

```Golang
func maxScoreSightseeingPair(values []int) int {
}
```

---
### [Solution = 1014. Best Sightseeing Pair](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1014-Best-Sightseeing-Pair/leetcodeonezeroonefour.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1014-Best-Sightseeing-Pair/leetcodeonezeroonefour.go")

---
* ### Metod => Dynamic Programming
```Golang
var (
   maxScore 
	
   maxLeftScore
   curRightScore
   curLeftScore
)
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1014_Best_Sightseeing_Pair.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1014_Best_Sightseeing_Pair_Time.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1014_Best_Sightseeing_Pair_Space.jpg)
