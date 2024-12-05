package day06

import (
	"bufio"
	"fmt"
	"strings"
)

func Part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	return 0
}
