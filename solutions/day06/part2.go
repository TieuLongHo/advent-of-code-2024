package day06

import "sync"

func Part2(input string) int {
	_, iStart, jStart, originalGrid, visited := helper(input)

	loopCount := 0

	var wg sync.WaitGroup
	ch := make(chan bool)

	for pos := range visited {
		if pos == [2]int{iStart, jStart} {
			continue
		}
		gridCopy := deepCopyGrid(originalGrid)
		gridCopy[pos[0]][pos[1]] = '#'

		wg.Add(1)
		go detectLoop(gridCopy, iStart, jStart, ch, &wg)

	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for range ch {
		loopCount++
	}

	return loopCount
}
func detectLoop(grid [][]byte, iStart, jStart int, ch chan<- bool, wg *sync.WaitGroup) {

	defer wg.Done()

	directions := [][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	direction := 0

	visitedStates := make(map[[3]int]bool)

	row, col := iStart, jStart
	visitedStates[[3]int{row, col, direction}] = true

	for {
		newRow := row + directions[direction][0]
		newCol := col + directions[direction][1]

		// case: out of bounds
		if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
			return
		}

		// case: obstruction
		if grid[newRow][newCol] == '#' {
			direction = (direction + 1) % 4
			continue
		}

		row, col = newRow, newCol

		state := [3]int{row, col, direction}
		// case: already visited in same dir -> loop
		if visitedStates[state] {
			ch <- true
			return
		}
		visitedStates[state] = true
	}
}

func deepCopyGrid(grid [][]byte) [][]byte {
	copyGrid := make([][]byte, len(grid))
	for i := range grid {
		copyGrid[i] = make([]byte, len(grid[i]))
		copy(copyGrid[i], grid[i])
	}
	return copyGrid
}
