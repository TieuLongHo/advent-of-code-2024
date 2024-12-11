package day11

import (
	"advent-of-code/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	input := utils.ReadFile("../../test_input/day11_test.txt")
	result := Part1(input)
	expected := 55312

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
