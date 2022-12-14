package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"example.com/adventofcode/solutions"
)

func main() {
	dayStr := os.Args[1]

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		fmt.Println("Error")
		return
	}

	input := getPuzzleInputFromFile(day)
	input = strings.TrimSuffix(input, "\n")
	answer := Solve(day, input)
	fmt.Println(answer)

}

func Solve(day int, input string) string {
	switch day {
	case 1:
		return solutions.One(input)
	case 2:
		return solutions.Two(input)
	case 3:
		return solutions.Three(input)
	case 4:
		return solutions.Four(input)
	case 5:
		return solutions.Five(input, 2)
	case 6:
		return solutions.Six(input)
	case 7:
		return solutions.Seven(input)
	case 8:
		return solutions.Eight(input)
	case 9:
		return solutions.Nine(input)
	case 10:
		return solutions.Ten(input)
	case 11:
		return solutions.Eleven(input, 2)
	case 12:
		return solutions.Twelve(input, 2)
	case 13:
		return solutions.Thirteen(input)
	case 14:
		return solutions.Fourteen(input)
	default:
		return "Not Yet Implemented"
	}
}

func getPuzzleInputOnline(day int) string {
	url := fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	return sb
}

func getPuzzleInputFromFile(day int) string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(path)
	fileName := fmt.Sprintf("../inputs/%d.txt", day)
	input, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	return string(input)
}
