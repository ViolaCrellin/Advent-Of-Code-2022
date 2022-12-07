package util

import (
	"math"
	"sort"
)

func Intersection(a, b []int) []int {
	m := make(map[int]uint8)
	for _, k := range a {
		m[k] |= (1 << 0)
	}
	for _, k := range b {
		m[k] |= (1 << 1)
	}

	var inAAndB []int
	for k, v := range m {
		a := v&(1<<0) != 0
		b := v&(1<<1) != 0
		if a && b {
			inAAndB = append(inAAndB, k)
		}
	}
	return inAAndB
}

func Sum(input []int, positiveOnly bool) int {
	result := 0
	for _, v := range input {
		if positiveOnly && v < 0 {
			continue
		}
		result += v
	}

	return result
}

func Product(input []int, positiveOnly bool) int {
	result := 1
	for _, v := range input {
		if positiveOnly && v < 0 {
			continue
		}
		result *= v
	}

	return result
}

func SumSquares(input []int) int {
	result := 0
	for _, v := range input {
		result += int(math.Pow(float64(v), 2))
	}

	return result
}

/*
 * An equivalent of php's array_count_values. A histogram of sorts
 *
 */
func SliceCountIntValues(input []int) []int {
	result := make(map[int]int)
	sizeNeeded := 0
	for _, v := range input {
		if _, ok := result[v]; !ok {
			result[v] = 0
			if v > sizeNeeded {
				sizeNeeded = v
			}
		}
		result[v]++
	}

	actualResult := make([]int, sizeNeeded+1)
	for value, frequency := range result {
		actualResult[value] = frequency
	}
	return actualResult
}

func SliceCountHighFrequencyOfIntValues(input []int) map[int]uint64 {
	sortedInput := input
	sort.Ints(sortedInput)
	result := make(map[int]uint64)
	for _, v := range sortedInput {
		if _, ok := result[v]; !ok {
			result[v] = 0
		}
		result[v]++
	}

	return result
}

func FillMapKeys(keys []interface{}, value interface{}) map[interface{}]interface{} {
	resultMap := make(map[interface{}]interface{}, 0)
	for _, key := range keys {
		resultMap[key] = value
	}

	return resultMap
}

func MapValues(elements map[interface{}]interface{}) []interface{} {
	i, vals := 0, make([]interface{}, len(elements))
	for _, val := range elements {
		vals[i] = val
		i++
	}
	return vals
}

func MapValuesInt(elements map[string]int) []int {
	i, vals := 0, make([]int, len(elements))
	for _, val := range elements {
		vals[i] = val
		i++
	}
	return vals
}

func InIntSlice(elements []int, element int) bool {
	for i := range elements {
		if elements[i] == element {
			return true
		}
	}
	return false
}

func InStringSlice(elements []string, element string) bool {
	for i := range elements {
		if elements[i] == element {
			return true
		}
	}

	return false
}

func GetByteKeyForIntValue(elements map[byte]int, givenValue int) byte {
	for key, value := range elements {
		if value == givenValue {
			return key
		}
	}
	return '0'
}

// Return the slice that results from removing elements in second from the first.
func Difference(first, second []string) []string {
	var diff = make(map[string]bool)
	var out = make([]string, 0)

	if len(first) == 0 {
		return out
	} else if len(second) == 0 {
		return first
	}

	for i := range second {
		item := second[i]
		diff[item] = true
	}

	for i := range first {
		item := first[i]
		if _, ok := diff[item]; !ok {
			out = append(out, item)
		}
	}

	return out
}

func IntersectionString(a, b []string) (c []string) {
	m := make(map[string]bool)

	for i := range a {
		m[a[i]] = true
	}

	for i := range b {
		item := b[i]
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func SumMapValues(input map[string]int, positiveOnly bool) int {
	result := 0
	for _, v := range input {
		if positiveOnly && v < 0 {
			continue
		}
		result += v
	}

	return result
}
