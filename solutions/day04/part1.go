package day04

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
)

var XMAS = []byte("XMAS")
var WORD_LEN = len(XMAS)

func Part1(input string) int {
	ch := make(chan bool)
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(strings.NewReader(input))
	var inputMatrix [][]byte

	for scanner.Scan() {
		inputMatrix = append(inputMatrix, []byte(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return 0
	}

	findXmas(inputMatrix, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()
	var counter int
	for range ch {
		counter++
	}
	return counter
}

func findXmas(grid [][]byte, ch chan<- bool, wg *sync.WaitGroup) {

	rows := len(grid)
	if rows == 0 {
		return
	}
	cols := len(grid[0])

	//direction vectors
	directions := [][2]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Up-Left
		{-1, 1},  // Up-Right
		{1, -1},  // Down-Left
		{1, 1},   // Down-Right
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'X' {
				for _, dir := range directions {
					wg.Add(1)
					go isValid(grid, i, j, dir[0], dir[1], ch, wg)
				}
			}
		}
	}
}

func isValid(grid [][]byte, i, j, rowDelta, colDelta int, ch chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for k := 0; k < WORD_LEN; k++ {
		row := i + k*rowDelta
		col := j + k*colDelta
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] != XMAS[k] {
			return
		}
	}
	ch <- true
}
