## 9. Palindrome Number => [original page](https://leetcode.com/problems/palindrome-number/description/ "https://leetcode.com/problems/palindrome-number/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Easy_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given an integer `x`, return `true` if `x` is a
_palindrome_, and `false` otherwise.


---
**Example 1:**

**Input:** `x = 121`  
**Output:** `true`  
**Explanation:** `121 reads as 121 from left to right and from right to left.`  

**Example 2:**

**Input:** `x = -121`  
**Output:** `false`  
**Explanation:** `From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.`  

**Example 3:**

**Input:** `x = 10`  
**Output:** `false`  
**Explanation:** `Reads 01 from right to left. Therefore it is not a palindrome.`  

---
**Constraints:**

 * $-2^{31}$ `<= x <=` $2^{31} - 1$

---
* ### Base data

```Golang
func isPalindrome(x int) bool {
}
```

---
### [Solution => 9. Palindrome Number](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0009-Palindrome-Number/leetcodenine.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0009-Palindrome-Number/leetcodenine.go")

---
* ### Metod => Two Pointers & strconv.Itoa(integer)
```Golang
for head < tail {
      //work
}
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/9_Palindrome_Number_String.jpg)

* ### Metod => Digit ('%10' & '/10' & '*10')
```Golang 
for x > 0 {
   digit = x % 10
   x /= 10
   number = number * 10 + digit
}
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/9_Palindrome_Number_Digit.jpg)
 
