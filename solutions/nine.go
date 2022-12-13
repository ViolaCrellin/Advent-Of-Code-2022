package solutions

import (
	"fmt"
	"strings"

	"example.com/adventofcode/util"
)

func Nine(input string) string {
	instructions := strings.Split(input, "\n")
	//xMax, yMax, xStart, yStart := SketchSize(instructions)
	//motions := util.BuildEmptyIntMatrix(xMax, yMax)
	motions := util.BuildEmptyIntMatrix(6, 6)
	positions := make(map[string]bool, 0)
	// motions[y][x]
	//headY := yStart
	//headX := xStart
	//tailY := yStart
	//tailX := xStart

	headY := 6
	headX := 0
	tailY := 6
	tailX := 0

	coordsStr := fmt.Sprintf("%d:%d", tailY, tailX)
	positions[coordsStr] = true
	motions[tailY][tailX] = 1
	for i := range instructions {
		direction, qty, err := util.SplitLetterAndNumber(instructions[i], " ")
		if err != nil {
			//Meh
		}

		for j := 0; j <= qty; j++ {
			if tailX == 3 {
				fmt.Println("argh")
			}

			switch direction {
			case "R":
				headX++
			case "L":
				headX--
			case "U":
				headY--
			case "D":
				headY++
			}

			if IsDiagonallyAdjacent(tailX, tailY, headX, headY) || IsAdjacent(tailX, tailY, headX, headY) {
				continue
			}

			// A diagonal move
			//
			//
			// ......
			// ......
			// ..T...
			// ......
			// ....H.
			//
			// ......
			// ......
			// ..T...
			// ...H..
			// ......
			if tailX+2 == headX && tailY-2 == headY {
				tailX++
				tailY--
				update(positions, motions, tailX, tailY)
				continue
			}

			// A diagonal move
			//
			//
			// ......
			// ......
			// ..H...
			// ......
			// ....T.
			//
			// ......
			// ......
			// ..H...
			// ...T..
			// ......
			if tailX-2 == headX && tailY-2 == headY {
				tailX--
				tailY--
				update(positions, motions, tailX, tailY)
				continue
			}

			// A diagonal move
			//
			//
			// ......
			// ......
			// ...T..
			// ......
			// .....H
			//
			// ......
			// ......
			// ......
			// ....T.
			// .....H
			if tailX+2 == headX && tailY+2 == headY {
				tailX++
				tailY++
				update(positions, motions, tailX, tailY)
				continue
			}

			// A diagonal move
			//
			//
			// .....H
			// ......
			// ...T..
			// ......
			// ......
			//
			// .....H
			// ......
			// ...T..
			// ......
			// ......
			if tailX+2 == headX && tailY-2 == headY {
				tailX++
				tailY--
				update(positions, motions, tailX, tailY)
				continue
			}

			// We are two to the left of H
			//
			//
			// ......
			// ......
			// ......
			// ......
			// ..T.H.
			//
			// ......
			// ......
			// ......
			// ......
			// ...TH.
			if tailX+2 == headX {
				tailX++
				update(positions, motions, tailX, tailY)
				continue
			}

			// We are two to the right of H
			//
			//
			// ......
			// ......
			// ......
			// ......
			// ..H.T.
			//
			// ......
			// ......
			// ......
			// ......
			// ..HT..
			if tailX-2 == headX {
				tailX--
				update(positions, motions, tailX, tailY)
				continue
			}

			// We are two to above H
			//
			// ......
			// ......
			// ...T..
			// ......
			// ...H..
			//
			// ......
			// ......
			// ......
			// ...T..
			// ...H..
			if tailY-2 == headY {
				tailY--
				update(positions, motions, tailX, tailY)
				continue
			}

			// We are two to below H
			//
			// ......
			// ......
			// ...H..
			// ......
			// ...T..
			//
			// ......
			// ......
			// ......
			// ...H..
			// ...T..
			if tailY+2 == headY {
				tailY++
				update(positions, motions, tailX, tailY)
				continue
			}
		}

	}

	part1 := len(positions)
	print(motions)
	part2 := util.MultidimensionalSum(motions, true)
	//minus 2 for the start and end positions where we will be one behind/or on top but not visited.
	return fmt.Sprintf("Part 1: %d, %d", part1, part2)
}

func update(positions map[string]bool, motions [][]int, tailX, tailY int) {
	coordsStr := fmt.Sprintf("%d:%d", tailY, tailX)
	positions[coordsStr] = true
	motions[tailY][tailX] = 1
	print(motions)
}

func SketchSize(instructions []string) (int, int, int, int) {
	headX := 0
	headY := 0
	maxX := 0
	minX := 0
	maxY := 0
	minY := 0
	for i := range instructions {
		direction, qty, err := util.SplitLetterAndNumber(instructions[i], " ")
		if err != nil {
			//Meh
		}
		switch direction {
		case "R":
			headX += qty
			if headX > maxX {
				maxX = headX
			}
		case "L":
			headX -= qty
			if headX < minX {
				minX = headX
			}
		case "U":
			headY -= qty
			if headY < minY {
				minY = headY
			}
		case "D":
			headY += qty
			if headY > maxY {
				maxY = headY
			}
		}
	}

	startX := util.Abs(minX)
	startY := util.Abs(minY)
	width := startX + maxX
	height := startY + maxY

	return width, height, startX, startY
}

// IsDiagonallyAdjacent
// This is intentionally clunky for readability purposes.
func IsDiagonallyAdjacent(tailX, tailY, headX, headY int) bool {

	// ......
	// ......
	// ..T...
	// ...H..
	// ......
	if tailX+1 == headX && tailY+1 == headY {
		return true
	}

	// ......
	// ......
	// ....T.
	// ...H..
	// ......
	if tailX-1 == headX && tailY+1 == headY {
		return true
	}

	// ......
	// ......
	// ..H...
	// .T....
	// ......
	if tailX+1 == headX && tailY-1 == headY {
		return true
	}

	// ......
	// ......
	// ......
	// .H....
	// ..T...
	if tailX-1 == headX && tailY-1 == headY {
		return true
	}

	return false
}

func IsAdjacent(tailX, tailY, headX, headY int) bool {
	if tailX == headX && tailY == headY {
		return true
	}

	if tailX+1 == headX && tailY == headY {
		return true
	}

	if tailX-1 == headX && tailY == headY {
		return true
	}

	if tailX == headX && tailY+1 == headY {
		return true
	}

	if tailX == headX && tailY-1 == headY {
		return true
	}

	return false
}

func print(motions [][]int) {
	for i := range motions {
		line := ""
		for j := range motions[i] {
			if motions[i][j] == 1 {
				line += "#"
			} else if motions[i][j] == -1 {
				line += "H"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}

	fmt.Print("\n")
}
