// Solution to problem 1368 from leetcode
package leetcodeonethreesixeight

import (
	"container/heap"
	"math"
)

func minCost(grid [][]int) int {
	nRows := len(grid)
	nCols := len(grid[nRows-1])

	// Distance array (cells) to store minimum cost for each cell
	cells := make([][]int, nRows)

	for i := 0; i < nRows; i++ {
		cells[i] = make([]int, nCols)

		// Mark each cell -> high cost
		sliceWithStartValue(cells[i], math.MaxInt)
	}
	cells[0][0] = 0 // Start -> cost 0

	queue := &Costs{PlaceCost{0, 0, 0}}
	heap.Init(queue)

	// Moves:         Right(1) Left(2)  Down(3)  Up (4)
	moves := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// Dijkstra's algorithm
	for queue.Len() > 0 {
		pc := heap.Pop(queue).(PlaceCost)
		cost := pc.cost
		row := pc.row
		col := pc.col

		// If visited this cell -> continue
		if cells[row][col] < cost {
			continue
		}

		// Check all possible moves
		for i, move := range moves {
			newRow := row + move[0]
			newCol := col + move[1]

			if newRow < 0 ||
				newRow >= nRows ||
				newCol < 0 ||
				newCol >= nCols {
				continue
			}

			// Calculate new cost
			newCost := cost

			if grid[row][col] != i+1 {
				newCost++
			}

			// If new cost is better, update and push
			if cells[newRow][newCol] > newCost {
				cells[newRow][newCol] = newCost
				heap.Push(queue, PlaceCost{newCost, newRow, newCol})
			}
		}
	}

	return cells[nRows-1][nCols-1]
}

func sliceWithStartValue(slice []int, startVal int) []int {
	for i, _ := range slice {
		slice[i] = startVal
	}

	return slice
}

type Costs []PlaceCost

type PlaceCost struct {
	cost int
	row  int
	col  int
}

func (c Costs) Len() int           { return len(c) }
func (c Costs) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Costs) Less(i, j int) bool { return c[i].cost < c[j].cost }
func (c *Costs) Push(obj any)      { *c = append(*c, obj.(PlaceCost)) }
func (c *Costs) Pop() any {
	old := *c
	lenght := len(old)
	item := old[lenght-1]
	*c = old[:lenght-1]
	return item
}
