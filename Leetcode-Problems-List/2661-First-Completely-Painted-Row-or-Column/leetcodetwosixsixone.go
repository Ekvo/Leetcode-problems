// Solution to problem 2661 from leetcode
package leetcodetwosixsixone

func firstCompleteIndex(arr []int, mat [][]int) int {
	valAndPoint := make(map[int][]int)

	for r, row := range mat {

		for c, val := range row {
			valAndPoint[val] = []int{r, c}
		}
	}

	rows := len(mat)
	cols := len(mat[0])
	arrRows := make([]int, rows)
	arrCols := make([]int, cols)

	for i, val := range arr {
		point := valAndPoint[val]
		r := point[0]
		c := point[1]

		arrRows[r]++
		arrCols[c]++

		if arrRows[r] == cols || arrCols[c] == rows {
			return i
		}
	}

	return -1
}
