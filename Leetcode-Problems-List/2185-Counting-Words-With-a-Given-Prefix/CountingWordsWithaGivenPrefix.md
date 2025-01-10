## 2185. Counting Words With a Given Prefix => [original page](https://leetcode.com/problems/counting-words-with-a-given-prefix/description/ "https://leetcode.com/problems/counting-words-with-a-given-prefix/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> |   _Easy_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
You are given an array of strings `words` and a string `pref`.

Return _the number of strings in `words` that contain `pref` as a **prefix**_.

A **prefix** of a string `s` is any leading contiguous substring of `s`.

---
**Example 1:**

**Input:** `words = ["pay","attention","practice","attend"], pref = "at"`  
**Output:** `2`  
**Explanation:** `The 2 strings that contain "at" as a prefix are: "attention" and "attend".`  

**Example 2:**

**Input:** `words = ["leetcode","win","loops","success"], pref = "code"`  
**Output:** `0`  
**Explanation:** `There are no strings that contain "code" as a prefix.`  

---
**Constraints:**

   * $1$ `<= words.length <=` $100$
   * $1$ `<= words[i].length, pref.length <=` $100$
   * `words[i] and pref consist of lowercase English letters.`

---
* ### Base data

```Golang
func prefixCount(words []string, pref string) int {
}
```

---
### [Solution => 2185. Counting Words With a Given Prefix](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2185-Counting-Words-With-a-Given-Prefix/leetcodetwooneeightfive.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2185-Counting-Words-With-a-Given-Prefix/leetcodetwooneeightfive.go")

---
* ### Metod => Compare string objects
```Golang
for _, word := range words {
	
   if find {
      // incremetm answer
   }
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |  $O(N*M)$  |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2185_Counting_Words_With_a_Given_Prefix.jpg)
