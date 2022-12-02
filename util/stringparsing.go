package util

import (
	"sort"
	"strconv"
)

func SliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func SortBytes(input string) []byte {
	inputBytes := []byte(input)
	sort.Slice(inputBytes, func(a, b int) bool {
		return inputBytes[a] < inputBytes[b]
	})
	return inputBytes
}

func BuildHistogramOfLetterOccurences(input string) map[string]int {
	result := make(map[string]int)
	for _, runeItem := range input {
		if _, ok := result[string(runeItem)]; !ok {
			result[string(runeItem)] = 0
		}
		result[string(runeItem)]++
	}

	return result
}
