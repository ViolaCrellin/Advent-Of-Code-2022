package solutions

import (
	"fmt"
	"strings"

	"example.com/adventofcode/util"
)

type instruction struct {
	name   string
	val    int
	cycles int
}

func Ten(input string) string {
	rawInstructions := strings.Split(input, "\n")
	instructions := buildInstructions(rawInstructions)
	CRT := mapSpriteLocations(instructions)
	return fmt.Sprintf("Part 1: %d \n\n %s", findSumOfSignalStrengths(instructions), drawCRT(CRT))
}

func buildInstructions(rawInstructions []string) []instruction {
	instructions := make([]instruction, len(rawInstructions))
	for i := range instructions {
		if rawInstructions[i] == "noop" {
			instructions[i] = instruction{
				cycles: 1,
				val:    0,
				name:   "noop",
			}
			continue
		}

		letter, val, err := util.SplitLetterAndNumber(rawInstructions[i], " ")
		if err != nil || letter != "addx" {
			fmt.Println("shit gone bang")
		}
		instructions[i] = instruction{
			cycles: 2,
			val:    val,
			name:   "addx",
		}
	}
	return instructions
}

func findSumOfSignalStrengths(instructions []instruction) int {
	i := 0
	X := 1
	sum := 0
	for cycle := 1; cycle <= 220; cycle++ {

		if (cycle-20)%40 == 0 {
			sum += X * cycle
		}

		X, i = process(&instructions[i], X, i)
	}

	return sum
}

func process(instruction *instruction, X, i int) (int, int) {
	switch instruction.name {
	case "addx":
		// lets make the instruction responsible for knowing what cycle its on
		instruction.cycles--
		// if its count down is complete, add the value
		if instruction.cycles == 0 {
			X += instruction.val
			i++
		}
	case "noop":
		// go next order
		i++
	}
	return X, i
}

func mapSpriteLocations(instructions []instruction) [6][40]string {
	CRT := [6][40]string{}
	for i, rows := range CRT {
		for j := range rows {
			CRT[i][j] = "."
		}
	}

	// C for center.
	C := 1
	i := 0
	for cycle := 1; i < len(instructions); cycle++ {

		// calculate which pixel is being drawn... ZERO INDEXED
		Y := (cycle - 1) / 40
		X := (cycle - 1) % 40

		// see if the spite's horizontal location overlaps that pixelCol
		spriteLeft, spriteRight := C-1, C+1
		if spriteLeft <= X && spriteRight >= X {
			CRT[Y][X] = "#"
		}

		C, i = process(&instructions[i], C, i)
	}
	return CRT
}

func drawCRT(CRT [6][40]string) string {
	log := ""
	for i := range CRT {
		for j := range CRT[i] {
			log += CRT[i][j]
		}
		log += "\n"
	}

	return log
}
