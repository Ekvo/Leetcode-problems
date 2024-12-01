// Given an array arr of integers, check if there exist two indices i and j such that :
//
// i != j
// 0 <= i, j < arr.length
// arr[i] == 2 * arr[j]
package leetcodeonethreefoursix

func checkIfExist(arr []int) bool {
	var jValue = make(map[int]bool, len(arr)<<1-1)

	for i := 0; i < len(arr); i++ {
		if jValue[arr[i]] {
			return true
		}
		if arr[i]%2 == 0 {
			jValue[arr[i]>>1] = true // arr[i]/2
		}
		jValue[arr[i]<<1] = true // arr[i]*2
	}
	return false
}
