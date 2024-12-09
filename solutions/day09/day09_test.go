package day09

import (
	"advent-of-code/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input := utils.ReadFile("../../test_input/day09_test.txt")
	result := Part1(input)
	expected := 1928

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := utils.ReadFile("../../test_input/day09_test.txt")
	result := Part2(input)
	expected := 2858

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
