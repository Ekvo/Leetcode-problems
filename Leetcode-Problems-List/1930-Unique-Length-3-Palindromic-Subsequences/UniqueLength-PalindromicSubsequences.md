## 1930. Unique Length-3 Palindromic Subsequences => [original page](https://leetcode.com/problems/unique-length-3-palindromic-subsequences/description/ "https://leetcode.com/problems/unique-length-3-palindromic-subsequences/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given a string `s`, return _the number of **unique palindromes of length three** that are a **subsequence** of `s`_.

Note that even if there are multiple ways to obtain the same subsequence, it is still only counted **once**.

A **palindrome** is a string that reads the same forwards and backwards.

A **subsequence** of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

   * For example, `"ace"` is a subsequence of `"abcde"`.

---
**Example 1:**

**Input:** `s = "aabca"`  
**Output:** `3`  
**Explanation:** `The 3 palindromic subsequences of length 3 are:`  
`- "aba" (subsequence of "aabca")`  
`- "aaa" (subsequence of "aabca")`  
`- "aca" (subsequence of "aabca")`  

**Example 2:**

**Input:** `s = "adc"`  
**Output:** `0`  
**Explanation:** `There are no palindromic subsequences of length 3 in "adc".`  

**Example 3:**

**Input:** `s = "bbcbaba"`  
**Output:** `4`  
**Explanation:** `The 4 palindromic subsequences of length 3 are:`  
`- "bbb" (subsequence of "bbcbaba")`  
`- "bcb" (subsequence of "bbcbaba")`  
`- "bab" (subsequence of "bbcbaba")`  
`- "aba" (subsequence of "bbcbaba")`  

---
**Constraints:**

   * $3$ `<= s.length <=` $10^5$
   * `s consists of only lowercase English letters.`
 
---
* ### Base data

```Golang
func countPalindromicSubsequence(s string) int {
}
```

---
### [Solution => 1930. Unique Length-3 Palindromic Subsequences](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1930-Unique-Length-3-Palindromic-Subsequences/leetcodeoneninethreezero.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1930-Unique-Length-3-Palindromic-Subsequences/leetcodeoneninethreezero.go")

---
* ### Metod => Compute First & Last Indices
```Golang
n := 'z' - 'a' + 1
first := make([]int, n)
last := make([]int, n)

for i := 0; i < len(s); i++ {	
      // load first and last	
}

numberOfPalindromes := 0
for i := 0; i < len(first); i++{
	
   if findPalindrome {
      // increment numberOfPalindromes
   } 
}
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1930-Unique-Length-3-Palindromic-Subsequences.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1930-Unique-Length-3-Palindromic-Subsequences_Time.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1930-Unique-Length-3-Palindromic-Subsequences_Space.jpg)
