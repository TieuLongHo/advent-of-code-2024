package main

import (
	"advent-of-code/solutions/day01"
	"advent-of-code/solutions/day02"
	"advent-of-code/solutions/day03"
	"advent-of-code/solutions/day04"
	"advent-of-code/solutions/day05"
	"advent-of-code/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile("input/day01.txt")
	fmt.Println("Day 1:")
	fmt.Println("\tPart 1:", day01.Part1(input))
	fmt.Println("\tPart 2:", day01.Part2(input))
	fmt.Println()

	input = utils.ReadFile("input/day02.txt")
	fmt.Println("Day 2:")
	fmt.Println("\tPart 1:", day02.Part1(input))
	fmt.Println("\tPart 2:", day02.Part2(input))
	fmt.Println()

	input = utils.ReadFile("input/day03.txt")
	fmt.Println("Day 3:")
	fmt.Println("\tPart 1:", day03.Part1(input))
	fmt.Println("\tPart 2:", day03.Part2(input))
	fmt.Println()

	input = utils.ReadFile("input/day04.txt")
	fmt.Println("Day 4:")
	fmt.Println("\tPart 1:", day04.Part1(input))
	fmt.Println("\tPart 2:", day04.Part2(input))
	fmt.Println()

	input = utils.ReadFile("input/day05.txt")
	fmt.Println("Day 5:")
	fmt.Println("\tPart 1:", day05.Part1(input))
	fmt.Println("\tPart 2:", day05.Part2(input))
	fmt.Println()
}
