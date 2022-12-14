package solutions

import (
	"fmt"
	"strings"

	"example.com/adventofcode/util"
)

func Fourteen(input string, part int) string {
	rawRockPaths := strings.Split(input, "\n")

	rockPaths, xMax, yMax, xMin, _ := buildRockPaths(rawRockPaths)
	if part == 2 {
		yMax += 2
		floor := []Coordinates{
			{
				x: 0,
				y: yMax,
			},
			{
				x: xMax + xMin,
				y: yMax,
			},
		}
		rockPaths = append(rockPaths, floor)
	}
	drawingOfCave, cave := buildCave(rockPaths)
	drawCave(drawingOfCave, "./day14_before.txt", xMax+xMin, yMax)
	sandGrains := trickleSand(cave, drawingOfCave, yMax, xMax, xMin, part)
	drawCave(drawingOfCave, "./day14_after.txt", xMax*xMax, yMax)
	return fmt.Sprintf("/n Part 1: %d Part 2: %d", sandGrains, sandGrains+1)
}

func trickleSand(cave map[Coordinates]bool, drawingOfMap map[Coordinates]string, yMax, xMax, xMin, part int) int {
	count := 0
	start := &Coordinates{x: 500, y: 0}
	drawingOfMap[*start] = "+"
	for {
		queue := []*Coordinates{start}
		var current *Coordinates

		for len(queue) > 0 {
			current, queue = queue[0], queue[1:]

			if part == 1 && current.y > yMax {
				fmt.Println("trickling away... like my life")
				return count
			}

			next := trickle(current, cave)
			//  Dead end
			if next == nil {
				if part == 2 && current == start {
					return count
				}

				// Sand stops here
				drawingOfMap[*current] = "o"
				cave[*current] = true
				break
			}
			queue = append(queue, next)
		}

		count++
		fmt.Println(count)
	}
}

func trickle(p *Coordinates, grid map[Coordinates]bool) *Coordinates {
	if !grid[Coordinates{x: p.x, y: p.y + 1}] {
		return &Coordinates{x: p.x, y: p.y + 1}
	} else {
		if !grid[Coordinates{x: p.x - 1, y: p.y + 1}] {
			return &Coordinates{x: p.x - 1, y: p.y + 1}
		} else if !grid[Coordinates{x: p.x + 1, y: p.y + 1}] {
			return &Coordinates{x: p.x + 1, y: p.y + 1}
		}
	}
	return nil
}

func buildRockPaths(rawRockPaths []string) ([][]Coordinates, int, int, int, int) {
	rockPaths := make([][]Coordinates, len(rawRockPaths))
	xMin := -1
	xMax := 0
	yMin := -1
	yMax := 0
	for i := range rawRockPaths {
		pathCorners := strings.Split(rawRockPaths[i], " -> ")
		path := make([]Coordinates, len(pathCorners))
		for j := range pathCorners {
			intCoords, err := util.SplitInts(pathCorners[j], ",")
			if err != nil {
				fmt.Println("Something up with parsing the paths to build the cave.")
			}
			x := intCoords[0]
			y := intCoords[1]
			if xMin == -1 || x < xMin {
				xMin = x
			}
			if yMin == -1 || x < yMin {
				yMin = y
			}
			if x > xMax {
				xMax = x
			}
			if y > yMax {
				yMax = y
			}

			path[j] = Coordinates{
				x: x,
				y: y,
			}
		}
		rockPaths[i] = path
	}

	// // make things fit nicely on the screen
	// for i := range rockPaths {
	// 	for j := range rockPaths[i] {
	// 		rockPaths[i][j].x = x

	// 	}
	// }

	return rockPaths, xMax, yMax, xMin, yMin
}

func buildCave(rockPaths [][]Coordinates) (map[Coordinates]string, map[Coordinates]bool) {
	cave := make(map[Coordinates]string)
	grid := make(map[Coordinates]bool)
	for i := range rockPaths {
		for j := range rockPaths[i] {
			corner := rockPaths[i][j]
			cave[Coordinates{y: corner.y, x: corner.x}] = "#"
			grid[Coordinates{y: corner.y, x: corner.x}] = true
			if j != len(rockPaths[i])-1 {
				nextCorner := rockPaths[i][j+1]

				// LEFT/RIGHT
				if corner.y == nextCorner.y {
					// MOVE TO LEFT
					if corner.x > nextCorner.x {
						// spaces left
						for l := corner.x; l >= nextCorner.x; l-- {
							cave[Coordinates{y: nextCorner.y, x: l}] = "#"
							grid[Coordinates{y: nextCorner.y, x: l}] = true
						}
					}

					// MOVE TO RIGHT
					if corner.x < nextCorner.x {
						for r := corner.x; r <= nextCorner.x; r++ {
							cave[Coordinates{y: nextCorner.y, x: r}] = "#"
							grid[Coordinates{y: nextCorner.y, x: r}] = true
						}
					}
				}

				// UP/DOWN
				if corner.x == nextCorner.x {
					// MOVE DOWNWARDS
					if corner.y > nextCorner.y {
						for d := corner.y; d >= nextCorner.y; d-- {
							cave[Coordinates{y: d, x: nextCorner.x}] = "#"
							grid[Coordinates{y: d, x: nextCorner.x}] = true
						}
					}

					// MOVE UPWARDS
					if corner.y < nextCorner.y {
						for u := corner.y; u <= nextCorner.y; u++ {
							cave[Coordinates{y: u, x: nextCorner.x}] = "#"
							grid[Coordinates{y: u, x: nextCorner.x}] = true
						}
					}
				}
			}
		}
	}

	return cave, grid
}

func drawCave(cave map[Coordinates]string, fileName string, xMax, yMax int) string {
	outputFile, err := util.NewFile(fileName).WithWriteableFile()
	if err != nil {
		// Don't make me care.
	}

	result := util.BuildEmptyStringMatrix(xMax, yMax, ".")
	for point, symbol := range cave {
		result[point.y][point.x] = symbol
	}

	res := ""
	for i := range result {
		res += fmt.Sprintf("\n%s", strings.Join(result[i], ""))
	}

	outputFile.Write([]byte(res))
	return res
}
