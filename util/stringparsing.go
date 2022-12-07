package util

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"unicode"
)

func SliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for i := range sa {
		x, err := strconv.Atoi(sa[i])
		if err != nil {
			return si, err
		}
		si[i] = x
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

func GetRegexMapOfNamedCaptureGroupValues(regex *regexp.Regexp, match []string) map[string]string {
	paramsMap := make(map[string]string)
	for i, name := range regex.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}

	return paramsMap
}

func GetRegexMapOfNamedCaptureGroupIntValues(regex *regexp.Regexp, match []string) (map[string]int, error) {
	paramsMap := make(map[string]int)
	for i, name := range regex.SubexpNames() {
		if i > 0 && i <= len(match) {
			integerVal, err := strconv.Atoi(match[i])
			if err != nil {
				return nil, fmt.Errorf("expected to capture an int, did not: %s", match[i])
			}
			paramsMap[name] = integerVal
		}
	}

	return paramsMap, nil
}
