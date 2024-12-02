## 1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence => [original page](https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/description/ "https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Easy_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given a `sentence` that consists of `some` words separated by a **single space**, and a `searchWord`, check if `searchWord` is a prefix of any word in `sentence`.

Return the index of the word in `sentence` (**1-indexed**) where searchWord is a prefix of this word. If `searchWord` is a prefix of more than one word, return the index of the first word (**minimum index**). If there is no such word return `-1`.

A **prefix** of a string s is any leading contiguous substring of `s`.


---
**Example 1:**

**Input:** `sentence = "i love eating burger", searchWord = "burg"`  
**Output:** `4`  
**Explanation:** `"burg" is prefix of "burger" which is the 4th word in the sentence.`  

**Example 2:**

**Input:** `sentence = "this problem is an easy problem", searchWord = "pro"`  
**Output:** `2`  
**Explanation:** `"pro" is prefix of "problem" which is the 2nd and the 6th word in the sentence, but we return 2 as it's the minimal index.`  

**Example 3:**

**Input:** `sentence = "i am tired", searchWord = "you"`  
**Output:** `-1`  
**Explanation:** `"you" is not a prefix of any word in the sentence.`  

---
**Constraints:**

   * $1$ `<= sentence.length <=` $100$
   * $1$ `<= searchWord.length <=` $10$
   * `sentence consists of lowercase English letters and spaces.`
   * `searchWord consists of lowercase English letters.`

---
* ### Base data
```Golang
func isPrefixOfWord(sentence string, searchWord string) int {
}
```

---
### [Solution => 1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1455-Check-If-a-Word-Occurs-As-a-Prefix-of-Any-Word-in-a-Sentence/leetcodeonefourfivefive.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/1455-Check-If-a-Word-Occurs-As-a-Prefix-of-Any-Word-in-a-Sentence/leetcodeonefourfivefive.go")

---
* ### Metod => Compare characters
```Golang
for len(sentence) {
	for len(searchWord) {		
    }
}
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1455_Check_If_a_Word_Occurs_As_a_Prefix_of_Any_Word_in_a_Sentence.jpg)
 
---
* ### Metod => Strings
```Golang
strings.Split
strings.HasPrefix
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/1455_Check_If_a_Word_Occurs_As_a_Prefix_of_Any_Word_in_a_Sentence_str.jpg)
 
