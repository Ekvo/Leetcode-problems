## 2558. Take Gifts From the Richest Pile => [original page](https://leetcode.com/problems/take-gifts-from-the-richest-pile/description/ "https://leetcode.com/problems/take-gifts-from-the-richest-pile/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Easy_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You are given an integer array `gifts` denoting the number of gifts in various piles. Every second, you do the following:

* Choose the pile with the maximum number of gifts.
* If there is more than one pile with the maximum number of gifts, choose any.
* Leave behind the floor of the square root of the number of gifts in the pile. Take the rest of the gifts.

Return _the number of gifts remaining after `k` seconds._

---
**Example 1:**

**Input:** `gifts = [25,64,9,4,100], k = 4`  
**Output:** `29`  
**Explanation:**  
`The gifts are taken in the following way:`  
`- In the first second, the last pile is chosen and 10 gifts are left behind.`  
`- Then the second pile is chosen and 8 gifts are left behind.`  
`- After that the first pile is chosen and 5 gifts are left behind.`  
`- Finally, the last pile is chosen again and 3 gifts are left behind.`  
`The final remaining gifts are [5,8,9,4,3], so the total number of gifts remaining is 29.`

**Example 2:**

**Input:** `gifts = [1,1,1,1], k = 4`  
**Output:** `4`  
**Explanation:**
`In this case, regardless which pile you choose, you have to leave behind 1 gift in each pile.`  
`That is, you can't take any pile with you.`  
`So, the total gifts remaining are 4.`

---
**Constraints:**

* $1$ `<= gifts.length <=` $10^3$
* $1$ `<= gifts[i] <=` $10^9$
* $1$ `<= k <=` $10^3$

---
* ### Base data

```Golang
func pickGifts(gifts []int, k int) int64 {
}
```

---
### [Solution => 2558. Take Gifts From the Richest Pile](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2558-Take-Gifts-From-the-Richest-Pile/leetcodetwofivefiveeight.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2558-Take-Gifts-From-the-Richest-Pile/leetcodetwofivefiveeight.go")

---
* ### Metod => Heap/Sort
```Golang
//heap.Fix(&data,0)
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2558_Take_Gifts_From_the_Richest_Pile.jpg)
![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2558_Take_Gifts_From_the_Richest_Pile_Time_Complexity.jpg)
![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2558_Take_Gifts_From_the_Richest_Pile_Space_Complexity.jpg)
