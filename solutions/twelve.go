package solutions

import (
	"fmt"
	"math"

	"example.com/adventofcode/util"
)

type Coordinates struct {
	x, y int
}

func Twelve(input string, part int) string {
	grid := util.BuildByteMatrixFromString(input, "\n")
	var (
		starts []Coordinates
		end    Coordinates
	)

	startCriteria := func(part, i, j int) bool {
		if part == 1 {
			return grid[i][j] == 'S'
		}
		return grid[i][j] == 'S' || grid[i][j] == 'a'
	}

	for i := range grid {
		for j := range grid[i] {
			if startCriteria(part, i, j) {
				starts = append(starts, Coordinates{i, j})
			} else if grid[i][j] == 'E' {
				end = Coordinates{i, j}
			}
		}
	}

	minSteps := math.MaxInt

	// to remove E from grid
	grid[end.x][end.y] = 'z'
	for _, start := range starts {
		grid[start.x][start.y] = 'a'
		// to remove possible 'S' from grid
		curSteps := BFS(grid, start, end)
		minSteps, _ = util.MinAndMax([]int{minSteps, curSteps})
	}

	// Add two for the step from the start and to the end
	return fmt.Sprintf("%d", minSteps)
}

func BFS(grid [][]byte, start, end Coordinates) int {
	visited := make(map[Coordinates]bool)
	visited[start] = true
	queue := []Coordinates{start}
	steps := 0
	var directions = []Coordinates{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	found := false

	// simple BFS
L:
	for len(queue) > 0 {
		k := len(queue)
		for i := 0; i < k; i++ {
			current := queue[0]
			queue = queue[1:]

			if current == end {
				found = true
				break L
			}

			for _, dir := range directions {
				newC := Coordinates{current.x + dir.x, current.y + dir.y}
				// Off the grid
				if newC.x < 0 || newC.x >= len(grid) || newC.y < 0 || newC.y >= len(grid[0]) {
					continue
				}

				// Been there done that
				if visited[newC] {
					continue
				}
				// Can't climb it
				isnewCGreater := grid[newC.x][newC.y] > grid[current.x][current.y]
				if isnewCGreater && grid[newC.x][newC.y]-grid[current.x][current.y] > 1 {
					continue
				}
				visited[newC] = true
				queue = append(queue, newC)
			}
		}
		steps++
	}

	if !found {
		fmt.Println(start, "to", end, "not found")
		return math.MaxInt
	} else {
		fmt.Println("FOUND")
	}

	return steps
}
