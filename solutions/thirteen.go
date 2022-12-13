package solutions

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"example.com/adventofcode/util"
	"github.com/benpate/convert"
)

func Thirteen(input string) string {
	signalPairs := strings.Split(input, "\n\n")
	areOrdered := make([]int, len(signalPairs))
	allSignals := []interface{}{}
	for i := range signalPairs {
		pair := strings.Split(signalPairs[i], "\n")
		left := util.ParseStringIntSliceRepresentationJSON(pair[0])
		right := util.ParseStringIntSliceRepresentationJSON(pair[1])
		areOrdered[i] = IsOrdered(left, right)
		if areOrdered[i] == 1 {
			allSignals = append(allSignals, left)
			allSignals = append(allSignals, right)
		} else {
			allSignals = append(allSignals, right)
			allSignals = append(allSignals, left)
		}
	}

	extra2 := util.ParseStringIntSliceRepresentationJSON("[[2]]")
	extra6 := util.ParseStringIntSliceRepresentationJSON("[[6]]")
	allSignals = append(allSignals, extra2)
	allSignals = append(allSignals, extra6)
	sort.Slice(allSignals, func(i, j int) bool {
		return IsOrdered(convert.SliceOfInterface(allSignals[i]), convert.SliceOfInterface(allSignals[j])) == 1
	})

	index2 := 0
	index6 := 0
	for i := range allSignals {
		signal := allSignals[i]
		signalStr := fmt.Sprintf("%+v", signal)
		if signalStr == fmt.Sprintf("%+v", extra2) {
			index2 = i + 1
		}
		if signalStr == fmt.Sprintf("%+v", extra6) {
			index6 = i + 1
		}
		fmt.Println(signalStr)
	}

	indicesSum := 0
	orderedIndexes := []int{}
	for i := range areOrdered {
		if areOrdered[i] == 1 {
			indicesSum += i + 1
			orderedIndexes = append(orderedIndexes, i+1)
		}
	}
	fmt.Printf("\n\n ORDERED INDEXES: %+v", orderedIndexes)

	return fmt.Sprintf("\nPart 1: %d, Part 2: %d", indicesSum, index2*index6)
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsOrdered(left []interface{}, right []interface{}) int {
	for i, leftItem := range left {
		if i > len(right)-1 {
			return 0
		}
		rightItem := right[i]
		if IsFloat64(leftItem) && IsFloat64(rightItem) {
			if convert.Int(leftItem) < convert.Int(rightItem) {
				return 1
			}
			if convert.Int(leftItem) > convert.Int(rightItem) {
				return 0
			}

			// They are equal, try to get a comparison on the next element
			continue
		}

		if IsSlice(leftItem) && IsSlice(rightItem) {
			if len(convert.SliceOfInterface(leftItem)) == 0 {
				if len(convert.SliceOfInterface(rightItem)) == 0 {
					continue
				}
				return 1
			}

			return IsOrdered(convert.SliceOfInterface(leftItem), convert.SliceOfInterface(rightItem))
		}

		if IsSlice(leftItem) && IsFloat64(rightItem) {
			return IsOrdered(convert.SliceOfInterface(leftItem), convert.SliceOfInterface([]interface{}{rightItem}))
		}

		if IsFloat64(leftItem) && IsSlice(rightItem) {
			return IsOrdered(convert.SliceOfInterface([]interface{}{leftItem}), convert.SliceOfInterface(rightItem))
		}
	}

	return 1
}

func IsInt(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Int
}

func IsFloat64(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Float64
}
