// You are given an array of integers nums of length n and a positive integer k.
//
// The power of an array is defined as:
//
// Its maximum element if all of its elements are consecutive and sorted in ascending order.
// -1 otherwise.
//
// You need to find the power of all
// subarrays
// of nums of size k.
//
// Return an integer array results of size n - k + 1, where results[i] is the power of nums[i..(i + k - 1)].
package leetcodethreetwofivefour

const negativPower = -1

func resultsArray(nums []int, k int) []int {
	var (
		sizeOfResult = len(nums) - k + 1
		results      = make([]int, 0, sizeOfResult)
		subNums      = NewSubArr(nil)
	)

	for i := 0; i+k <= len(nums); {
		subNums.nextSubArr(nums[i : i+k])

		subNums.compute()

		for j := 0; j < subNums.addToIndex() && i+j+k <= len(nums); j++ {
			results = append(results, subNums.powerValue())
		}

		i += subNums.addToIndex()
	}
	return results
}

type subArr struct {
	sArr                   []int
	power                  int
	increaseIndexOfMainArr int
}

func NewSubArr(data []int) *subArr {
	return &subArr{
		sArr:                   data,
		power:                  0,
		increaseIndexOfMainArr: 0,
	}
}

func (sA *subArr) nextSubArr(data []int) {
	sA.sArr = data
	sA.power = 0
	sA.increaseIndexOfMainArr = 0
}

func (sA *subArr) addToIndex() int {
	return sA.increaseIndexOfMainArr
}

func (sA *subArr) powerValue() int {
	return sA.power
}

func (sA *subArr) compute() {
	// create alias for convenience
	var arr = sA.sArr

	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i]-arr[i-1] != 1 {

			//find the beginning of the positive subarray
			//`consecutive` and `sorted` in `ascending order`.
			for i++; i < len(arr) && arr[i] <= arr[i-1]; i++ {
			}

			i--
			sA.increaseIndexOfMainArr = i
			sA.power = negativPower

			return
		}
		if sA.power < arr[i] {
			sA.power = arr[i]
		}
	}
	sA.increaseIndexOfMainArr = 1
}
