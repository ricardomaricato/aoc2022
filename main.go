package main

import (
	"fmt"
	"github.com/ricardomaricato/aoc2022/day01"
	"github.com/ricardomaricato/aoc2022/day02"
	"os"
	"strconv"
	"time"
)

var (
	Repeats         = 10
	UseCachedResult = false
)

func init() {
	if repeatsStr := os.Getenv("AOC_REPEATS"); repeatsStr != "" {
		repeats, err := strconv.Atoi(repeatsStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "AOC_REPEATS could not be parsed into int: %v\n", err)
		} else {
			Repeats = repeats
		}
	}

	if useCachedStr := os.Getenv("AOC_USE_CACHED_RESULT"); useCachedStr != "" {
		useCached, err := strconv.ParseBool(useCachedStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "AOC_USE_CACHED_RESULT could not be parsed into bool: %v\n", err)
		} else {
			UseCachedResult = useCached
		}
	}

}

type DayAOC interface {
	Part1(input string) string
	Part2(input string) string
	Notes() string
}

type PuzzleStats struct {
	Results struct {
		Part1 string
		Part2 string
	}
	Timing struct {
		Part1 time.Duration
		Part2 time.Duration
	}
}

func main() {
	fmt.Println("######################################################")
	fmt.Println("Resolvendo os puzzles do https://adventofcode.com/2022")
	fmt.Println("######################################################")

	fmt.Println()
	fmt.Println(" Dia | Parte 1    ( time ms ) | Parte 2    ( time ms ) |")
	fmt.Println("-----+------------------------+------------------------+")

	aocDays := []DayAOC{
		// existing days:
		&day01.Puzzle{},
		&day02.Puzzle{},
	} // ci:addNewDay
	for day0, aocDay := range aocDays {
		day := day0 + 1
		puzzleStats, err := buildPuzzleStats(day, aocDay)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error building PuzzleStats for day %d: %v\n", day, err)
			continue
		}

		// Printing results
		fmt.Printf("  %02d | %-10s (%9.1f) | %-10s (%9.1f) | %s\n", day,
			puzzleStats.Results.Part1, float64(puzzleStats.Timing.Part1/time.Microsecond)/1000,
			puzzleStats.Results.Part2, float64(puzzleStats.Timing.Part2/time.Microsecond)/1000,
			aocDay.Notes(),
		)
	}
}
