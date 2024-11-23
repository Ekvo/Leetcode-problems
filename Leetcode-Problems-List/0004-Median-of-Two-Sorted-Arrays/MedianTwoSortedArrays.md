## 4. Median of Two Sorted Arrays => [original page](https://leetcode.com/problems/median-of-two-sorted-arrays/description/ "https://leetcode.com/problems/median-of-two-sorted-arrays/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Hard_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median** of the two sorted arrays.

The overall run time complexity should be `O(log (m+n))`.

---
**Example 1:**

**Input:** `nums1 = [1,3], nums2 = [2]`  
**Output:** `2.00000`  
**Explanation:** `merged array = [1,2,3] and median is 2.`  

**Example 2:**

**Input:** `nums1 = [1,2], nums2 = [3,4]`  
**Output:** `2.50000`  
**Explanation:** `merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.`  

---
**Constraints:**

   * `nums1.length == m`
   * `nums2.length == n`
   * $0$ `<= m <=` $1000$
   * $0$ `<= n <=` $1000$
   * $1$ `<= m + n <=` $2000$
   * $-10^6$ `<= nums1[i], nums2[i] <=` $10^6$

---
* ### Base data

```Golang
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
}
```

---
### [Solution => 4. Median of Two Sorted Arrays](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0004-Median-of-Two-Sorted-Arrays/leetcodefour.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0004-Median-of-Two-Sorted-Arrays/leetcodefour.go")

---
* ### Metod => Two Pointers
```Golang
var (
	lenght = len(nums1) + len(nums2)
	l1 = 0 //pointer to nums1
	l2 = 0 //pointer to nums2
)
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/4_Median_of_Two_Sorted_Arrays.jpg)

* ### Metod => Brute Force >> (sort.Ints)
```Golang
var slice = append(nums1, nums2...)
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/4_Median_of_Two_Sorted_Arrays_Brute.jpg)
 


