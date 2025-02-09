## 1726. Tuple with Same Product => [original page](https://leetcode.com/problems/tuple-with-same-product/description/ "https://leetcode.com/problems/tuple-with-same-product/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> | _Medium_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
Given an array `nums` of **distinct** positive integers, return _the number of tuples `(a, b, c, d)` such that `a * b = c * d` where `a, b, c`, and `d` are elements of `nums`, and `a != b != c != d`_.

---
**Example 1:**

**Input:** `nums = [2,3,4,6]`  
**Output:** `8`  
**Explanation:** `There are 8 valid tuples:`  
`(2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)`  
`(3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)`  

**Example 2:**

**Input:** `nums = [1,2,4,5,10]`  
**Output:** `16`  
**Explanation:** `There are 16 valid tuples:`  
`(1,10,2,5) , (1,10,5,2) , (10,1,2,5) , (10,1,5,2)`  
`(2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1)`  
`(2,10,4,5) , (2,10,5,4) , (10,2,4,5) , (10,2,5,4)`  
`(4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2)`  

---
**Constraints:**

   * $1$ `<= nums.length <=` $1000$
   * $1$ `<= nums[i] <=` $10^4$
   * `All elements in nums are distinct.`
 
---
* ### Base data

```Golang
func tupleSameProduct(nums []int) int {
}
```

---
### [Solution => 1726. Tuple with Same Product](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1726-Tuple-with-Same-Product/leetcodeoneseventwosix.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1726-Tuple-with-Same-Product/leetcodeoneseventwosix.go")

---
* ### Metod => HashMap & Triangular Numbers 'n * (n - 1) / 2'
```Golang
multiplyTwoNums map[int]int

for i := n - 1; i > 0; i-- {

   for j := i - 1; j > -1; j-- {
     //add nums[i]*nums[j] to multiplyTwoNums
   }
}

for _, count := range multiplyTwoNums {
   // if count>1 -> add result	
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |  $O(N^2)$  |
| Space    |   $O(N)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1726_Tuple_with_Same_Product.jpg)
