package solutions

import "fmt"

func Six(input string) string {
	charWindow := input[:1]
	requiredLength := 4
	answer := ""
	for i := range input {
		if i == 0 {
			continue
		}
		char := input[i]
		for j := range charWindow {
			if charWindow[j] == char {
				charWindow = charWindow[j+1:]
				break
			}
		}
		charWindow += string(char)
		if len(charWindow) == requiredLength {
			if requiredLength != 14 {
				answer += fmt.Sprintf("Part1: %d letter sequence: %s, first position: %d", requiredLength, charWindow, i+1)
				requiredLength = 14
			} else {
				answer += fmt.Sprintf("Part2: %d letter sequence: %s, first position %d", requiredLength, charWindow, i+1)
			}
		}
	}

	return answer
}
