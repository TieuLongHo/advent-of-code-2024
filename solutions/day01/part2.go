package day01

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part2() int {
	input := utils.ReadFile("input/day01.txt")
	scanner := bufio.NewScanner(strings.NewReader(input))

	slice1 := make([]int, 0)
	slice2 := make([]int, 0)
	map1 := make(map[int]int)
	for scanner.Scan() {
		splitAndCountString(scanner.Text(), map1, &slice1, &slice2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}
	return calculateSimilarityScore(slice1, map1)
}

func splitAndCountString(line string, m map[int]int, l1, l2 *[]int) {
	tempSlice := strings.Fields(line)

	num1, err := strconv.Atoi(tempSlice[0])
	if err != nil {
		panic(err)
	}
	*l1 = append(*l1, num1)

	num2, err := strconv.Atoi(tempSlice[1])
	if err != nil {
		panic(err)
	}
	m[num2]++

	*l2 = append(*l2, num2)
}

func calculateSimilarityScore(l []int, m map[int]int) int {
	score := 0
	for _, val := range l {
		if frequency, ok := m[val]; ok {
			score += frequency * val
		}
	}
	return score
}
