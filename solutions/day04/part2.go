package day04

import (
	"bufio"
	"fmt"
	"strings"
	"sync"
)

func Part2(input string) int {
	ch := make(chan bool)
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(strings.NewReader(input))
	var grid [][]byte

	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return 0
	}
	findA(grid, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	count := 0
	for e := range ch {
		if e {
			count++
		}
	}

	return count
}

func findA(grid [][]byte, ch chan<- bool, wg *sync.WaitGroup) {
	rows := len(grid)
	if rows == 0 {
		return
	}
	cols := len(grid[0])

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if grid[i][j] == 'A' {
				wg.Add(1)
				go isValidXMAS(grid, i, j, ch, wg)
			}
		}
	}
}

func isValidXMAS(grid [][]byte, i, j int, ch chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()

	rows := len(grid)
	cols := len(grid[0])
	var patterns = [][]byte{
		[]byte("MAS"),
		[]byte("SAM"),
	}

	validDiag1 := false
	validDiag2 := false

	for _, pattern := range patterns {
		if i-1 >= 0 && j-1 >= 0 && i+1 < rows && j+1 < cols {
			if grid[i-1][j-1] == pattern[0] && grid[i+1][j+1] == pattern[2] {
				validDiag1 = true
				break
			}
		}
	}

	for _, pattern := range patterns {
		if i-1 >= 0 && j+1 < cols && i+1 < rows && j-1 >= 0 {
			if grid[i-1][j+1] == pattern[0] && grid[i+1][j-1] == pattern[2] {
				validDiag2 = true
				break
			}
		}
	}

	ch <- validDiag1 && validDiag2
}
