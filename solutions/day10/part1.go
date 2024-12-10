package day10

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input string) int {
	return helper(input, false)
}

func helper(input string, isPart2 bool) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var grid [][]int
	var trailHeads [][2]int

	var i int
	for scanner.Scan() {
		newLine := scanner.Text()
		newLineInt := make([]int, len(newLine))
		for j, e := range newLine {
			num, _ := strconv.Atoi(string(e))
			if num == 0 {
				trailHeads = append(trailHeads, [2]int{i, j})
			}
			newLineInt[j] = num
		}
		grid = append(grid, newLineInt)
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	return solve(grid, trailHeads, isPart2)
}

func solve(grid [][]int, trailHeads [][2]int, isPart2 bool) int {
	var score int
	for _, e := range trailHeads {
		visited := make(map[[2]int]bool)
		score += DFU(grid, e, visited, isPart2)
	}
	return score
}

func DFU(grid [][]int, node [2]int, visited map[[2]int]bool, isPart2 bool) int {
	directions := [4][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}
	var score int
	visited[node] = true

	if grid[node[0]][node[1]] == 9 {
		return 1
	}

	for _, dir := range directions {
		i := node[0] + dir[0]
		j := node[1] + dir[1]
		// case: out of bound
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid) {
			continue
		}
		// case: height difference different than +1
		if grid[i][j]-grid[node[0]][node[1]] != 1 {
			continue
		}
		// case: already visited
		if _, ok := visited[[2]int{i, j}]; ok && !isPart2 {
			continue
		}
		score += DFU(grid, [2]int{i, j}, visited, isPart2)
	}
	return score
}
