package solutions

import (
	"fmt"
	"strings"

	"example.com/adventofcode/util"
)

func Fourteen(input string) string {
	rawRockPaths := strings.Split(input, "\n")

	rockPaths, xMax, yMax, xMin, _ := buildRockPaths(rawRockPaths)
	cave := buildCave(rockPaths, xMax-xMin, yMax)
	cave[0][500-xMin] = "+"
	drawCave(cave)
	fmt.Println("\n\n")
	sandGrains := trickleSand(cave, xMax, xMin, yMax)
	return fmt.Sprintf("/n Part 1: %d", sandGrains)
}

func trickleSand(cave [][]string, xMax, xMin, yMax int) int {
	sandStart := Coordinates{
		x: 500 - xMin,
		y: 0,
	}
	sandGlobs := 0
	for {
		placed := false
		x := sandStart.x
		y := sandStart.y
		for y+1 < yMax {
			y++
			// Until firm ground below
			if cave[y+1][x] != "." {
				pathY := y
				pathX := x

				// Try to move diagonally left
				// space is available and has ground under it and there's no further fall option
				for {
					if cave[pathY][pathX] != "." {
						break
					}

					if pathY+1 > yMax || pathX-2 < 0 {
						return sandGlobs
					}

					if cave[pathY+1][pathX-1] == "." {
						down := pathY + 2
						for {
							if down > yMax-1 {
								break
							}
							if cave[down][pathX-1] != "." {
								break
							}
							down++
						}

						if down >= 54 {
							fmt.Sprintf("debug")
						}
						if cave[down][pathX-2] != "." {
							if down >= 55 && pathX >= 34 {
								fmt.Sprintf("debug")
							}
							cave[down-1][pathX-1] = "o"
							placed = true
							break
						}
					}

					// else if cave[pathY][pathX-1] == "." {
					// 	cave[pathY][pathX-1] = "X"
					// 	placed = true
					// 	break
					// }

					pathY++
					pathX--
				}

				if placed {
					break
				}

				pathY = y
				pathX = x
				// Try to move diagonally right
				// space is available and has ground under it
				for {
					if cave[pathY][pathX] != "." {
						break
					}

					if pathY+1 > yMax || pathX+2 > xMax-xMin {
						return sandGlobs
					}

					if cave[pathY+1][pathX+1] == "." {
						down := pathY + 2
						for {
							if down > yMax-1 {
								break
							}
							if cave[down][pathX+1] != "." {
								break
							}
							down++
						}

						// Further diagonally is blocked
						if cave[down][pathX+2] != "." {
							cave[down-1][pathX+1] = "o"
							placed = true
							break
						}

					}

					// else if cave[pathY][pathX+1] == "." {
					// 	cave[pathY][pathX+1] = "X"
					// 	placed = true
					// 	break
					// }

					pathY++
					pathX++
				}

				if placed {
					break
				}

				if cave[y][x] == "." {
					cave[y][x] = "o"
					placed = true
					break
				}
			}
		}

		if placed {
			sandGlobs++

			if sandGlobs == 381 {
				drawCave(cave)
			}
		} else {
			break
		}
	}

	return sandGlobs
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

	// make things fit nicely on the screen
	for i := range rockPaths {
		for j := range rockPaths[i] {
			rockPaths[i][j].x -= xMin
		}
	}

	return rockPaths, xMax, yMax, xMin, yMin
}

func buildCave(rockPaths [][]Coordinates, xMax, yMax int) [][]string {
	cave := util.BuildEmptyStringMatrix(xMax, yMax, ".")
	for i := range rockPaths {
		for j := range rockPaths[i] {
			corner := rockPaths[i][j]
			cave[corner.y][corner.x] = "#"
			if j != len(rockPaths[i])-1 {
				nextCorner := rockPaths[i][j+1]

				// LEFT/RIGHT
				if corner.y == nextCorner.y {
					// MOVE TO LEFT
					if corner.x > nextCorner.x {
						// spaces left
						for l := corner.x; l >= nextCorner.x; l-- {
							cave[nextCorner.y][l] = "#"
						}
					}

					// MOVE TO RIGHT
					if corner.x < nextCorner.x {
						for r := corner.x; r <= nextCorner.x; r++ {
							cave[nextCorner.y][r] = "#"
						}
					}
				}

				// UP/DOWN
				if corner.x == nextCorner.x {
					// MOVE DOWNWARDS
					if corner.y > nextCorner.y {
						for d := corner.y; d >= nextCorner.y; d-- {
							cave[d][nextCorner.x] = "#"
						}
					}

					// MOVE UPWARDS
					if corner.y < nextCorner.y {
						for u := corner.y; u <= nextCorner.y; u++ {
							cave[u][nextCorner.x] = "#"
						}
					}
				}
			}
		}
	}

	return cave
}

func drawCave(cave [][]string) {
	for i := range cave {
		fmt.Printf("\n%s", strings.Join(cave[i], ""))
	}
}
