## 1790. Check if One String Swap Can Make Strings Equal => [original page](https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/description/ "https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> |   _Easy_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
You are given two strings `s1` and `s2` of equal length. A **string swap** is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.

Return _`true` if it is possible to make both strings equal by performing **at most one string swap** on **exactly one** of the strings. Otherwise, return `false`_.

---
**Example 1:**

**Input:** `s1 = "bank", s2 = "kanb"`  
**Output:** `true`  
**Explanation:** `For example, swap the first character with the last character of s2 to make "bank".`  

**Example 2:**

**Input:** `s1 = "attack", s2 = "defend"`  
**Output:** `false`  
**Explanation:** `It is impossible to make them equal with one string swap.`  

**Example 3:**

**Input:** `s1 = "kelb", s2 = "kelb"`  
**Output:** `true`  
**Explanation:** `The two strings are already equal, so no string swap operation is required.`  

---
**Constraints:**

  * $1$ `<= s1.length, s2.length <=` $100$
  * `s1.length == s2.length`
  * `s1 and s2 consist of only lowercase English letters.`
 
---
* ### Base data

```Golang
func areAlmostEqual(s1 string, s2 string) bool {
}
```

---
### [Solution => 1790. Check if One String Swap Can Make Strings Equal](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1790-Check-if-One-String-Swap-Can-Make-Strings-Equal/leetcodeonesevenninezero.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1790-Check-if-One-String-Swap-Can-Make-Strings-Equal/leetcodeonesevenninezero.go")

---
* ### Metod => Count Difference
```Golang
for i := 0; i < len(s1); i++ {
   // if s1[i] not equal s2[i] 
   // increment 'count Difference'	
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1790_Check_if_One_String_Swap_Can_Make_Strings_Equal.jpg)
