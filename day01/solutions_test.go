package day01

import (
	"os"
	"testing"
)

var p = &Puzzle{}
var exampleInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "24000"},
	}

	for i, tt := range tests {
		solution, _ := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution, _ := p.Part1(string(data))
	if solution != "72017" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "45000"},
	}

	for i, tt := range tests {
		solution, _ := p.Part2(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution, _ := p.Part2(string(data))
	if solution != "212520" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}
