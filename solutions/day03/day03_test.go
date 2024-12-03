package day03

import (
	"advent-of-code/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input := utils.ReadFile("../../test_input/day03_test.txt")
	result := Part1(input)
	expected := 161

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// func TestPart2(t *testing.T) {
// 	input := utils.ReadFile("../../test_input/day02_test.txt")
// 	result := Part2(input)
// 	expected := 4

// 	if result != expected {
// 		t.Errorf("Expected %d, got %d", expected, result)
// 	}
// }
