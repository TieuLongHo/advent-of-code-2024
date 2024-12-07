package day06

import (
	"bufio"
	"fmt"
	"strings"
)

func Part1(input string) int {
	visited, _, _, _, _ := helper(input)
	return visited
}
func helper(input string) (visitedCount, newIStart, newJStart int, originalMatrix [][]byte, visitedPlaces map[[2]int]bool) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var inputMatrix [][]byte

	var iStart, jStart int
	foundStart := false
	i := 0

	for scanner.Scan() {
		newLine := []byte(scanner.Text())
		if !foundStart {
			for j := 0; j < len(newLine); j++ {
				if newLine[j] == '^' {
					iStart, jStart = i, j
					foundStart = true
				}
			}
		}
		inputMatrix = append(inputMatrix, newLine)
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	visited := simulateGuardPositions(inputMatrix, iStart, jStart)
	return len(visited), iStart, jStart, inputMatrix, visited
}

func simulateGuardPositions(grid [][]byte, iStart, jStart int) map[[2]int]bool {
	directions := [][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	dirSymbol := [4]byte{
		'^', // Up
		'>', // Right
		'v', // Down
		'<', // Left
	}

	direction := 0

	visited := make(map[[2]int]bool)
	visited[[2]int{iStart, jStart}] = true

	row, col := iStart, jStart

	for {
		newRow := row + directions[direction][0]
		newCol := col + directions[direction][1]

		// case: out of bounds
		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
			break
		}

		// case: obstruction
		if grid[newRow][newCol] == '#' {
			direction = (direction + 1) % 4
			continue
		}

		row, col = newRow, newCol
		visited[[2]int{row, col}] = true
		grid[row][col] = dirSymbol[direction]
	}

	return visited
}
