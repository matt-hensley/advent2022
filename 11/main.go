package main

import (
	_ "embed"
	"fmt"
  "sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Operand struct {
	isOld bool
	value int
}

type Monkey struct {
	items       []int
	operation   string
	operand     Operand
	divisible   int
	ifTrue      int
	ifFalse     int
	inspections int
}

func main() {
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
	return solve(input, 20, true)
}

func part2(input string) int {
	return solve(input, 10000, false)
}

func solve(input string, rounds int, reduceWorry bool) int {
  monkeys := parse_monkeys(input)

  for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				worry := operation(item, monkey.operation, monkey.operand)

        if reduceWorry == true {
          worry = worry / 3
        }

				if worry%monkey.divisible == 0 {
					//fmt.Println(item, worry, "to", monkey.ifTrue)
					monkeys[monkey.ifTrue].items = append(monkeys[monkey.ifTrue].items, worry)
				} else {
					//fmt.Println(item, worry, "to", monkey.ifFalse)
					monkeys[monkey.ifFalse].items = append(monkeys[monkey.ifFalse].items, worry)
				}

				monkey.inspections += 1
			}

			monkey.items = []int{}
		}
	}
  
  inspections := []int{}

	for _, monkey := range monkeys {
    inspections = append(inspections, monkey.inspections)
	}
  
  fmt.Println(inspections)

  sort.Slice(inspections, func (l, r int) bool {
    return inspections[r] < inspections[l]
  })

  
	return inspections[0] * inspections[1]
}

func parse_monkeys(input string) []*Monkey {
	sections := strings.Split(input, "\n\n")
	monkeys := []*Monkey{}

	for _, section := range sections {
		//fmt.Println(section)

		var monkey Monkey
		monkey.items = []int{}
		lines := strings.Split(section, "\n")

		for _, val := range strings.Split(strings.Split(lines[1], ":")[1], ",") {
			item, _ := strconv.Atoi(strings.TrimSpace(val))
			monkey.items = append(monkey.items, item)
		}

		var operand string
		fmt.Sscanf(strings.TrimSpace(lines[2]), "Operation: new = old %s %s", &monkey.operation, &operand)

		if operand == "old" {
			monkey.operand = Operand{isOld: true, value: 0}
		} else {
			value, _ := strconv.Atoi(strings.TrimSpace(operand))
			monkey.operand = Operand{isOld: false, value: value}
		}

		fmt.Sscanf(strings.TrimSpace(lines[3]), "Test: divisible by %d", &monkey.divisible)
		fmt.Sscanf(strings.TrimSpace(lines[4]), "If true: throw to monkey %d", &monkey.ifTrue)
		fmt.Sscanf(strings.TrimSpace(lines[5]), "If false: throw to monkey %d", &monkey.ifFalse)
		monkey.inspections = 0
		monkeys = append(monkeys, &monkey)
	}

	return monkeys
}

func operation(item int, operation string, operand Operand) int {
	switch operation {
	case "+":
		if operand.isOld {
			return item + item
		} else {
			return item + operand.value
		}
	case "*":
		if operand.isOld {
			return item * item
		} else {
			return item * operand.value
		}
	}

	panic("unrecognized operation")
}
