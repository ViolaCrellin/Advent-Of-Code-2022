package util

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
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

func AlphabetSlice() []string {
	a := 'a'
	z := 'z'
	result := make([]string, 26)
	idx := 0
	for i := a; i <= z; i++ {
		result[idx] = fmt.Sprintf("%c", i)
		idx++
	}

	return result
}

func AlphabetMap() map[string]int {
	a := 'a'
	z := 'z'
	result := make(map[string]int)
	idx := 0
	for i := a; i <= z; i++ {
		idx++
		result[fmt.Sprintf("%c", i)] = idx
	}

	return result
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
