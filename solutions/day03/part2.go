package day03

import (
	"bufio"
	"strings"
)

func Part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return 0
}
