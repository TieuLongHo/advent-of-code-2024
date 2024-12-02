package day02

import (
	"bufio"
	"strings"
)

func Part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	safeCounter := 0
	for scanner.Scan() {
		checkedLine := splitAndConvertString(scanner.Text())

		_, isIncreasing := sliceIsSorted(checkedLine)
		if isSafe := checkSafety(checkedLine, isIncreasing); isSafe == 1 {
			safeCounter++
			continue
		} else if damp(checkedLine) {
			safeCounter++
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return safeCounter
}

func damp(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		tempSlice := make([]int, 0, len(levels)-1)
		tempSlice = append(tempSlice, levels[:i]...)
		tempSlice = append(tempSlice, levels[i+1:]...)

		if isSorted, isIncreasing := sliceIsSorted(tempSlice); isSorted {
			safety := checkSafety(tempSlice, isIncreasing)
			if safety == 1 {
				return true
			}
		}
	}
	return false
}
