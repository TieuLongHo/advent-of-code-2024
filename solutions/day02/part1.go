package day02

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	safeCounter := 0

	for scanner.Scan() {
		checkedLine := splitAndConvertString(scanner.Text())
		isSorted, isIncreasing := sliceIsSorted(checkedLine)
		if !isSorted {
			continue
		}
		safeCounter += checkSafety(checkedLine, isIncreasing)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return safeCounter
}

func splitAndConvertString(line string) []int {
	levelsInString := strings.Fields(line)
	levelsInInt := make([]int, len(levelsInString))

	for i, e := range levelsInString {
		val, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}

		levelsInInt[i] = val
	}
	return levelsInInt
}

func sliceIsSorted(slice []int) (isSorted bool, isInc bool) {
	isIncSorted := sort.SliceIsSorted(slice, func(p, q int) bool {
		return slice[p] < slice[q]
	})
	isDecSorted := sort.SliceIsSorted(slice, func(p, q int) bool {
		return slice[p] > slice[q]
	})

	if isIncSorted || isDecSorted {
		return true, isIncSorted
	}

	return false, isIncreasing(slice)

}

func isIncreasing(slice []int) bool {
	sum := 0
	for i := 0; i < len(slice)-1; i++ {
		sum += slice[i] - slice[i+1]
	}
	return sum < 0
}

func checkSafety(levels []int, isIncreasing bool) int {
	start, end := -3, -1
	if !isIncreasing {
		start, end = 1, 3
	}
	for i, e := range levels {
		if i == 0 {
			continue
		}
		if !isDifferenceOk(levels[i-1], e, start, end) {
			return 0
		}
	}
	return 1

}

func isDifferenceOk(val1, val2, start, end int) bool {
	difference := val1 - val2
	return difference >= start && difference <= end
}
