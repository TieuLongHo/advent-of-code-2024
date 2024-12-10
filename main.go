package main

import (
	"advent-of-code/solutions/day01"
	"advent-of-code/solutions/day02"
	"advent-of-code/solutions/day03"
	"advent-of-code/solutions/day04"
	"advent-of-code/solutions/day05"
	"advent-of-code/solutions/day06"
	"advent-of-code/solutions/day07"
	"advent-of-code/solutions/day08"
	"advent-of-code/solutions/day09"
	"advent-of-code/solutions/day10"
	"advent-of-code/utils"
	"fmt"
	"strings"
	"time"
)

func main() {
	runDay("Day01", day01.Part1, day01.Part2)
	runDay("Day02", day02.Part1, day02.Part2)
	runDay("Day03", day03.Part1, day03.Part2)
	runDay("Day04", day04.Part1, day04.Part2)
	runDay("Day05", day05.Part1, day05.Part2)
	runDay("Day06", day06.Part1, day06.Part2)
	runDay("Day07", day07.Part1, day07.Part2)
	runAltSolution("Day07", "Recursive", day07.Part1Recursive, day07.Part2Recursive)
	runDay("Day08", day08.Part1, day08.Part2)
	runDay("Day09", day09.Part1, day09.Part2)
	runDay("Day10", day10.Part1, day10.Part2)
}

func runDay(fnName string, fn1 func(input string) int, fn2 func(input string) int) {
	fileName := fmt.Sprintf("input/%s.txt", strings.ToLower(fnName))
	input := utils.ReadFile(fileName)

	fmt.Printf("\n%s:\n", fnName)

	executePart("Part 1", fn1, input)
	executePart("Part 2", fn2, input)
}

func runAltSolution(fnName string, altName string, fn1 func(input string) int, fn2 func(input string) int) {
	fileName := fmt.Sprintf("input/%s.txt", strings.ToLower(fnName))
	input := utils.ReadFile(fileName)

	fmt.Printf("%s:\n", altName)
	executePart("Part 1", fn1, input)
	executePart("Part 2", fn2, input)
}

func executePart(partName string, fn func(input string) int, input string) {
	startTime := time.Now()
	result := fn(input)
	elapsedTime := time.Since(startTime)

	fmt.Printf("\t%s:\t%v (Execution time: %s)\n", partName, result, elapsedTime)
}
