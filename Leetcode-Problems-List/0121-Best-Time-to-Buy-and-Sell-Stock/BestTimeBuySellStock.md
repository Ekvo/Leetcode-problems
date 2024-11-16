## 121. Best Time to Buy and Sell Stock => [original page](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/ "https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/")

---

| Property               |      |   Target |              
|------------------------|:----:|---------:|
| complexity of the task | ---> |   _Easy_ |
| status                 | ---> | _Solved_ |

---
**Task:**  
You are given an array `prices` where `prices[i]` is the price of a given stock on the `i-th` day.

You want to maximize your profit by choosing a **single day** to buy one stock and choosing a **different day in the future** to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, `return 0`.


---
**Example 1:**

**Input:** `prices = [7,1,5,3,6,4]`  
**Output:** `5`  
**Explanation:** `Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.`  
`Note` that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.`

**Example 2:**

**Input:** `prices = [7,6,4,3,1]`  
**Output:** `0`  
**Explanation:** `In this case, no transactions are done and the max profit = 0.`

---
**Constraints:**  
    * $1$ `<= prices.length <=` $10^5$  
    * $0$ `<= prices[i] <=` $10^4$  

---
* ### Base data

```Golang
func maxProfit(prices []int) int {
}
```

---
### [Solution => 121. Best Time to Buy and Sell Stock](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0121-Best-Time-to-Buy-and-Sell-Stock/bestTimeBuySellStock.go "https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-List/0121-Best-Time-to-Buy-and-Sell-Stock/bestTimeBuySellStock.go")

---
* ### Metod => Dynamic Programming
```Golang
var (
    maxProfit = 0
    curProfit = 0
    minPrice  = prices[0]
)
```

![submit](https://github.com/Ekvo/Leetcode-problems/blob/main/Leetcode-Problems-Submit-Screenshots/121_Best_Time_to_Buy_and_Sell_Stock.jpg)
 
