package day01

import (
	"advent-of-code/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input := utils.ReadFile("../../test_input/day01_test.txt")
	result := Part1(input)
	expected := 11

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := utils.ReadFile("../../test_input/day01_test.txt")
	result := Part2(input)
	expected := 31

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
