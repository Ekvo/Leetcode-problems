## 2593. Find Score of an Array After Marking All Elements => [original page](https://leetcode.com/problems/find-score-of-an-array-after-marking-all-elements/description/ "https://leetcode.com/problems/find-score-of-an-array-after-marking-all-elements/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You are given an array `nums` consisting of positive integers.

Starting with `score = 0`, apply the following algorithm:

   * Choose the smallest integer of the array that is not marked. If there is a tie, choose the one with the smallest index.
   * Add the value of the chosen integer to score.
   * Mark **the chosen element and its two adjacent elements if they exist.**
   * Repeat until all the array elements are marked.

Return _the score you get after applying the above algorithm._

---
**Example 1:**

**Input:** `nums = [2,1,3,4,5,2]`  
**Output:** `7`  
**Explanation:** `We mark the elements as follows:`  
`- 1 is the smallest unmarked element, so we mark it and its two adjacent elements: [2,1,3,4,5,2].`  
`- 2 is the smallest unmarked element, so we mark it and its left adjacent element: [2,1,3,4,5,2].`  
`- 4 is the only remaining unmarked element, so we mark it: [2,1,3,4,5,2].`
` Our score is 1 + 2 + 4 = 7.`

**Example 2:**

**Input:** `nums = [2,3,5,1,3,2]`  
**Output:** `5`  
**Explanation:** `We mark the elements as follows:`  
`- 1 is the smallest unmarked element, so we mark it and its two adjacent elements: [2,3,5,1,3,2].`  
`- 2 is the smallest unmarked element, since there are two of them, we choose the left-most one, so we mark the one at index 0 and its right adjacent element: [2,3,5,1,3,2].`  
`- 2 is the only remaining unmarked element, so we mark it: [2,3,5,1,3,2].`  
`  Our score is 1 + 2 + 2 = 5.` 

---
**Constraints:**

   * $1$ `<= nums.length <=` $10^5$
   * $1$ `<= nums[i] <=` $10^6$

 
---
* ### Base data

```Golang
func findScore(nums []int) int64 {
}
```

---
### [Solution => 2593. Find Score of an Array After Marking All Elements](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2593_Find_Score_of_an_Array_After_Marking_All_Elements/leetcodetwofiveninethree.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2593_Find_Score_of_an_Array_After_Marking_All_Elements/leetcodetwofiveninethree.go")

---
* ### Metod => Struct & sort.Slice
```Golang
type indexValue struct {
	index int
	value int
}
store    []*indexValue
visited  []bool
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2593_Find_Score_of_an_Array_After_Marking_All_Elements.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2593_Find_Score_of_an_Array_After_Marking_All_Elements_Time.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2593_Find_Score_of_an_Array_After_Marking_All_Elements_Space.jpg)
