package solutions

import (
	"fmt"
	"strings"

	"example.com/adventofcode/util"
)

func One(input string) string {
	perElfCalorieList := strings.Split(input, "\n\n")
	totalCaloriesPerElf := make([]int, len(perElfCalorieList))
	for i := range perElfCalorieList {
		calorieListPerElf, err := util.SliceAtoi(strings.Split(perElfCalorieList[i], "\n"))
		if err != nil {
			return err.Error()
		}

		totalCaloriesPerElf[i] = util.Sum(calorieListPerElf, true)
	}

	_, max := util.MinAndMax(totalCaloriesPerElf)
	sumOfTop3 := util.SumOfTopN(totalCaloriesPerElf, 3, true)
	return fmt.Sprintf("Answer 1: %d\n Answer 2: %d\n", max, sumOfTop3)
}
