// You are given a 2D integer array items where items[i] = [pricei, beautyi]
// denotes the price and beauty of an item respectively.
//
// You are also given a 0-indexed integer array queries.
// For each queries[j], you want to determine the maximum beauty of an item whose price is less than or equal to queries[j]. If no such item exists, then the answer to this query is 0.
//
// Return an array answer of the same length as queries where answer[j]
// is the answer to the jth query
package leetcodetwozerosevenzero

import "sort"

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] > items[j][1]
		}
		return items[i][0] < items[j][0]
	})

	// All the time max beauty at the beginning of the price
	// example: [1,7], [1,4], [1,3], [1,2], [2,5], [2,3], [2,1]
	var maxBeauty = items[0][1]

	for i := 0; i < len(items); i++ {
		if maxBeauty < items[i][1] {
			maxBeauty = items[i][1]
			continue
		}
		if items[i][1] < maxBeauty {
			items[i][1] = maxBeauty
		}
	}

link:
	for i := 0; i < len(queries); i++ {
		var (
			head         = 0
			tail         = len(items) - 1
			mid          = (head + tail) / 2
			currentPrice = items[mid][0]
		)

		for head <= tail {
			if currentPrice < queries[i] {
				head = mid + 1
			} else if currentPrice > queries[i] {
				tail = mid - 1
			} else {
				queries[i] = items[mid][1]
				continue link
			}
			mid = (head + tail) / 2
			currentPrice = items[mid][0]
		}

		if -1 < tail && tail < len(items) {
			queries[i] = items[tail][1]
		} else {
			queries[i] = 0
		}
	}

	return queries
}
