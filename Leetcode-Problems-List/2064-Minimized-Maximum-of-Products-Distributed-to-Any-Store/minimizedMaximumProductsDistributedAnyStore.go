// You are given an integer n indicating there are n specialty retail stores.
// There are m product types of varying amounts, which are given as a 0-indexed integer array quantities,
// where quantities[i] represents the number of products of the ith product type.
//
// You need to distribute all products to the retail stores following these rules:
//
// A store can only be given at most one product type but can be given any amount of it.
// After distribution, each store will have been given some number of products (possibly 0).
// Let x represent the maximum number of products given to any store. You want x to be as small as possible,
// i.e., you want to minimize the maximum number of products that are given to any store.
//
// Return the minimum possible x.
package minimizedMaximumProductsDistributedAnyStore

func minimizedMaximum(n int, quantities []int) int {
	var (
		minProduct = 1
		maxProduct = 0
	)

	for i := 0; i < len(quantities); i++ {
		if maxProduct < quantities[i] {
			maxProduct = quantities[i]
		}
	}

	var (
		result           = 0
		currentProduct   = (minProduct + maxProduct) / 2
		goodDistribution func(product int) bool
	)

	goodDistribution = func(product int) bool {
		var store = 0

		for i := 0; i < len(quantities); i++ {
			store += 1 + (quantities[i]-1)/product
		}

		if store <= n {
			return true
		}
		return false
	}

	for minProduct <= maxProduct {
		if goodDistribution(currentProduct) {
			result = currentProduct
			maxProduct = currentProduct - 1
		} else {
			minProduct = currentProduct + 1
		}
		currentProduct = (minProduct + maxProduct) / 2
	}

	return result
}
