package util

import (
	"sort"
)

func Mean(input []int) float64 {
	return float64(Sum(input, true)) / float64(len(input))
}

/*
 * Returns slice of most frequently occuring values in the input
 *
 */
func Mode(input []int) []int {
	highestFrequency := 0
	histogram := SliceCountIntValues(input)
	for _, v := range histogram {
		if v >= highestFrequency {
			highestFrequency = v
		}
	}

	modes := make([]int, 0)
	for i, v := range histogram {
		if v == highestFrequency {
			modes = append(modes, i)
		}
	}

	return modes
}

func Median(input []int) float64 {
	sort.Ints(input)
	if len(input)%2 != 0 {
		medianIndex := len(input) + 1/2
		return float64(input[medianIndex-1])
	} else {
		medianIndex := len(input) / 2
		return Mean([]int{input[medianIndex-1], input[medianIndex]})
	}
}

func LeastSquaresRegression(input [][2]int) func(*float64, *float64) *float64 {
	var sumX, sumY, sumXX, sumXY float64
	n := float64(len(input))
	for _, v := range input {
		x := v[0]
		y := v[1]
		sumX += float64(x)
		sumY += float64(y)
		sumXX += float64(x * x)
		sumXY += float64(x * y)
	}
	sumXsumY := sumX * sumY
	beta := (n*sumXY - sumXsumY) / (n*sumXX - sumXsumY)
	alpha := sumY/n - sumX*beta/n
	return func(x *float64, y *float64) *float64 {
		//Solve for X given Y
		if (x == nil && y == nil) || (x != nil && y != nil) {
			return nil
			//Solve for y given x
		} else if y == nil {
			resultY := alpha + beta*(*x)
			return &resultY
			//Solve for x given y
		} else {
			resultX := (*y - alpha) / beta
			return &resultX
		}

	}
}

func MinAndMax(input []int) (min int, max int) {
	min = input[0]
	max = input[0]
	for _, value := range input {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func SumOfTopN(input []int, n int, positiveOnly bool) (sum int) {
	sort.IntSlice(input).Sort()
	topN := input[len(input)-n:]
	return Sum(topN, positiveOnly)
}

func IndexOfMinAndMax(input []int) (minIndex int, maxIndex int) {
	min := input[0]
	minIndex = 0
	max := input[0]
	maxIndex = 0
	for i, value := range input {
		if value < min {
			min = value
			minIndex = i
		}
		if value > max {
			max = value
			maxIndex = i
		}
	}
	return minIndex, maxIndex
}
