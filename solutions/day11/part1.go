package day11

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) int {
	return helper(input, 25)
}

func helper(input string, blinkCount int) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	stoneCount := make(map[int]int)
	for scanner.Scan() {
		for _, val := range strings.Fields(scanner.Text()) {
			numInt, _ := strconv.Atoi(val)
			stoneCount[numInt]++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	stoneMemo := make(map[int][]int)
	for i := 0; i < blinkCount; i++ {
		stoneCount = blink(stoneCount, stoneMemo)
	}
	var score int
	for _, count := range stoneCount {
		score += count
	}

	return score
}

func blink(stoneCount map[int]int, stoneMemo map[int][]int) map[int]int {
	newStoneCount := make(map[int]int)
	for stone, count := range stoneCount {
		result := transform(stone, stoneMemo)
		for _, res := range result {
			newStoneCount[res] += count
		}

	}
	return newStoneCount
}

func transform(stone int, stoneMemo map[int][]int) []int {
	if res, ok := stoneMemo[stone]; ok {
		return res
	}

	valStr := strconv.Itoa(stone)
	length := len(valStr)

	var results []int
	switch {
	case stone == 0:
		results = []int{1}
	case length%2 == 0:
		mid := length / 2
		numA, _ := strconv.Atoi(valStr[:mid])
		numB, _ := strconv.Atoi(valStr[mid:])
		results = []int{numA, numB}
	default:
		results = []int{stone * 2024}
	}

	stoneMemo[stone] = results
	return results

}
