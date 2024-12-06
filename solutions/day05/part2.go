package day05

import (
	"bufio"
	"fmt"
	"strings"
)

func Part2(input string) int {
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
		score += checkLine(scanner.Text(), greaterPairs, true)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	return score
}

// func sort(slice []int, greaterPairs [][2]int) []int {
// 	for i := 0; i < len(slice)-1; i++ {
// 		for j := i + 1; j < len(slice); j++ {
// 			if slices.Contains(greaterPairs, [2]int{slice[j], slice[i]}) {
// 				itemJ := slice[j]
// 				newSlice := make([]int, 0)
// 				newSlice = append(newSlice, slice[:j]...)
// 				newSlice = append(newSlice, slice[j+1:]...)
// 				newSlice = append(newSlice[:i], append([]int{itemJ}, newSlice[i:]...)...)
// 				return sort(newSlice, greaterPairs)

// 			}
// 		}
// 	}
// 	return slice
// }
