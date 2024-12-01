package main

import (
	"advent-of-code/solutions/day01"
	"advent-of-code/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile("input/day01.txt")
	fmt.Println("Day 1, Part 1:", day01.Part1(input))
	fmt.Println("Day 1, Part 2:", day01.Part2(input))

	input = utils.ReadFile("input/day02.txt")
}
