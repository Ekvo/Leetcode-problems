// Solution to problem 2658 from leetcode
package leetcodetwosixfiveeight

func findMaxFish(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	countFish := 0
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var getFishs = func(row, col int) int {
		queue := [][]int{{row, col}}
		currentFish := grid[row][col]
		grid[row][col] = 0

		for len(queue) > 0 {
			p := queue[0]
			row = p[0]
			col = p[1]
			queue = queue[1:]

			for _, move := range moves {
				newRow := row + move[0]
				newCol := col + move[1]

				if newRow < 0 ||
					newCol < 0 ||
					newRow >= rows ||
					newCol >= cols ||
					grid[newRow][newCol] == 0 {
					continue
				}

				currentFish += grid[newRow][newCol]
				grid[newRow][newCol] = 0
				queue = append(queue, []int{newRow, newCol})
			}

		}
		return currentFish
	}

	for r := 0; r < rows; r++ {

		for c := 0; c < cols; c++ {

			if grid[r][c] != 0 {
				countFish = max(countFish, getFishs(r, c))
			}
		}
	}

	return countFish
}
