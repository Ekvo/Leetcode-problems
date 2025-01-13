## 32. Longest Valid Parentheses => [original page](https://leetcode.com/problems/longest-valid-parentheses/description/ "https://leetcode.com/problems/longest-valid-parentheses/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> |   _Hard_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
Given a string containing just the characters `'('` and `')'`, return _the length of the longest valid (well-formed) parentheses substring_.

---
**Example 1:**

**Input:** `s = "(()"`  
**Output:** 2`  
**Explanation:** `The longest valid parentheses substring is "()".`  

**Example 2:**

**Input:** `s = ")()())"`  
**Output:** `4`  
Explanation: The longest valid parentheses substring is "()()".`  

**Example 3:**

**Input:** `s = ""`  
**Output:** `0`  

---
**Constraints:**

   * $0$ `<= s.length <=` $3 * 10^4$
   * `s[i] is '(', or ')'.`
 
---
* ### Base data

```Golang
func longestValidParentheses(s string) int {
}
```

---
### [Solution => 32. Longest Valid Parentheses]( "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0032-Longest-Valid-Parentheses/leetcodethreetwo.go")

---
* ### Metod => Cycle right & Cycle left
```Golang
for i := 0; i < len(s); i++ {
    // find valid Parentheses
}

for i := len(s)-1; i > -1; i-- {
    // find valid Parentheses
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/32_Longest_Valid_Parentheses.jpg)
