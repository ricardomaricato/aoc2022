package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//[]string{
//	"1000\n2000\n3000",
//	"4000"
//}

//[]string{
//	"1000",
//	"2000",
//  "3000",
//}

//[][]int{
//	{1000, 2000, 3000},
//	{4000}
//}

type Puzzle struct {
}

type ElfsLists [][]int

func parseInput(raw string) (ElfsLists, error) {
	elfLists := strings.Split(raw, "\n\n")

	result := ElfsLists{}
	for _, elfList := range elfLists {
		itemStr := strings.Split(elfList, "\n")

		var elfItems []int
		for _, itemStr := range itemStr {
			itemInt, err := strconv.Atoi(itemStr)
			if err != nil {
				return nil, err
			}
			elfItems = append(elfItems, itemInt)
		}
		result = append(result, elfItems)
	}
	return result, nil
}

func (p *Puzzle) Part1(input string) string {
	elfLists, _ := parseInput(input)

	maxSum := 0
	for _, elfList := range elfLists {
		elfCalSum := 0
		for _, item := range elfList {
			elfCalSum += item
		}

		if elfCalSum > maxSum {
			maxSum = elfCalSum
		}
	}

	return fmt.Sprint(maxSum)
}

func (p *Puzzle) Part2(input string) string {
	elfLists, _ := parseInput(input)

	var listSum []int
	for _, elfList := range elfLists {
		elfCalSum := 0
		for _, item := range elfList {
			elfCalSum += item
		}
		listSum = append(listSum, elfCalSum)
	}

	sort.Ints(listSum)
	L := len(listSum)
	sumTop3 := listSum[L-1] + listSum[L-2] + listSum[L-3]

	return fmt.Sprint(sumTop3)
}

func (p *Puzzle) Notes() string {
	return ""
}
