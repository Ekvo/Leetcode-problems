// Solution to problem 1765 from leetcode
package leetcodeonesevensixfive

func highestPeak(isWater [][]int) [][]int {
	rows := len(isWater)
	cols := len(isWater[0])

	heights := make([][]int, rows)
	queue := make([][]int, 0, rows*cols)

	for r := 0; r < rows; r++ {
		heights[r] = make([]int, cols)

		for c := 0; c < cols; c++ {

			if isWater[r][c] == 1 {
				heights[r][c] = 0
				queue = append(queue, []int{r, c})
			} else {
				heights[r][c] = -1
			}
		}
	}

	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		point := queue[0]

		row := point[0]
		col := point[1]
		height := heights[row][col]

		queue = queue[1:]

		for _, move := range moves {
			newRow := row + move[0]
			newCol := col + move[1]

			if newRow < 0 ||
				newCol < 0 ||
				newRow >= rows ||
				newCol >= cols ||
				heights[newRow][newCol] != -1 {
				continue
			}

			heights[newRow][newCol] = height + 1
			queue = append(queue, []int{newRow, newCol})
		}
	}

	return heights
}
