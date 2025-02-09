## 2349. Design a Number Container System => [original page](https://leetcode.com/problems/design-a-number-container-system/description/ "https://leetcode.com/problems/design-a-number-container-system/description/")

---
| Property                |      |   Target |              
|:------------------------|:----:|---------:|
| complexity of the task  | ---> | _Medium_ |
| status                  | ---> | _Solved_ |

---
**Task:**  
Design a number container system that can do the following:

   * **Insert** or **Replace** a number at the given index in the system.
   * **Return** the smallest index for the given number in the system.

Implement the `NumberContainers` class:

   * `NumberContainers()` Initializes the number container system.
   * `void change(int index, int number)` Fills the container at `index` with the number. If there is already a number at that `index`, replace it.
   * `int find(int number)` Returns the smallest index for the given `number`, or `-1` if there is no index that is filled by `number` in the system.

---
**Example 1:**

**Input**  
`["NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"]`  
`[[], [10], [2, 10], [1, 10], [3, 10], [5, 10], [10], [1, 20], [10]]`  
**Output**  
`[null, -1, null, null, null, null, 1, null, 2]`  
**Explanation**  
`NumberContainers nc = new NumberContainers();`  
`nc.find(10); // There is no index that is filled with number 10. Therefore, we return -1.`  
`nc.change(2, 10); // Your container at index 2 will be filled with number 10.`  
`nc.change(1, 10); // Your container at index 1 will be filled with number 10.`  
`nc.change(3, 10); // Your container at index 3 will be filled with number 10.`  
`nc.change(5, 10); // Your container at index 5 will be filled with number 10.`  
`nc.find(10); // Number 10 is at the indices 1, 2, 3, and 5. Since the smallest index that is filled with 10 is 1, we return 1.`  
`nc.change(1, 20); // Your container at index 1 will be filled with number 20. Note that index 1 was filled with 10 and then replaced with 20.`  
`nc.find(10); // Number 10 is at the indices 2, 3, and 5. The smallest index that is filled with 10 is 2. Therefore, we return 2.`  

---
**Constraints:**

   * $1$ `<= index, number <=` $10^9$
   * `At most` $10^5$ `calls will be made in total to change and find.`
 
---
* ### Base data

```Golang
type NumberContainers struct {    
}

func Constructor() NumberContainers {    
}

func (this *NumberContainers) Change(index int, number int)  {    
}

func (this *NumberContainers) Find(number int) int {    
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
```

---
### [Solution => 2349. Design a Number Container System](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2349-Design-a-Number-Container-System/leetcodetwothreefournine.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/2349-Design-a-Number-Container-System/leetcodetwothreefournine.go")

---
* ### Metod => Struct 'store' & Two HashMap
```Golang
type store struct {
   addr     []int //indexs for number
   addrSort bool  //mark for check sort indexs
}

type NumberContainers struct {
   indexNumber map[int]int   // index : number
   numberIndex map[int]store // number : index
}

func Constructor() NumberContainers
func (this *NumberContainers) Find(number int) int
func (this *NumberContainers) Change(index int, number int)

// add 'number' with 'index' in to NumberContainers 
func (this *NumberContainers) add(index int, number int)
// delete 'number' with 'index' from NumberContainers 
func (this *NumberContainers) remove(index int, number int)
```
| Property     | Complexity |              
|:-------------|:----------:|
| Time:        |            |
| ------------ | ---------- |
| Constructor  |   $O(N)$   |
| Find         |   $O(N)$   |
| Change       |   $O(1)$   |
| ------------ | ---------- |
| Space:       |            |
| ------------ | ---------- |
| Constructor  |   $O(N)$   |
| Find         |   $O(N)$   |
| Change       |   $O(1)$   |

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/2349_Design_a_Number_Container_System.jpg)
