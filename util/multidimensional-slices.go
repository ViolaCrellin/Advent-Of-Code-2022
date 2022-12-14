package util

import (
	"strings"
)

/*
 * Builds a multidimensional array from input like this:
 * 39 26 33 65 32
 * 78 72 80 51  0
 * 35 64 60 18 31
 * 93 59 83 54 74
 * 86  5  9 98 69
 */
func BuildIntMatrixFromString(input string, nextRowDelimiter string, nextColumnDelimiter string) [][]int {
	rows := strings.Split(input, nextRowDelimiter)
	var result = make([][]int, len(rows))
	for i, row := range rows {
		var rowStrValues []string
		if nextColumnDelimiter == " " {
			rowStrValues = strings.Fields(row)
		} else {
			rowStrValues = strings.Split(row, nextColumnDelimiter)
		}

		result[i], _ = SliceAtoi(rowStrValues)
	}

	return result
}

func BuildByteMatrixFromString(input string, nextRowDelimiter string) [][]byte {
	rows := strings.Split(input, nextRowDelimiter)
	var result = make([][]byte, len(rows))
	for i, row := range rows {
		result[i] = []byte(row)
	}

	return result
}

func MultidimensionalSum(input [][]int, positiveOnly bool) int {
	result := 0
	for _, row := range input {
		result += Sum(row, positiveOnly)
	}

	return result
}

func GetDiagonalValuesInIntMatrix(matrix [][]int, diagonal rune) []int {
	var result []int
	switch diagonal {
	case '/':
		for i := 0; i < len(matrix); i++ {
			result = append(result, matrix[i][i])
		}

	case '\\':
		for i := len(matrix); i > 0; i-- {
			result = append(result, matrix[i-1][len(matrix)-i])
		}

	}

	return result
}

func GetColumnValuesFromIntMatrix(matrix [][]int, columnNumber int) []int {
	column := make([]int, 0)
	for _, row := range matrix {
		column = append(column, row[columnNumber])
	}
	return column
}

func ListCoordinatesBetweenPoints(A []int, B []int, includeDiagonals bool) [][]int {
	result := make([][]int, 0)
	isHorizontal := A[0] == B[0]
	isVertical := A[1] == B[1]
	isDiagonal := false
	if !isHorizontal && !isVertical {
		isDiagonal = true
		if !includeDiagonals {
			return result
		}
	}

	if !isDiagonal {

		index := 0
		nonVaryingIndex := 1
		if isHorizontal {
			index = 1
			nonVaryingIndex = 0
		}
		if A[index] > B[index] {
			for i := B[index]; i <= A[index]; i++ {
				includedCoordinate := make([]int, 2)
				includedCoordinate[index] = i
				includedCoordinate[nonVaryingIndex] = A[nonVaryingIndex]
				result = append(result, includedCoordinate)
			}
		} else {
			for i := A[index]; i <= B[index]; i++ {
				includedCoordinate := make([]int, 2)
				includedCoordinate[index] = i
				includedCoordinate[nonVaryingIndex] = A[nonVaryingIndex]
				result = append(result, includedCoordinate)
			}
		}
	} else {
		indexA := 0
		indexB := 1
		if A[indexA] < B[indexA] {
			for i := A[indexA]; i <= B[indexA]; i++ {
				includedCoordinate := make([]int, 2)
				includedCoordinate[indexA] = i
				result = append(result, includedCoordinate)
			}
		} else {
			for i := A[indexA]; i >= B[indexA]; i-- {
				includedCoordinate := make([]int, 2)
				includedCoordinate[indexA] = i
				result = append(result, includedCoordinate)
			}
		}

		for x, halfDoneCoord := range result {
			if A[indexB] > B[indexB] {
				halfDoneCoord[indexB] = A[indexB] - x
			} else {
				halfDoneCoord[indexB] = A[indexB] + x
			}
			result[x] = halfDoneCoord
		}
	}

	return result
}

func FindMaxAxisValsFromCoordinatePairs(coordinatePairs [][]int) (xMax int, yMax int) {
	for _, coordinatePair := range coordinatePairs {
		x := coordinatePair[0]
		if x > xMax {
			xMax = x
		}

		y := coordinatePair[1]
		if y > yMax {
			yMax = y
		}
	}

	return
}

func BuildEmptyIntMatrix(xMax int, yMax int) [][]int {
	xMax++
	yMax++
	result := make([][]int, yMax)
	for rowNum, _ := range result {
		rowContents := make([]int, xMax)
		result[rowNum] = rowContents
	}

	return result
}

func BuildEmptyStringMatrix(xMax int, yMax int, blankChar string) [][]string {
	xMax++
	yMax++
	result := make([][]string, yMax)
	for rowNum, _ := range result {
		rowContents := make([]string, xMax)
		if blankChar != "" {
			for i := range rowContents {
				rowContents[i] = blankChar
			}
		}
		result[rowNum] = rowContents
	}

	return result
}

func BuildEmptyIntSliceMatrix(xMax, yMax, sliceLength int) [][][]int {
	xMax++
	yMax++
	result := make([][][]int, yMax)
	for rowNum, _ := range result {
		rowContents := make([][]int, xMax)
		for i := range rowContents {
			rowContents[i] = make([]int, sliceLength)
		}
		result[rowNum] = rowContents
	}
	return result
}
