## 1408. String Matching in an Array => [original page](https://leetcode.com/problems/string-matching-in-an-array/description/ "https://leetcode.com/problems/string-matching-in-an-array/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Easy_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given an array of string `words`, return _all strings in `words` that is a **substring** of another word_. You can return the answer in **any order**.

A **substring** is a contiguous sequence of characters within a string

---
**Example 1:**

**Input:** `words = ["mass","as","hero","superhero"]`  
**Output:** `["as","hero"]`  
**Explanation:** `"as" is substring of "mass" and "hero" is substring of "superhero".`  
`["hero","as"] is also a valid answer.`  

**Example 2:**

**Input:** `words = ["leetcode","et","code"]`  
**Output:** `["et","code"]`  
**Explanation:** `"et", "code" are substring of "leetcode".`  

**Example 3:**

**Input:** `words = ["blue","green","bu"]`  
**Output:** `[]`  
**Explanation:** `No string of words is substring of another string.`  

---
**Constraints:**

   * $1$ `<= words.length <=` $100$
   * $1$ `<= words[i].length <=` $30$
   * `words[i] contains only lowercase English letters.`
   * `All the strings of words are unique.`
 
---
* ### Base data

```Golang
func stringMatching(words []string) []string {
}
```

---
### [Solution => 1408. String Matching in an Array](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1408-String-Matching-in-an-Array/leetcodeonefourzeroeight.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1408-String-Matching-in-an-Array/leetcodeonefourzeroeight.go")

---
* ### Metod => sort.Slice & strings.Contain
```Golang
sort.Slice //return len(words[i]) < len(words[j])

for i := 0; i < len(words); i++ {

   for j := i + 1; j < len(words); j++ {
           //find substring
   }   
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     |  $O(N^2)$  |
| Space    |   $O(N)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1408_String_Matching_in_an_Array_Contain.jpg)
