package day03

import (
	"bufio"
	"strings"
)

var DO_PATTERN = [4]string{"d", "o", "(", ")"}
var DONT_PATTERN = [7]string{"d", "o", "n", "'", "t", "(", ")"}

func Part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	score := 0

	for scanner.Scan() {
		stringSlice := strings.Split(scanner.Text(), "")

		score += parseFunction(stringSlice, 0, true)

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return score
}

// returns
//
//	input: Puzzle input
//	0 -> mismatch,
//	1 -> do(),
//	2 -> don't()
func parseDoFunction(input []string, index int) (do, newIndex int) {

	isDoFunc := false

	for i, e := range DONT_PATTERN {
		switch input[index] {
		// case: don't()
		case e:
			if !isDoFunc {
				index++
				if len(DONT_PATTERN)-1 == i {
					return 2, index
				}
				continue
			}

		// case: do()
		case DO_PATTERN[i]:
			index++
			isDoFunc = true
			if len(DO_PATTERN)-1 == i {
				return 1, index
			}
			continue

		// case: mismatch
		default:
			return 0, index
		}
	}
	// case: incomplete function
	return 0, index
}
