// Solution to problem 2381 from leetcode
package leetcodetwothreeeightone

//Approach: Difference Array
/*
Building on the idea of cumulative sums, we can use a difference array to handle range updates more efficiently.
A difference array helps us record changes in values between consecutive elements rather than updating every element in a range directly.

Instead of keeping track of how many shifts should be applied to each character in the alphabet, we’ll use the difference array
to store how many more shifts should be applied to the current character compared to the previous one.
This allows us to record changes only at the starting and ending points of shifts, rather than updating each character in the range.

For convenience, a positive shift means that the character must move forward in the alphabet, and a negative shift means that it must move backward.
*/
func shiftingLettersDifArr(s string, shifts [][]int) string {
	difArr := make([]int, len(s)+1)

	for _, shift := range shifts {

		if shift[2] != 0 {
			difArr[shift[0]] += 1
			difArr[shift[1]+1] -= 1

		} else {
			difArr[shift[0]] -= 1
			difArr[shift[1]+1] += 1
		}
	}

	arr := []byte(s)
	n := int('z' - 'a' + 1)
	difArrSum := 0

	for i, ch := range arr {
		difArrSum += difArr[i]

		newCh := (int(ch-'a') + difArrSum) % n
		if newCh < 0 {
			newCh += n
		}
		arr[i] = byte(newCh) + 'a'
	}

	return string(arr)
}
