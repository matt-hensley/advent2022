package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string
//go:embed output.txt
var output string

func Test_part1_1(t *testing.T) {
	want := 13140
	got := part1(sample)

	if got != want {
		t.Errorf("part1_sample = %d; want %d", got, want)
	}
}

//*
func Test_part2(t *testing.T) {
	want := output
	got := part2(sample)
	if got != want {
		t.Errorf("part2_sample = \n%s; want = \n%s", got, want)
	}
}
//*/
