package solutions

import (
	"fmt"
	"strings"

	"example.com/adventofcode/util"
)

func Three(input string) string {
	rucksacks := strings.Split(input, "\n")
	alphabet := util.AlphabetMap()
	scoreA := 0
	scoreB := 0
	threeChunk := [][]string{}
	for i := range rucksacks {
		rucksack := rucksacks[i]
		compartmentSize := len(rucksack) / 2
		all := strings.Split(rucksack, "")
		threeChunk = append(threeChunk, all)
		compartmentA := strings.Split(rucksack[:compartmentSize], "")
		compartmentB := strings.Split(rucksack[compartmentSize:], "")

		result := util.IntersectionString(compartmentA, compartmentB)
		incorrectItem := result[0]
		scoreA += addPriorityScore(alphabet, incorrectItem)
		if len(threeChunk) == 3 {
			commonItems := util.IntersectionString(threeChunk[0], threeChunk[1])
			commonItems = util.IntersectionString(commonItems, threeChunk[2])
			badge := commonItems[0]
			threeChunk = [][]string{}
			scoreB += addPriorityScore(alphabet, badge)
		}

	}
	return fmt.Sprintf("Part 1: %d\nPart 2: %d", scoreA, scoreB)
}

func addPriorityScore(alphabet map[string]int, item string) int {
	itemScore, ok := alphabet[item]
	if !ok {
		lowerCase := strings.ToLower(item)
		itemScore = alphabet[lowerCase] + 26
	}
	return itemScore
}
