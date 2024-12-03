package main

import (
	"advent-of-code/solutions/day01"
	"advent-of-code/solutions/day02"
	"advent-of-code/solutions/day03"
	"advent-of-code/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile("input/day01.txt")
	fmt.Println("Day 1, Part 1:", day01.Part1(input))
	fmt.Println("Day 1, Part 2:", day01.Part2(input))

	input = utils.ReadFile("input/day02.txt")
	fmt.Println("Day 2, Part 1:", day02.Part1(input))
	fmt.Println("Day 2, Part 2:", day02.Part2(input))

	input = utils.ReadFile("input/day03.txt")
	fmt.Println("Day 3, Part 1:", day03.Part1(input))
	fmt.Println("Day 3, Part 2:", day03.Part2(input))
}
