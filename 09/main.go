package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const (
	Up    string = "U"
	Right        = "R"
	Down         = "D"
	Left         = "L"
)

type Point struct{ x, y int }

func main() {
	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}

func part1(input string) int {
	return solve(input, 2)
}

func part2(input string) int {
	return solve(input, 10)
}

func solve(input string, length int) int {
	ropes := make([]Point, length)

	for i := 0; i < len(ropes); i++ {
		ropes[i] = Point{0, 0}
	}

	lines := strings.Split(input, "\n")
	visited := make(map[Point]bool)
	visited[Point{0, 0}] = true
	head := &ropes[0]
	tail := &ropes[len(ropes)-1]

	for _, line := range lines {
		var direction string
		var distance int
		fmt.Sscanf(line, "%s %d", &direction, &distance)

		for step := 0; step < distance; step++ {
			switch direction {
			case Up:
				head.y -= 1
			case Right:
				head.x += 1
			case Down:
				head.y += 1
			case Left:
				head.x -= 1
			}

			for i := 1; i < len(ropes); i++ {
				previous := &ropes[i-1]
				current := &ropes[i]
				delta := Point{(previous.x - current.x), (previous.y - current.y)}

				if -1 <= delta.x && delta.x <= 1 && -1 <= delta.y && delta.y <= 1 {
					continue
				}

				current.x += move(delta.x)
				current.y += move(delta.y)
			}

			visited[Point{tail.x, tail.y}] = true
		}
	}

	return len(visited)
}

func move(delta int) int {
	if delta < 0 {
		return -1
	} else if delta > 0 {
		return 1
	}

	return 0
}
