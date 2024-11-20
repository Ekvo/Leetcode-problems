## 224. Basic Calculator => [original page](https://leetcode.com/problems/basic-calculator/description/ "https://leetcode.com/problems/basic-calculator/description/")

---
| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Hard_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
Given a string `s` representing a valid expression, implement a basic calculator to evaluate it, and return the result of the evaluation.

**Note:** You are **not** allowed to use any built-in function which evaluates strings as mathematical expressions, such as `eval()`.

---
**Example 1:**

**Input:** `s = "1 + 1"`  
**Output:** `2`  

**Example 2:**

**Input:** `s = " 2-1 + 2 "`  
**Output:** `3`  

**Example 3:**

**Input:** `s = "(1+(4+5+2)-3)+(6+8)"`  
**Output:** `23`  

---
**Constraints:**
  *  $1$ `<= s.length <=` $3 * 10^5$
  *  `s consists of digits, '+', '-', '(', ')', and ' '.`
  *  `s represents a valid expression.`
  *  `'+' is not used as a unary operation (i.e., "+1" and "+(2 + 3)" is invalid).`
  *  `'-' could be used as a unary operation (i.e., "-1" and "-(2 + 3)" is valid).`
  *  `There will be no two consecutive operators in the input.`
  *  `Every number and running calculation will fit in a signed 32-bit integer.`

---
* ### Base data

```Golang
func calculate(s string) int {	
}
```

---
### [Solution => 224. Basic Calculator](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0224-Basic-Calculator/leetcodetwotwofour.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0224-Basic-Calculator/leetcodetwotwofour.go")

---
* ### Metod => Biern Stroustrup ---> Recursion Calculator
```Golang
type compute struct {
    data  string
    index int
}
func newCalculate(s string) *compute
func (c *compute) work() bool //index check
func (c *compute) increment() //index++
func (c *compute) decrement() //index--
func (c *compute) newSimbol() byte 
func (c *compute) newValue() int
// calculates '+' or  '-' and return result one operation
func (c *compute) expression() int
// find: value, expression(), unaryMinus and return value
func (c *compute) prime() int   
/*
example: 4 + 3
4 ---> number ---> prime() ---->| expression
+ ------------------------------|----------> expression
3 ---> number ---> prime()------|
*/
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/224_Basi_Calculator.jpg)
 
