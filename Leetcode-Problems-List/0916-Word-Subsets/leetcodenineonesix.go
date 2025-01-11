// Solution to problem 916 from leetcode
package leetcodenineonesix

func wordSubsets(words1 []string, words2 []string) []string {
	letters := make([]int, 26)

	//find max count of each letter in words2
	for _, word := range words2 {
		arr := arrLetters(word)
		letters = arrMaxNum(letters, arr)
	}

	universalStrings := make([]string, 0, len(words1))

link:
	for _, word := range words1 {
		arr := arrLetters(word)

		for i, num := range letters {

			if num > arr[i] {
				continue link
			}
		}

		universalStrings = append(universalStrings, word)
	}

	return universalStrings
}

// count the number of letters in one word
func arrLetters(text string) []int {
	arr := make([]int, 26)

	for _, ch := range text {
		arr[ch-'a']++
	}

	return arr
}

// find max count of each letter in two arrays
func arrMaxNum(arr1, arr2 []int) []int {
	for i, num := range arr2 {

		if arr1[i] < num {
			arr1[i] = num
		}
	}

	return arr1
}
