package solutions

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"example.com/adventofcode/util"
)

type monkey struct {
	items           []int
	operation       operation
	testPredicate   operation
	testResult      map[bool]int
	inspectionCount int
	number          int
}

type operation struct {
	operator string
	LHS      int
	RHS      int
}

func Eleven(input string, part int) string {
	monkeyData := strings.Split(input, "\n\n")
	monkeys := make([]*monkey, len(monkeyData))
	for i := range monkeyData {
		monkeys[i] = buildMonkey(monkeyData[i])
		monkeys[i].number = i
	}

	commonMultiplier := 1
	for _, m := range monkeys {
		commonMultiplier *= m.testPredicate.RHS
	}

	round := 0
	rounds := 20
	if part == 2 {
		rounds = 10000
	}

	for round < rounds {
		fmt.Printf("\n ROUND %d", round)
		for i := range monkeys {
			items := monkeys[i].items
			if len(items) == 0 {
				continue
			}
			monkeys[i].inspectionCount += len(items)
			for j := range items {
				worryLevel := items[j]
				operation := monkeys[i].operation
				if operation.LHS < 0 {
					operation.LHS = worryLevel
				}
				if operation.RHS < 0 {
					operation.RHS = worryLevel
				}
				worryLevel = operation.Evaluate()
				if part == 1 {
					worryLevel /= 3
				} else {
					worryLevel %= commonMultiplier
				}
				test := monkeys[i].testPredicate
				test.LHS = worryLevel
				recipientMonkey := monkeys[i].testResult[test.EvaluatePredicate()]
				monkeys[recipientMonkey].items = append(monkeys[recipientMonkey].items, worryLevel)
			}
			monkeys[i].items = []int{}
		}

		// for i := range monkeys {
		// 	fmt.Printf("\n%s", monkeys[i].ToString())
		// }
		round++
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectionCount > monkeys[j].inspectionCount
	})

	monkeyBusiness := monkeys[0].inspectionCount * monkeys[1].inspectionCount
	return fmt.Sprintf("Part 1: %d", monkeyBusiness)
}

func (m *monkey) ToString() string {
	return fmt.Sprintf("\nMonkey %d\nItems: %+v\nInspectionCount:%d", m.number, m.items, m.inspectionCount)
}

func (o operation) Evaluate() int {
	switch o.operator {
	case "*":
		return o.LHS * o.RHS
	case "+":
		return o.LHS + o.RHS
	case "-":
		return o.LHS - o.RHS
	case "/":
		return o.LHS / o.RHS
	default:
		return -1
	}
}

func (o operation) EvaluatePredicate() bool {
	switch o.operator {
	case "divisible by":
		return o.LHS%o.RHS == 0
	default:
		return false
	}
}

// Is this like a code monkey but they just deal with the build? Like Ralph.
func buildMonkey(monkeyData string) *monkey {
	// Muwahahahah I am the regex master. Bow to your liege.
	r := regexp.MustCompile(`^Monkey\s\d+\:\s+Starting\sitems:\s(?P<items>[\d+,\s]+)\n\s+Operation:\snew\s\=\s(?P<operationLHS>.*)\s(?P<operationOperator>[\+|\/|\*|\-])\s(?P<operationRHS>.+)\s+Test\:\s(?P<testPredicateOperator>[\w\s]+)\s(?P<testPredicateRHS>\d+)\n\s+If\strue:\sthrow\sto\smonkey\s(?P<testResultTrue>\d+)\n\s+If\sfalse:\sthrow\sto\smonkey\s(?P<testResultFalse>\d+)$`)
	match := r.FindStringSubmatch(monkeyData)
	monkeyStuff := util.GetRegexMapOfNamedCaptureGroupValues(r, match)
	monkey := monkey{
		testResult: map[bool]int{
			true:  0,
			false: 0,
		},
	}
	// Go anti-pattern of the week. Ranging over a map with a switch statement.
	for key, val := range monkeyStuff {
		switch key {
		case "items":
			items, err := util.SliceAtoi(strings.Split(val, ", "))
			if err != nil {
				break
			}
			monkey.items = items
		case "operationLHS":
			lhs, err := strconv.Atoi(val)
			if err != nil {
				lhs = -1
			}
			monkey.operation.LHS = lhs
		case "operationRHS":
			rhs, err := strconv.Atoi(val)
			if err != nil {
				rhs = -1
			}
			monkey.operation.RHS = rhs
		case "operationOperator":
			monkey.operation.operator = val
		case "testPredicateOperator":
			monkey.testPredicate.operator = val
		case "testPredicateRHS":
			rhs, err := strconv.Atoi(val)
			if err != nil {
				rhs = -1
			}
			monkey.testPredicate.RHS = rhs
		case "testResultTrue":
			trueMonkey, err := strconv.Atoi(val)
			if err != nil {
				break
			}
			monkey.testResult[true] = trueMonkey
		case "testResultFalse":
			falseMonkey, err := strconv.Atoi(val)
			if err != nil {
				break
			}
			monkey.testResult[false] = falseMonkey
		}
	}

	return &monkey
}
