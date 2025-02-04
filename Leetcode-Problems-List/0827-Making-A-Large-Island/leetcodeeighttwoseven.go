// Solution to problem 827 from leetcode
package leetcodeeighttwoseven

func largestIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	islandSize := make(map[int]int)
	islandID := 2 // start ID

	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// label each coordinate of island with a islandID
	// find size of the island
	createIsland := func(r, c int) {
		size := 1
		queue := [][]int{{r, c}}
		grid[r][c] = islandID

		for len(queue) > 0 {
			p := queue[0]
			r := p[0]
			c := p[1]
			queue = queue[1:]

			for _, move := range moves {
				newR := r + move[0]
				newC := c + move[1]

				if newR < 0 || newC < 0 || newR >= rows || newC >= cols ||
					grid[newR][newC] != 1 {
					continue
				}

				grid[newR][newC] = islandID
				size++
				queue = append(queue, []int{newR, newC})
			}
		}
		islandSize[islandID] = size
	}

	// we found coordinate with '0' and create size of the new island
	newIslandWithSize := func(r, c int) int {
		size := 1
		visitedIsland := make(map[int]bool, 4)

		for _, move := range moves {
			newR := r + move[0]
			newC := c + move[1]

			if newR < 0 || newC < 0 || newR >= rows || newC >= cols ||
				visitedIsland[grid[newR][newC]] {
				continue
			}

			visitedIsland[grid[newR][newC]] = true
			size += islandSize[grid[newR][newC]]
		}

		return size
	}

	// label all islands on 'grid'
	for r := 0; r < rows; r++ {

		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				createIsland(r, c)
				islandID++
			}
		}
	}

	maxSizeOfIsland := 0

	// find a possible big island
	for r := 0; r < rows; r++ {

		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				maxSizeOfIsland = max(maxSizeOfIsland, newIslandWithSize(r, c))
			}
		}
	}

	// no '0' on the grid
	if maxSizeOfIsland == 0 {
		return rows * cols
	}

	return maxSizeOfIsland
}
