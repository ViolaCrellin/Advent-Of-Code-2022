package solutions

import (
	"fmt"
	"reflect"
	"strings"

	"example.com/adventofcode/util"
	"github.com/benpate/convert"
)

func Thirteen(input string) string {
	signalPairs := strings.Split(input, "\n\n")
	areOrdered := make([]int, len(signalPairs))
	for i := range signalPairs {
		pair := strings.Split(signalPairs[i], "\n")
		left := util.ParseStringIntSliceRepresentationJSON(pair[0])
		right := util.ParseStringIntSliceRepresentationJSON(pair[1])
		//fmt.Printf("\n LEFT %+v", left)
		//fmt.Printf("\n RIGHT %+v", right)
		areOrdered[i] = IsOrdered(left, right)
		fmt.Printf("\n ORDERED %d", areOrdered[i])
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
	return fmt.Sprintf("\nPart 1: %d", indicesSum)
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsOrdered(left []interface{}, right []interface{}) int {
	fmt.Printf("\n COMPARING LEFT %+v VS RIGHT %+v", left, right)
	// if len(left) == 0 && len(right) > 0 {
	// 	return 1
	// }

	for i, leftItem := range left {
		if i > len(right)-1 {
			return 0
		}
		rightItem := right[i]
		fmt.Printf("\n COMPARING LEFT ITEM %+v VS RIGHT ITEM %+v", leftItem, rightItem)
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
			// if isOrdered == 0 {
			// 	return 0
			// }
		}

		if IsSlice(leftItem) && IsFloat64(rightItem) {
			return IsOrdered(convert.SliceOfInterface(leftItem), convert.SliceOfInterface([]interface{}{rightItem}))
			// if isOrdered == 0 {
			// 	return 0
			// }
		}

		if IsFloat64(leftItem) && IsSlice(rightItem) {
			return IsOrdered(convert.SliceOfInterface([]interface{}{leftItem}), convert.SliceOfInterface(rightItem))
			// if isOrdered == 0 {
			// 	return 0
			// }
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
