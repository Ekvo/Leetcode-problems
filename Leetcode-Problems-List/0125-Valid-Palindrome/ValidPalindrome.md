## 125. Valid Palindrome => [original page ](https://leetcode.com/problems/valid-palindrome/description/ "https://leetcode.com/problems/valid-palindrome/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Easy_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string `s`, return `true` if it is a **palindrome**, or `false` otherwise.

---
**Example 1:**

**Input:** `s = "A man, a plan, a canal: Panama"`  
**Output:** `true`  
**Explanation:** `"amanaplanacanalpanama" is a palindrome.`  

**Example 2:**

**Input:** `s = "race a car"`  
**Output:** `false`  
**Explanation:** `"raceacar" is not a palindrome.`  

**Example 3:**

**Input:** `s = " "`  
**Output:** `true`  
**Explanation:** `s is an empty string "" after removing non-alphanumeric characters.`  
`Since an empty string reads the same forward and backward, it is a palindrome.`

---
**Constraints:**

   * $1$ `<= s.length <=` $2 * 10^5$  
   * `s consists only of printable ASCII characters.`

---
* ### Base data

```Golang
func isPalindrome(s string) bool {
}
```

---
### [Solution => 125. Valid Palindrome](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0125-Valid-Palindrome/leetcodeonetwofive.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0125-Valid-Palindrome/leetcodeonetwofive.go")

---
* ### Metod => Head >> Center << Tail
```Golang
var (
	head   = 0
	tail   = len(s) - 1
	chHead rune
	chTail rune
)
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/125_Valid_Palindrome.jpg)
 
