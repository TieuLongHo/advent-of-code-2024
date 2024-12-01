package day01

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	slice1 := make([]int, 0)
	slice2 := make([]int, 0)
	for scanner.Scan() {
		splitString(scanner.Text(), &slice1, &slice2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	sort.Ints(slice1)
	sort.Ints(slice2)

	return calculateDist(slice1, slice2)
}

func splitString(line string, l1, l2 *[]int) {
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
	*l2 = append(*l2, num2)
}

func calculateDist(slice1, slice2 []int) int {
	distance := 0
	for i := range slice1 {
		diff := slice1[i] - slice2[i]
		if diff < 0 {
			diff *= -1
		}
		distance += diff
	}
	return distance
}
