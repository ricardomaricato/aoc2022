package day02

import (
	"os"
	"testing"
)

// A e X - Pedra - 1 ponto
// B e Y - Papel - 2 pontos
// C e Z - Tesoura - 3 pontos

// 0 — Perdeu
// 3 — Empate
// 6 — Ganhou

var p = &Puzzle{}
var exampleInput = `A Y
B X
C Z
`

func TestPart1(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "15"},
	}

	for i, tt := range tests {
		solution := p.Part1(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part1(string(data))
	if solution != "14827" {
		t.Errorf("Solution for Part1: %s", solution)
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{exampleInput, "12"},
	}

	for i, tt := range tests {
		solution := p.Part2(tt.Input)
		if solution != tt.Expected {
			t.Errorf("[%02d] Expected %q, got %q", i, tt.Expected, solution)
		}
	}

	data, _ := os.ReadFile("input.txt")
	solution := p.Part2(string(data))
	if solution != "13889" {
		t.Errorf("Solution for Part2: %s", solution)
	}
}
