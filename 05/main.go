package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1: ", part1(input))
  fmt.Println("Part 2: ", part2(input))
}

func part1(input string) string {
  return solve(input, false)
}

func part2(input string) string {
  return solve(input, true)
}

func solve(input string, multimove bool) string {
  fmt.Println("multimove", multimove)
  parts := strings.Split(input, "\n\n")
  stacks, keys := extract_crates(parts[0])
  procedures := strings.Split(parts[1], "\n")
  re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
  printmap(stacks)

  for _, procedure := range procedures {
    matches := re.FindStringSubmatch(procedure)
    moves, err := strconv.Atoi(matches[1])

    if err != nil {
      panic(err)
    }
    
    from, to := matches[2], matches[3]

    if multimove == true {
      val := stacks[from][0:moves]
      targets := make([]string, len(val))
      copy(targets, val)
      stacks[from] = stacks[from][moves:]
      stacks[to] = append(targets, stacks[to]...)
    } else {
      for i := 0; i < moves; i++ {
        if len(stacks[from]) == 0 {
          continue
        }
  
        target := stacks[from][0]
        stacks[from] = stacks[from][1:]
        stacks[to] = append([]string{target}, stacks[to]...)
      }
    }
  }

  solution := make([]string, 0)

  for _, key := range keys {
    stack := stacks[key]

    if len(stack) > 0 {
      solution = append(solution, stack[0])
    }
  }

  return strings.Join(solution, "")
}

func extract_crates(text string) (map[string][]string, []string) {
  data := make(map[string][]string)
  lines := strings.Split(text, "\n")
  last := lines[len(lines) - 1]
  stacks := lines[0:len(lines) - 1]
  keys := make([]string, 0)

  for i, c := range strings.Split(last, "") {
    if c == "" || c == " " {
      continue
    }

    keys = append(keys, c)

    for _, line := range stacks {
      cell := rune(line[i])
      
      if unicode.IsSpace(cell) {
        continue
      }

      data[c] = append(data[c], string(cell))
    }
  }

  return data, keys
}

func printmap(data map[string][]string) {
  for key, values := range data {
    fmt.Print(key, "->")

    for _, v := range values {
      fmt.Print(v)
    }

    fmt.Print("\n")
  }
}

