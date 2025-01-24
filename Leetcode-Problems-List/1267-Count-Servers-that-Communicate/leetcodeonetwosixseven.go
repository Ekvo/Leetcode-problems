// Solution to problem 1267 from leetcode
package leetcodeonetwosixseven

func countServers(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	queue := make([][]int, 0, 1)
	connected := make([][]bool, rows)

	for r := 0; r < rows; r++ {
		connected[r] = make([]bool, cols)

		for c := 0; c < cols; c++ {

			if grid[r][c] == 1 {
				queue = append(queue, []int{r, c})
				break
			}
		}
	}

	numberOfServers := 0

	for len(queue) > 0 {
		point := queue[0]

		row := point[0]
		col := point[1]
		curServer := 0

		queue = queue[1:]

		for r := row + 1; r < rows; r++ {

			if grid[r][col] == 1 {

				if !connected[row][col] {
					connected[row][col] = true
					curServer++
				}
				if !connected[r][col] {
					connected[r][col] = true
					curServer++

					queue = append(queue, []int{r, col})
				}

				break
			}
		}

		for c := col + 1; c < cols; c++ {

			if grid[row][c] == 1 {

				if !connected[row][col] {
					connected[row][col] = true
					curServer++
				}
				if !connected[row][c] {
					connected[row][c] = true
					curServer++
					queue = append(queue, []int{row, c})
				}

				break
			}
		}

		numberOfServers += curServer
	}

	return numberOfServers
}
