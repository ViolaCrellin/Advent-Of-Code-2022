package solutions

import (
	"fmt"
	"strings"
)

var (
	outcomeScores = map[string]int{
		"lose": 0,
		"draw": 3,
		"win":  6,
	}

	requiredOutcomes = map[string]string{
		"X": "lose",
		"Y": "draw",
		"Z": "win",
	}

	//A rock
	//B paper
	//C scissors

	outcomes = map[string]map[string]string{
		"win": {
			"A": "Y",
			"B": "Z",
			"C": "X",
		},
		"lose": {
			"A": "Z",
			"B": "X",
			"C": "Y",
		},
		"draw": {
			"A": "X",
			"B": "Y",
			"C": "Z",
		},
	}

	weaponScores = map[string]int{
		"X": 1, //Rock
		"Y": 2, //Paper
		"Z": 3, //Scissors
	}

	flatOutcomes = map[string]string{
		"A Y": "win",
		"A Z": "lose",
		"A X": "draw",
		"B Z": "win",
		"B X": "lose",
		"B Y": "draw",
		"C X": "win",
		"C Y": "lose",
		"C Z": "draw",
	}
)

func Two(input string) string {
	matches := strings.Split(input, "\n")

	return fmt.Sprintf("Part 1: %d\nPart 2: %d", Part1(matches), Part2(matches))
}

func Part1(matches []string) int {
	score := 0
	for i := range matches {
		match := matches[i]
		outcome := flatOutcomes[match]
		score += outcomeScores[outcome]
		score += weaponScores[match[len(match)-1:]]
	}

	return score
}

func Part2(matches []string) int {
	score := 0
	for i := range matches {
		match := matches[i]
		requiredOutcome := requiredOutcomes[match[len(match)-1:]]
		theirWeapon := match[:1]
		requiredWeapon := outcomes[requiredOutcome][theirWeapon]
		score += outcomeScores[requiredOutcome]
		score += weaponScores[requiredWeapon]
	}

	return score
}
