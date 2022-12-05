package solutions

import (
	"fmt"
	"strings"

	"example.com/adventofcode/util"
)

func Four(input string) string {
	pairSectionList := strings.Split(input, "\n")
	fullyOverlappingCount := 0
	partiallyOverlappingCount := 0
	for i := range pairSectionList {
		sections := strings.Split(pairSectionList[i], ",")
		startA, endA := getStartAndEnd(sections[0])
		startB, endB := getStartAndEnd(sections[1])

		// A is fully enclosed by B ir B is fully enclosed by A
		if (startA >= startB && endA <= endB) || (startB >= startA && endB <= endA) {
			fullyOverlappingCount++
		}

		if (endB >= startA && endB <= endA) || (startB <= endA && startB >= startA) {
			partiallyOverlappingCount++
			continue
		}

		if (endA >= startB && endA <= endB) || (startA <= endB && startA >= startB) {
			partiallyOverlappingCount++
		}
	}
	return fmt.Sprintf("Part 1:%d\n Part2:%d\n", fullyOverlappingCount, partiallyOverlappingCount)
}

func getStartAndEnd(input string) (int, int) {
	startAndEnd, _ := util.SliceAtoi(strings.Split(input, "-"))
	return startAndEnd[0], startAndEnd[1]
}
