## 27. Remove Element => [original page](https://leetcode.com/problems/remove-element/description/ "https://leetcode.com/problems/remove-element/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Easy_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given an integer array nums and an integer `val`, remove all occurrences of `val` in `nums` **in-place**. The order of the elements may be changed. Then return _the number of elements in_ `nums` _which are not equal to_ `val`.

Consider the number of elements in `nums` which are not equal to `val` be `k`, to get accepted, you need to do the following things:

   * Change the array `nums` such that the first `k` elements of `nums` contain the elements which are not equal to `val`. The remaining elements of `nums` are not important as well as the size of `nums`.
   * Return `k`.

**Custom Judge:**

The judge will test your solution with the following code:

`int[] nums = [...]; // Input array`  
`int val = ...; // Value to remove`  
`int[] expectedNums = [...]; // The expected answer with correct length.`  
`// It is sorted with no values equaling val.`  

`int k = removeElement(nums, val); // Calls your implementation`

`assert k == expectedNums.length;`  
`sort(nums, 0, k); // Sort the first k elements of nums`  
`for (int i = 0; i < actualLength; i++) {`  
`assert nums[i] == expectedNums[i];`  
`}`  

If all assertions pass, then your solution will be **accepted**.


---
**Example 1:**

**Input:** `nums = [3,2,2,3], val = 3`  
**Output:** `2, nums = [2,2,_,_]`  
**Explanation:** `Your function should return k = 2, with the first two elements of nums being 2.`  
`It does not matter what you leave beyond the returned k (hence they are underscores).`  

**Example 2:**

**Input:** `nums = [0,1,2,2,3,0,4,2], val = 2`  
**Output:** `5, nums = [0,1,4,0,3,_,_,_]`  
**Explanation:** `Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.`  
`Note that the five elements can be returned in any order.`  
`It does not matter what you leave beyond the returned k (hence they are underscores).`  

---
**Constraints:**

   * $0$ `<= nums.length <=` $100$
   * $0$ `<= nums[i] <=` $50$
   * $0$ `<= val <=` $100$

---
* ### Base data

```Golang
func removeElement(nums []int, val int) int {
}
```

---
### [Solution => 27. Remove Element](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0027-Remove-Element/leetcodetwoseven.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0027-Remove-Element/leetcodetwoseven.go")

---
* ### Metod => Two Pointers
```Golang
  for {
    if nums[i] == val {
        continue
     }
   //array refactor 
   }
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/27_Remove_Element.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/27_Remove_Element_Time.jpg)

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/27_Remove_Element_Space.jpg)
