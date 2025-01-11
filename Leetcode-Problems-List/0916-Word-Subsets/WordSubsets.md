## 916. Word Subsets => [original page](https://leetcode.com/problems/word-subsets/description/ "https://leetcode.com/problems/word-subsets/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> | _Medium_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
You are given two string arrays `words1` and `words2`.

A string `b` is a **subset** of string `a` if every letter in `b` occurs in `a` including multiplicity.

   * For example, `"wrr"` is a subset of `"warrior"` but is not a subset of `"world"`.

A string `a` from `words1` is **universal** if for every string `b` in `words2`, `b` is a subset of `a`.

Return an array of all the **universal** strings in `words1`. You may return the answer in **any order**.

---
**Example 1:**

**Input:** `words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","o"]`  
**Output:** `["facebook","google","leetcode"]`  

**Example 2:**

**Input:** `words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["l","e"]`  
**Output:** `["apple","google","leetcode"]`  

---
**Constraints:**

  *  $1$ `<= words1.length, words2.length <=` $10^4$
  *  $1$ `<= words1[i].length, words2[i].length <=` $10$
  *  `words1[i] and words2[i] consist only of lowercase English letters.`
  *  `All the strings of words1 are unique.`
 
---
* ### Base data

```Golang
func wordSubsets(words1 []string, words2 []string) []string {
}
```

---
### [Solution => 916. Word Subsets](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0916-Word-Subsets/leetcodenineonesix.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0916-Word-Subsets/leetcodenineonesix.go")

---
* ### Metod => Max count of each letter in a word in words2 array
```Golang
for _, word := range words2 {
      //find max count of each letter - 'maxCountLetters'
}

link:
for _, word := range words1 {
	
      //create arr of count letters for each word in the words1

      for i, num := range maxCountLetters {	
			
              //compare two arrays			
              if bad {
                 continue link
              }  
      }	
      // increment answer
}
```
| Property | Complexity |              
|:---------|:----------:|
| Time     | $O(N + M)$ |
| Space    |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/0916_Word_Subsets.jpg)
