package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
	pairs := strings.Split(input, "\n\n")
  sum := 0

	for _, pair := range pairs {
		var left, right []interface{}
		packets := strings.Split(pair, "\n")
		json.Unmarshal([]byte(packets[0]), &left)
		json.Unmarshal([]byte(packets[1]), &right)

    if sorted(left, right) {
      sum++
    }
	}

	return sum
}

func part2(input string) int {
	return -1
}

func sorted(left, right []interface{}) bool {
	for i := 0; i < len(left); i++ {
    if len(right) < i {
      return false
    }
    
		leftInt, leftIsInt := left[i].(int)
		rightInt, rightIsInt := right[i].(int)

		if leftIsInt && rightIsInt {
			if leftInt < rightInt {
				return true
			}

			if leftInt > rightInt {
				return false
			}

			if leftInt == rightInt {
				continue
			}
		}

		leftList, leftIsList := left[i].([]interface{})
		rightList, rightIsList := right[i].([]interface{})

		if leftIsList && rightIsList {
			if sorted(leftList, rightList) == false {
				return false
			}

			continue
		}

    if leftIsInt && rightIsList {
      if sorted([]interface{}{leftInt}, rightList) == false {
        return false
      }

      continue
    } else if leftIsList && rightIsInt {
      if sorted(leftList, []interface{}{rightInt}) == false {
        return false
      }

      continue
    }
	}

  return false
}
