package day05

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var greaterPairs [][2]int
	isRule := true
	var score int

	for scanner.Scan() {
		if scanner.Text() == "" {
			isRule = false
			continue
		}
		if isRule {
			splitAndAppend(scanner.Text(), &greaterPairs)
			continue
		}
		score += checkLine(scanner.Text(), greaterPairs, false)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	return score
}

func checkLine(line string, greaterPairs [][2]int, isPart2 bool) int {
	pageList := strings.Split(line, ",")
	intSlice, err := convertStringToInt([]string(pageList))
	if err != nil {
		return 0
	}
	if !isSorted(intSlice, greaterPairs) {
		if isPart2 {
			return sort(intSlice, greaterPairs)[len(intSlice)/2]
		}
		return 0

	}
	if isPart2 {
		return 0
	}
	return intSlice[len(intSlice)/2]
}

func convertStringToInt(strSlice []string) (newIntSlice []int, err error) {
	intSlice := make([]int, len(strSlice))
	for i, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting '%s' to integer: %v\n", str, err)
			return []int{}, fmt.Errorf("error converting")
		}
		intSlice[i] = num
	}
	return intSlice, nil
}

func splitAndAppend(line string, slice *[][2]int) {
	parts := strings.Split(line, "|")
	firstPart, _ := strconv.Atoi(parts[0])
	secondPart, _ := strconv.Atoi(parts[1])
	*slice = append(*slice, [2]int{firstPart, secondPart})
}

func isSorted(slice []int, greaterPairs [][2]int) bool {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slices.Contains(greaterPairs, [2]int{slice[j], slice[i]}) {
				return false
			}
		}
	}
	return true
}
