package day09

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var newLine string

	for scanner.Scan() {
		newLine = (scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	return solvePart2(newLine)
}

func solvePart2(input string) int {
	compressBlocks := compressFileBlockwise(mapToFileBlock(input))
	var score int
	for _, e := range compressBlocks {
		if e.id == -1 {
			continue
		}
		score += e.id * (e.getLength()) * (e.start + e.end) / 2
	}
	return score
}

func compressFileBlockwise(fileBlocks []Block) []Block {
	var offsetCount int
	compressedBlocks := make([]Block, 0)
	compressedBlocks = append(compressedBlocks, fileBlocks...)
	for i := len(fileBlocks) - 1; i >= 0; i-- {
		for j, e := range compressedBlocks {
			if i+offsetCount <= j || compressedBlocks[i+offsetCount].id == -1 {
				break
			}
			if e.id != -1 {
				continue
			}
			newBlock, err := e.splitBlock(&compressedBlocks[i+offsetCount])
			if err != nil {
				continue
			}
			if !newBlock.IsEmpty() {
				compressedBlocks = append(compressedBlocks[:j+1], append([]Block{newBlock}, compressedBlocks[j+1:]...)...)
				offsetCount++
			}
			tempBlock := compressedBlocks[j]
			compressedBlocks[j] = compressedBlocks[i+offsetCount] //
			compressedBlocks[i+offsetCount] = tempBlock           //
			break

		}
	}
	return compressedBlocks

}

func mapToFileBlock(input string) []Block {
	var fileBlocks []Block
	start := 0
	for i, e := range input {
		fileVal, _ := strconv.Atoi(string(e))
		var blockId int
		if i%2 == 0 {
			blockId = i / 2
		} else {
			blockId = -1
		}
		if fileVal == 0 {
			continue
		}
		fileBlocks = append(fileBlocks, Block{blockId, start, start + fileVal - 1})
		start += fileVal
	}
	return fileBlocks
}
