package day06

import (
	"advent-of-code/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input := utils.ReadFile("../../test_input/day06_test.txt")
	result := Part1(input)
	expected := 0

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// func TestPart2(t *testing.T) {
// 	input := utils.ReadFile("../../test_input/dayXY_part2_test.txt")
// 	result := Part2(input)
// 	expected := 0

// 	if result != expected {
// 		t.Errorf("Expected %d, got %d", expected, result)
// 	}
// }
