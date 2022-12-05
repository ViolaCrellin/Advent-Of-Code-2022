package solutions

import (
	"fmt"
	"regexp"
	"strings"

	"example.com/adventofcode/util"
)

func Five(input string, part int) string {
	allInput := strings.Split(input, "\n\n")
	towersPlanRows := strings.Split(allInput[0], "\n")
	// TODO: I feel like 8 is a *bit* (ahahah, I crack myself up - in both senses) too significant
	towers := make([]util.Stack, 9)
	for j := 0; j <= 8; j++ {
		for i := 7; i >= 0; i-- {
			letter := string(towersPlanRows[i][j])
			if string(towersPlanRows[i][j]) != "." {
				towers[j].Push(letter)
			}
		}
	}

	// towers are now in 8 stacks we can push and pop from
	instructions := strings.Split(allInput[1], "\n")

	for i := range instructions {
		r := regexp.MustCompile(`move\s(?P<qty>\d+)\sfrom\s(?P<fromTower>\d+)\sto\s(?P<toTower>\d+)`)
		match := r.FindStringSubmatch(instructions[i])
		instruction, err := util.GetRegexMapOfNamedCaptureGroupIntValues(r, match)
		if err != nil {
			fmt.Println(fmt.Errorf("something looks funky with this instruction %s", instructions[i]))
			return ""
		}

		if part == 1 {
			towers = MoveOneByOne(towers, instruction)
		} else {
			towers = MoveInChunks(towers, instruction)
		}
	}

	result := ""
	for i := range towers {
		result += fmt.Sprintf("%s", towers[i].TopVal())
	}

	return result
}

func MoveOneByOne(towers []util.Stack, instruction map[string]int) []util.Stack {
	for j := 0; j < instruction["qty"]; j++ {
		fromTower := instruction["fromTower"] - 1
		toTower := instruction["toTower"] - 1

		toPush, ok := towers[fromTower].Pop()
		if !ok {
			fmt.Println(fmt.Errorf("ran out of blocks to pop from on tower %d (index %d)", instruction["fromTower"], fromTower))
		}

		towers[toTower].Push(toPush)
	}

	return towers
}

func MoveInChunks(towers []util.Stack, instruction map[string]int) []util.Stack {
	spareTower := util.Stack{}
	for j := 0; j < instruction["qty"]; j++ {
		fromTower := instruction["fromTower"] - 1
		toPush, ok := towers[fromTower].Pop()
		if !ok {
			fmt.Println(fmt.Errorf("ran out of blocks to pop from on tower %d (index %d)", instruction["fromTower"], fromTower))
		}

		spareTower.Push(toPush)
	}

	for j := 0; j < instruction["qty"]; j++ {
		toTower := instruction["toTower"] - 1
		toPush, ok := spareTower.Pop()
		if !ok {
			fmt.Println("ran out of blocks to pop from from spare tower")
		}

		towers[toTower].Push(toPush)
	}

	return towers
}
