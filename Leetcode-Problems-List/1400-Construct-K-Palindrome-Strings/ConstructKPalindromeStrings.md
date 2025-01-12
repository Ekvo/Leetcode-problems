## 1400. Construct K Palindrome Strings => [original page](https://leetcode.com/problems/construct-k-palindrome-strings/description/ "https://leetcode.com/problems/construct-k-palindrome-strings/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> | _Medium_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
Given a string `s` and an integer `k`, return `true` _if you can use all the characters in `s` to construct `k` palindrome strings or `false` otherwise_.

---
**Example 1:**

**Input:** `s = "annabelle", k = 2`  
**Output:** true`  
**Explanation:** `You can construct two palindromes using all characters in s.`  
`Some possible constructions "anna" + "elble", "anbna" + "elle", "anellena" + "b"`  

**Example 2:**

**Input:** `s = "leetcode", k = 3`  
**Output:** `false`
**Explanation:** `It is impossible to construct 3 palindromes using all the characters of s.`  

**Example 3:**

**Input:** `s = "true", k = 4`  
**Output:** `true`
**Explanation:** `The only possible solution is to put each character in a separate string.`  

---
**Constraints:**

   * $1$ `<= s.length <=` $10^5$
   * `s consists of lowercase English letters.`
   * $1$ `<= k <=` $10^5$
 
---
* ### Base data

```Golang
func canConstruct(s string, k int) bool {
}
```

---
### [Solution => 1400. Construct K Palindrome Strings](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1400-Construct-K-Palindrome-Strings/leetcodeonefourzerozero.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1400-Construct-K-Palindrome-Strings/leetcodeonefourzerozero.go")

---
* ### Metod => Count Odd Frequencies
```Golang
for _, ch := range s {
     //find count of each letter
}

for _, num := range letters {
     //find count odd letters 	
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |   $O(N)$   |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1400_Construct_K_Palindrome_Strings.jpg)
