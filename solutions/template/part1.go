package dayXY

import (
	"bufio"
	"fmt"
	"strings"
)

func Part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	return 0
}
