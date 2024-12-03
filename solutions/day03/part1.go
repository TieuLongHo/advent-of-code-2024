package day03

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var TOKEN = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var PATTERN = [8]string{"m", "u", "l", "(", "X", ",", "Y", ")"}

const MAX_PARAM_LEN = 3

func Part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	score := 0

	for scanner.Scan() {
		stringSlice := strings.Split(scanner.Text(), "")

		score += parseFunction(stringSlice, 0)

	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return score
}

func parseFunction(input []string, index int) int {
	var paramX int
	var paramXAssigned bool
	var paramY int
	var paramYAssigned bool
	var sum int
	var err error

	for _, e := range PATTERN {
		if index == len(input) {
			return sum
		}

		//case: X and Y handling
		if e == "X" {
			paramX, index, err = parseParameter(input, index)
			if err != nil {
				return sum + parseFunction(input, index)
			}
			paramXAssigned = true
			continue
		} else if e == "Y" {
			paramY, index, err = parseParameter(input, index)
			if err != nil {

				return sum + parseFunction(input, index)
			}
			paramYAssigned = true
			continue
		}

		//case: closing parentheses
		if e == ")" && input[index] == ")" {
			if paramXAssigned && paramYAssigned {
				sum += paramX * paramY
			}
			return sum + parseFunction(input, index+1)
		}
		//case: mismatch symbol
		if e != input[index] {
			return sum + parseFunction(input, index+1)
		}

		index++
	}

	return sum + parseFunction(input, index+1)
}

func parseParameter(input []string, index int) (parameter, newIndex int, err error) {
	parameterString := ""

	for i := 0; i <= MAX_PARAM_LEN; i++ {
		if index == len(input)-1 && input[index] != ")" {
			return 0, index, fmt.Errorf("EOF in Param")
		}
		if i > 0 && (input[index] == "," || input[index] == ")") {
			parameterInt, err := strconv.Atoi(parameterString)
			return parameterInt, index, err
		}
		if slices.Contains(TOKEN, input[index]) {
			parameterString += input[index]
			index++
			continue
		}

		return 0, index, fmt.Errorf("mismatch")

	}
	parameterInt, err := strconv.Atoi(parameterString)
	return parameterInt, index + 1, err
}
