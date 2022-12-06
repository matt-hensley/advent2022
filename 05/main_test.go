package main

import (
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sample string

func Test_part1(t *testing.T) {
	want := "CMZ"
	got := part1(sample)
	if got != want {
		t.Errorf("part1_sample = %s; want %s", got, want)
	}
}

func Test_part2(t *testing.T) {
	want := "MCD"
	got := part2(sample)
	if got != want {
		t.Errorf("part2_sample = %s; want %s", got, want)
	}
}
