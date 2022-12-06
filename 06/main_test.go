package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

type TestCase struct {
  input string
  want int
}

func Test_part1(t *testing.T) {
  test_cases := []TestCase {
    {"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
    {"nppdvjthqldpwncqszvftbrmjlhg", 6},
    {"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
    {"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
  }

  for _, test_case := range test_cases {
    want := test_case.want
    got := part1(test_case.input)

    if got != want {
		  t.Errorf("part1_sample = %d; want %d", got, want)
	  }    
  }
}

func Test_part2(t *testing.T) {
  test_cases := []TestCase {
    {"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
    {"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
    {"nppdvjthqldpwncqszvftbrmjlhg", 23},
    {"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
    {"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
  }

  for _, test_case := range test_cases {
    want := test_case.want
    got := part1(test_case.input)

    if got != want {
		  t.Errorf("part2_sample = %d; want %d", got, want)
	  }    
  }
}