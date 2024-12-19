## 3. Longest Substring Without Repeating Characters => [original page](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/ "https://leetcode.com/problems/longest-substring-without-repeating-characters/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> | _Medium_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given a string `s`, find the length of the **longest** substring without repeating characters.

---
**Example 1:**

**Input:** `s = "abcabcbb"`  
**Output:** `3`  
**Explanation:** `The answer is "abc", with the length of 3.`  

**Example 2:**

**Input:** `s = "bbbbb"`  
**Output:** `1`  
**Explanation:** `The answer is "b", with the length of 1.`

**Example 3:**

**Input:** `s = "pwwkew"`  
**Output:** `3`  
**Explanation:** `The answer is "wke", with the length of 3.`  
`Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.`  

---
**Constraints:**

   * $0$ `<= s.length <=` $5 * 10^4$
   * `s consists of English letters, digits, symbols and spaces.`
 
---
* ### Base data

```Golang
func lengthOfLongestSubstring(s string) int {
}
```

---
### [Solution => 3. Longest Substring Without Repeating Characters](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0003-Longest-Substring-Without-Repeating-Characters/leetcodethree.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0003-Longest-Substring-Without-Repeating-Characters/leetcodethree.go")

---
* ### Metod => Two Pointers
```Golang
simbols = make([]int,1<<8)
for i<len(s) {
	for simbols[s[i]]>1{
		//move start pointer
   }     
}
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/3_Longest_Substring_Without_Repeating_Characters._Two_Pointers.jpg)

---
* ### Metod => Brute Force
```Golang
for i<len(s) {
	for j<len(s){
         bytes.IndexByte
   }     
}
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/3_Longest_Substring_Without_Repeating_Characters_Brute.jpg)
