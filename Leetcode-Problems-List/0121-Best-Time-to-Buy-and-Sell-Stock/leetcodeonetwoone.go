// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
package leetcodeonetwoone

func maxProfit(prices []int) int {
	var (
		maxProfit = 0
		curProfit = 0
		minPrice  = prices[0]
	)

	for i := 1; i < len(prices); i++ {
		curProfit = prices[i] - minPrice

		if curProfit > maxProfit {
			maxProfit = curProfit
		} else if minPrice > prices[i] {
			minPrice = prices[i]
		}

	}
	return maxProfit
}
