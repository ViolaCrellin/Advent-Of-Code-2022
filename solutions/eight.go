package solutions

import (
	"fmt"

	"example.com/adventofcode/util"
)

func Eight(input string) string {
	forest := util.BuildIntMatrixFromString(input, "\n", "")

	xLen := len(forest[0])
	yLen := len(forest)

	visibilityMap := util.BuildEmptyIntMatrix(xLen-1, yLen-1)
	numOfTrees := xLen * yLen

	// How many trees are NOT visible may be easier
	// A tree is NOT visible if any of the trees between it and the edge are taller than it
	flippedForest := AssessWhatTreesAreVisibleInRow(forest, visibilityMap, false)
	AssessWhatTreesAreVisibleInRow(flippedForest, visibilityMap, true)
	invisibleTreesCount := 0
	for i := range visibilityMap {
		row := ""
		for j := range visibilityMap[i] {
			sidesNotVisibleFrom := visibilityMap[i][j]
			row += fmt.Sprintf("%d", visibilityMap[i][j])
			if sidesNotVisibleFrom == 4 {
				if j == 0 || j == xLen-1 || i == 0 || i == yLen-1 {
					continue
				}
				invisibleTreesCount++
			}
		}
		fmt.Println(row)
	}

	score := CalculateMaxScenicScore(forest)
	return fmt.Sprintf("Part 1: %d, Part 2: %d", numOfTrees-invisibleTreesCount, score)
}

func AssessWhatTreesAreVisibleInRow(forest [][]int, visibilityMap [][]int, isFlipped bool) [][]int {
	xLen := len(forest[0])
	yLen := len(forest)
	// How many trees are NOT visible may be easier
	// A tree is NOT visible if any of the trees between it and the edge are taller than it
	// We search for the index of the highest value in the rows and columns
	forestRotated := util.BuildEmptyIntMatrix(xLen-1, yLen-1)
	for i := range forest {
		leftMax := forest[i][0]
		rightMax := forest[i][xLen-1]
		for j := range forest[i] {
			leftTreeHeight := forest[i][j]
			if leftTreeHeight <= leftMax {
				// Not visible from the left
				if isFlipped {
					visibilityMap[j][i]++
				} else {
					visibilityMap[i][j]++
				}
			}
			if leftTreeHeight > leftMax {
				leftMax = leftTreeHeight
			}

			rightTreeHeight := forest[i][xLen-1-j]
			if rightTreeHeight <= rightMax {
				// Not visible from the right
				if isFlipped {
					visibilityMap[xLen-1-j][i]++
				} else {
					visibilityMap[i][xLen-1-j]++
				}
			}
			if rightTreeHeight > rightMax {
				rightMax = rightTreeHeight
			}

			forestRotated[j][i] = leftTreeHeight
		}
	}

	return forestRotated
}

func CalculateMaxScenicScore(forest [][]int) int {
	max := 0
	xLen := len(forest[0])
	yLen := len(forest)
	for i := range forest {
		for j := range forest[i] {
			left := 0
			right := 0
			up := 0
			down := 0
			treeHeight := forest[i][j]
			for y := j - 1; y >= 0; y-- {
				left++
				if forest[i][y] >= treeHeight {
					break
				}
			}

			for y := j + 1; y < yLen; y++ {
				right++
				if forest[i][y] >= treeHeight {
					break
				}
			}

			for x := i - 1; x >= 0; x-- {
				up++
				// if we get to a larger tree we stop.
				if forest[x][j] >= treeHeight {
					break
				}
			}

			for x := i + 1; x < xLen; x++ {
				down++
				if forest[x][j] >= treeHeight {
					break
				}
			}

			score := left * right * down * up
			if score > max {
				max = score
			}
		}
	}

	return max
}
