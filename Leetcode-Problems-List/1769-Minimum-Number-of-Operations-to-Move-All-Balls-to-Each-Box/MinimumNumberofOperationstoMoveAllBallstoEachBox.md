## 1769. Minimum Number of Operations to Move All Balls to Each Box => [original page](https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/description/ "https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You have `n` boxes. You are given a binary string `boxes` of length `n`, where `boxes[i]` is `'0'` if the $i^{th}$ box is empty, and `'1'` if it contains one ball.

In one operation, you can move **one** ball from a box to an adjacent box. Box `i` is adjacent to box `j` if `abs(i - j) == 1`. Note that after doing so, there may be more than one ball in some boxes.

Return an array `answer` of size `n`, where `answer[i]` is the **minimum** number of operations needed to move all the balls to the $i^{th}$ box.

Each `answer[i]` is calculated considering the **initial** state of the boxes.


---
**Example 1:**

**Input:** `boxes = "110"`  
**Output:** `[1,1,3]`  
**Explanation:** `The answer for each box is as follows:`  
`1) First box: you will have to move one ball from the second box to the first box in one operation.`  
`2) Second box: you will have to move one ball from the first box to the second box in one operation.`  
`3) Third box: you will have to move one ball from the first box to the third box in two operations, and move one ball from the second box to the third box in one operation.`  

**Example 2:**

**Input:** `boxes = "001011"`  
**Output:** `[11,8,5,4,3,4]`  

---
**Constraints:**

   * `n == boxes.length`
   * $1$ `<= n <=` $2000$
   * `boxes[i] is either '0' or '1'.`
 
---
* ### Base data

```Golang
func minOperations(boxes string) []int {	
}
```

---
### [Solution => 1769. Minimum Number of Operations to Move All Balls to Each Box](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1769-Minimum-Number-of-Operations-to-Move-All-Balls-to-Each-Box/leetcodeonesevensixnine.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1769-Minimum-Number-of-Operations-to-Move-All-Balls-to-Each-Box/leetcodeonesevensixnine.go")

---
* ### Metod => Sum of Left and Right Moves
```Golang
// head and tail
for h, t := 0, n-1; h < n; h, t = h+1, t-1 {     	
      answer[h] += moveLeft
      //find moveLeft		

      answer[t] += moveRight
      //find moveRight			
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1769_Minimum_Number_of_Operations_to_Move_All_Balls_to_Each_Box_Sum_Left_Right.jpg)

---
* ### Metod => Brute Force
```Golang
onesIndex := make([]int, 0, len(boxes))

for i, ch := range boxes {
      //find index with '1'	
}

for i, _ := range boxes {
	
    for _, index := range onesIndex {
      //find  answer[i]
    }
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |  $O(N^2)$  |
| Space    |   $O(N)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1769_Minimum_Number_of_Operations_to_Move_All_Balls_to_Each_Box_Brute_Force.jpg)

